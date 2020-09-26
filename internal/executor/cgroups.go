package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func addPIDtoCGroup(rootName string) error {
	return ioutil.WriteFile(
		filepath.Join(
			"/sys/fs/cgroup/pids",
			rootName,
			"cgroup.procs",
		),
		[]byte(strconv.Itoa(os.Getpid())),
		0700,
	)
}
