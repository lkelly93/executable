package cgroup

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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

	if err := setupCPUCGroup(cgroup, rootName); err != nil {
		return err
	}

	return nil
}

//GetMemoryUsage gets the memory usage for the given executable.
func GetMemoryUsage(rootName string) (int, error) {
	memoryLoc := filepath.Join(
		"/sys/fs/cgroup/memory/",
		rootName,
		"memory.max_usage_in_bytes",
	)
	memoryUsage, err := ioutil.ReadFile(memoryLoc)
	if err != nil {
		return -1, err
	}

	memoryUsageString := strings.TrimSuffix(string(memoryUsage), "\n")
	return strconv.Atoi(string(memoryUsageString))
}

//GetCPUTime gets the time the executable spent on the cpu in USER_HZ time.
func GetCPUTime(rootName string) (int, error) {
	timeLoc := filepath.Join(
		"/sys/fs/cgroup/cpuacct/",
		rootName,
		"cpuacct.stat",
	)

	userHz, err := ioutil.ReadFile(timeLoc)
	if err != nil {
		return -1, err
	}

	// indexSlashN := strings.Index(string(userHz), "\n")
	// userVal := string(userHz[:indexSlashN])
	// userVal = strings.ReplaceAll(userVal, "user ", "")
	// return strconv.Atoi(string(userVal))
	return parseCPUSTatFile(string(userHz))
}

func parseCPUSTatFile(data string) (int, error) {

	indexFirstNewLine := strings.Index(data, "\n")
	firstLine := strings.ReplaceAll(data[:indexFirstNewLine], "user ", "")

	user, err := strconv.Atoi(firstLine)
	if err != nil {
		return -1, err
	}

	return user, nil
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
