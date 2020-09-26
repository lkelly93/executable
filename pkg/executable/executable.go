package executable

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"syscall"
	"time"

	"github.com/lkelly93/executable/internal/environment"
	"github.com/lkelly93/executable/internal/runners"
)

//NewExecutable returns something that implements the Executable interface
//If the given language is not supported NewExecutable will throw an error.
//A uniqueIdentifier is required, this package does not check if it is actually
//unique but you must give a non-empty string as an argument. If it is not
//unique it could cause a data race and/or unknown behavior.
func NewExecutable(lang, code, uniqueIdentifier string) (Executable, error) {
	if len(uniqueIdentifier) == 0 {
		return nil, &IllegalUniqueIdentifier{
			ErrMessage: "The field \"uniqueIdentifier\" cannot be empty.",
		}
	}
	runner := runners.GetRunner(lang)
	if runner != nil {
		return &executableState{
			code:             code,
			lang:             lang,
			uniqueIdentifier: uniqueIdentifier,
			runner:           runner,
		}, nil
	}

	return nil, &UnsupportedLanguageError{Lang: lang}
}

//Run will run the executable in a secure container. It returns the output
//of the program and/or an error. See errors.go for more all the possible
//errors it can return.
func (state *executableState) Run() (string, error) {
	uniqueID := state.uniqueIdentifier
	//Setup the executable's new root file system.
	envData, err := environment.Setup(uniqueID)
	if err != nil {
		return "", fatalServerError(err, uniqueID)
	}

	//Create the runner file.
	sysCommand, fileName, err := state.runner.CreateFile(
		state.code,
		filepath.Join(envData.RootPath, "runner_files"),
	)
	if err != nil {
		return "", fatalServerError(err, uniqueID)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(MaxExecutableRunTime)*time.Second,
	)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"executor",
		envData.RootPath,
		uniqueID,
		sysCommand,
		fileName,
	)

	var stdOut bytes.Buffer
	cmd.Stdout = &stdOut
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}

	err = cmd.Run()
	output, outputErr := getOutput(stdOut, stdErr)
	checkLoggerFile(envData.RootPath)

	//Parse output and check for all possible errors
	if ctx.Err() == context.DeadlineExceeded {
		log.Println(err)
		if errCleanup := envData.CleanUp(); errCleanup != nil {
			return output, fatalServerError(errCleanup, uniqueID)
		}
		return output, &TimeLimitExceededError{MaxTime: MaxExecutableRunTime}
	}
	if err != nil {
		log.Println(err)
		if errCleanup := envData.CleanUp(); errCleanup != nil {
			return output, fatalServerError(errCleanup, uniqueID)
		}
		return output, &RuntimeError{ErrMessage: err.Error()}
	}

	if stdErr.Len() != 0 {
		if errCleanup := envData.CleanUp(); errCleanup != nil {
			return output, fatalServerError(errCleanup, uniqueID)
		}
		return output, &RuntimeError{ErrMessage: outputErr}
	}

	if errCleanup := envData.CleanUp(); errCleanup != nil {
		return output, fatalServerError(errCleanup, uniqueID)
	}

	return output, nil
}

func fatalServerError(err error, uniqueID string) error {
	log.Println(err)
	return &SystemError{
		Err: err,
	}
}

func checkLoggerFile(rootPath string) {
	fileLoc := filepath.Join(rootPath, "log/serverOutput.log")

	output, err := ioutil.ReadFile(fileLoc)

	if err != nil && os.IsExist(err) {
		log.Println("Couldn't open the logger file but it does exist.")
		log.Println(err.Error())
	}

	if output != nil {
		log.Println(string(output))
	}

	os.RemoveAll(fileLoc)
}

func parseOutput(message []byte) string {
	//Remove unneeded time stamp.
	regex, err := regexp.Compile("[0-9]{4}/[0-9]{2}/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2} ")
	if err != nil {
		log.Fatal("Could not compile regex expression.")
	}

	return string(regex.ReplaceAll(message, []byte("")))
}

func getOutput(stdOut, stdErr bytes.Buffer) (string, string) {
	var out string = ""
	var err string = ""
	if stdOut.Len() != 0 {
		out = parseOutput(stdOut.Bytes())
	}
	if stdErr.Len() != 0 {
		err = parseOutput(stdErr.Bytes())
	}

	return out, err
}
