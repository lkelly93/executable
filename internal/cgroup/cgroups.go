package cgroup

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//Setup sets up all the cgroup settings for an executable.
func Setup(rootName string) error {
	cgroup := "/sys/fs/cgroup"
	if err := setupPidsCGroup(cgroup, rootName); err != nil {
		return err
	}
	if err := setupMemoryCGroup(cgroup, rootName); err != nil {
		return err
	}

	return nil
}

//GetMemoryUsage gets the memory usage for the given executable.
func GetMemoryUsage(rootName string) (string, error) {
	memoryLoc := filepath.Join(
		"/sys/fs/cgroup/memory/",
		rootName,
		"memory.max_usage_in_bytes",
	)
	memoryUsage, err := ioutil.ReadFile(memoryLoc)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(memoryUsage), "\n"), nil
}

//AddPIDToCGroup adds the given pid to all the implemented CGroups for a given
//executable rootName.
func AddPIDToCGroup(rootName string, pid []byte) error {
	for i := 0; i < len(usedCGroups); i++ {
		err := ioutil.WriteFile(
			filepath.Join(
				"/sys/fs/cgroup",
				usedCGroups[i],
				rootName,
				"cgroup.procs",
			),
			pid,
			0700,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

//Cleanup will remove all the cgroup files left over after a executable has been
//run.
func Cleanup(rootName string) error {
	cgroup := "/sys/fs/cgroup"

	for i := 0; i < len(usedCGroups); i++ {
		err := os.RemoveAll(filepath.Join(
			cgroup,
			usedCGroups[i],
			rootName),
		)

		if err != nil {
			return err
		}
	}

	return nil
}
