package environment

import (
	"os"
	"path/filepath"
	"strings"
)

//SetupRunnerFileSystem sets up a file system in /securefs/. It assumes that
//the file /securefs/ already exists.
//The paramater "rootName" should be the name of the folder you want created
//This method returns the absolute path to the file system and an error if it
//encounters any problems setting up the file system.
func SetupRunnerFileSystem(rootName string) (string, error) {

	var neededFiles = []string{
		"bin",
		"sbin",
		"lib",
		"lib64",
		"usr",
		"etc",
		"var",
		"boot",
		"media",
		"mnt",
		"root",
		"srv",
		"sys",
		"dev",
		"proc",
		"runner_files",
	}

	var rootFilePath strings.Builder
	rootFilePath.WriteString("/securefs/")
	rootFilePath.WriteString(rootName)

	rootPath := rootFilePath.String()

	err := os.Mkdir(rootPath, 0755)
	if err != nil {
		return "", err
	}

	for _, file := range neededFiles {
		err = os.Mkdir(filepath.Join(rootPath, file), 0755)
		if err != nil {
			return "", err
		}
	}

	return rootPath, nil
}

//RemoveRunnerFileSystem deltes the file at /securesfs/xxxx.
//WARNING: Do not call this method if you have anything mounted to this location
//doing so could cause unwanted behavior and/or break your root file system.
//rootPath is the absolute path to the folder you want removed. It should be
//safe to pass in the string you received from SetupRunnerFileSystem .
func RemoveRunnerFileSystem(rootPath string) error {
	return os.RemoveAll(rootPath)
}
