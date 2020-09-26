package environment

import (
	"os/exec"
	"path/filepath"

	"github.com/otiai10/copy"
)

var neededBinds = map[string]bool{
	"lib":   true,
	"lib64": true,
	"bin":   true,
	"sbin":  true,
	"usr":   true,
	"dev":   false,
}

//bindAndCopyRequiredFiles binds all the required folders and copies over the
//the data in the folders that can not be bound.
func bindAndCopyRequiredFiles(rootPath string) error {

	// for i := 0; i < len(neededBinds); i++ {
	for bindLoc, readOnly := range neededBinds {
		if readOnly {
			err := runCommand(
				"mount",
				"--rbind",
				"--make-runbindable",
				"--read-only",
				filepath.Join("/", bindLoc),
				filepath.Join(rootPath, bindLoc),
			)
			if err != nil {
				return err
			}
		} else {
			err := runCommand(
				"mount",
				"--rbind",
				"--make-runbindable",
				filepath.Join("/", bindLoc),
				filepath.Join(rootPath, bindLoc),
			)
			if err != nil {
				return err
			}
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

//unbindAll unbinds everything that was previously bound from the root.
func unbindAll(rootPath string) error {
	for bind := range neededBinds {
		target := filepath.Join(rootPath, bind)
		err := runCommand("umount", "-l", target)
		if err != nil {
			return err
		}
	}

	return nil
}

func runCommand(command string, args ...string) error {
	return exec.Command(command, args[0:]...).Run()
}
