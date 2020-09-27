package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/lkelly93/executable/internal/cgroup"
)

func main() {
	if len(os.Args) != 5 {
		var message strings.Builder
		message.WriteString("This file should only be run through the executable package. ")
		message.WriteString("Failure to do so could cause damage to the system it is running on.\n")
		log.Println(message.String())
		os.Exit(242)
	}
	initContainerAndRunProgram()
}

func initContainerAndRunProgram() {
	rootPath := os.Args[1]
	rootName := os.Args[2]
	sysCommand := os.Args[3]
	fileName := os.Args[4]

	err := cgroup.AddPIDToCGroup(rootName, []byte(strconv.Itoa(os.Getpid())))
	if err != nil {
		serverFatalBeforeChroot(rootPath, err)
	}

	err = syscall.Chroot(rootPath)
	if err != nil {
		serverFatalBeforeChroot(rootPath, err)
	}

	err = os.Chdir("/")
	if err != nil {
		serverFatal(err)
	}

	err = syscall.Sethostname([]byte("runner"))
	if err != nil {
		serverFatal(err)
	}

	mountAllTempDirs()

	runProgramInContainer(
		sysCommand,
		filepath.Join("/runner_files", fileName),
	)
}

func runProgramInContainer(sysCommand string, fileName string) {
	cmd := exec.Command(sysCommand, fileName)

	var stdErr bytes.Buffer
	var stdOut bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()

	if err != nil || stdErr.Len() != 0 {
		log.Print(parseOutput(stdErr.String()))
	}

	fmt.Print(parseOutput(stdOut.String()))
}

func parseOutput(message string) string {
	return strings.ReplaceAll(message, "/runner_files/", "")
}
