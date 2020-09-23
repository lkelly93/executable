package executable

import (
	"os"
	"path/filepath"
	"strings"
)

func setupRunnerFileSystem(rootName string) (string, error) {

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
	rootFilePath.WriteString("/")
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

//DO NOT CALL THIS BEFORE UNMOUNTING! You will delete your entire root filesystem.
func teardownRunnerFileSystem(rootPath string) error {
	return os.RemoveAll(rootPath)
}
