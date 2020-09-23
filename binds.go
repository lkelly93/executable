package executable

import (
	"os/exec"
	"path/filepath"
	"syscall"
)

var neededBinds = []string{
	"usr",
	"bin",
	"sbin",
	"lib",
	"lib64",
	"etc",
	"var",
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
	return nil
}

func teardownAllFileSystemBinds(rootPath string) error {

	for i := 0; i < len(neededBinds); i++ {
		target := filepath.Join(rootPath, neededBinds[i])
		err := syscall.Unmount(target, 0)
		// err := runCommand(
		// 	"umount",
		// 	filepath.Join(rootPath, bindLoc),
		// )
		if err != nil {
			return err
		}
	}
	return nil
}

func runCommand(command string, args ...string) error {
	return exec.Command(command, args[0:]...).Run()
}
