package main

import (
	"log"
	"os"
	"path/filepath"
)

//serverPrintln assumes chroot as already been called.
func serverPrintln(message string) {
	writeToFile(message, "/log/serverOutput.log")

}

//serverFatal assumes chroot as already been called.
func serverFatal(err error) {
	serverPrintln(err.Error())
	os.Exit(100)
}

func serverFatalBeforeChroot(rootPath string, err error) {
	writeToFile(
		err.Error(),
		filepath.Join(rootPath, "log/serverOutput.log"),
	)
	os.Exit(100)
}

func writeToFile(message string, fileLocation string) {
	fileptr, err := os.OpenFile(
		fileLocation,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0700,
	)

	if err != nil {
		log.Fatalf(
			"Couldn't create server logging output file - %s",
			err.Error(),
		)
	}

	_, err = fileptr.Write([]byte(message + "\n"))

	if err != nil {
		log.Fatalf(
			"Could not write to logging file after it was created. - %s",
			err.Error(),
		)
	}

}
