package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	if len(os.Args) != 4 {
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

	sysCommand := os.Args[2]
	fileName := os.Args[3]

	err := syscall.Chroot(rootPath)
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

func runProgramInContainer(sysCommand string, fileLocation string) {
	cmd := exec.Command(sysCommand, fileLocation)

	var stdErr bytes.Buffer
	var stdOut bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()

	if err != nil || stdErr.Len() != 0 {
		log.Print(stdErr.String())
		// log.Print(parseOutput(stdErr.String(), fileLocation, fileNamePrefix))
	}

	fmt.Print(stdOut.String())
	// fmt.Print(parseOutput(stdOut.String(), fileLocation, fileNamePrefix))
}
