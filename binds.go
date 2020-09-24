package executable

import (
	"os/exec"
	"path/filepath"

	"github.com/otiai10/copy"
)

var neededBinds = []string{
	"lib",
	"lib64",
	"bin",
	"sbin",
	"usr",
	// "var",
	// "etc",
}

func setupAllFileSystemBinds(rootPath string) error {

	for i := 0; i < len(neededBinds); i++ {
		bindLoc := neededBinds[i]
		err := runCommand(
			"mount",
			"--rbind",
			filepath.Join("/", bindLoc),
			filepath.Join(rootPath, bindLoc),
		)
		if err != nil {
			return err
		}
	}

	err := copy.Copy("/var", filepath.Join(rootPath, "var"))
	if err != nil {
		return err
	}
	err = copy.Copy("/etc", filepath.Join(rootPath, "etc"))
	if err != nil {
		return err
	}
	return nil
}

func teardownAllFileSystemBinds(rootPath string) error {

	for i := 0; i < len(neededBinds); i++ {
		target := filepath.Join(rootPath, neededBinds[i])
		err := runCommand("umount", target)
		if err != nil {
			return err
		}
	}

	return nil
}

func runCommand(command string, args ...string) error {
	return exec.Command(command, args[0:]...).Run()
}
