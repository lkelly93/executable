package main

import (
	"syscall"
)

func mountAllTempDirs() {
	mountProc()
	mountSys()
}

func mountProc() {
	source := "proc"
	fstype := "proc"
	target := "/proc"
	flags := uintptr(0)
	data := ""

	err := syscall.Mount(source, target, fstype, flags, data)
	if err != nil {
		serverFatal(err)
	}
}

func mountSys() {
	source := "sysfs"
	fstype := "sysfs"
	target := "/sys"
	flags := uintptr(0)
	data := ""

	err := syscall.Mount(source, target, fstype, flags, data)
	if err != nil {
		serverFatal(err)
	}
}
