package executable

import (
	"bytes"
	"fmt"
	"os"
)

//NewExecutable returns something that implements the Executable interface
//If the given language is not supported NewExecutable will throw an error.
//If Settings is nil the default settings will be used for that language
func NewExecutable(lang, code string, settings *Settings) (Executable, error) {
	return &executableState{
		code:     code,
		lang:     lang,
		settings: settings,
	}, nil
}

//Run TODO: COMMENT
func (state *executableState) Run() (string, error) {
	uniqueID := state.settings.UniqueIdentifier
	//Setup the executable's new root file system.
	rootPath, err := setupRunnerFileSystem(uniqueID)
	if err != nil {
		return "", fatalServerError(err, uniqueID)
	}

	//Bind all the needed files
	err = setupAllFileSystemBinds(rootPath)
	if err != nil {
		return "", fatalServerError(err, uniqueID)
	}

	err = cleanUp(rootPath)
	if err != nil {
		return "", fatalServerError(err, uniqueID)
	}

	return "", nil
}

func cleanUp(rootPath string) error {
	err := teardownAllFileSystemBinds(rootPath)
	if err != nil {
		return err
	}
	err = teardownRunnerFileSystem(rootPath)
	if err != nil {
		return err
	}

	return nil
}

func fatalServerError(err error, uniqueID string) error {
	var errMessage bytes.Buffer
	errMessage.WriteString(
		fmt.Sprintf("Error while executing %s: - ", uniqueID),
	)
	errMessage.WriteString(err.Error())
	errMessage.WriteString("\n")
	os.Stderr.Write(errMessage.Bytes())
	// log.Fatal(err)
	// log.Println(err)
	return &SystemError{}
}
