package environment

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

//BindAndCopyRequiredFiles binds all the required folders and copies over the
//the data in the folders that can not be bound.
func BindAndCopyRequiredFiles(rootPath string) error {

	for i := 0; i < len(neededBinds); i++ {
		bindLoc := neededBinds[i]
		err := runCommand(
			"mount",
			//Recursively bind these folders to the destination
			"--rbind",
			//Make it so this file system can't be cloned
			"--make-runbindable",
			//Make it so the running code can't write to the root file system
			"--read-only",
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

//UnbindAll unbinds everything that was previously bound from the root.
func UnbindAll(rootPath string) error {

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
