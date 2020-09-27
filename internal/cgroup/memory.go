package cgroup

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func setupMemoryCGroup(cgroup, rootName string) error {
	memoryLoc := filepath.Join(cgroup, "memory", rootName)
	if err := os.MkdirAll(memoryLoc, 0755); err != nil {
		return err
	}

	fiveHundredMBInB := "500000000"
	err := setMemoryFileValue(memoryLoc, "memory.limit_in_bytes", fiveHundredMBInB)
	if err != nil {
		return err
	}

	err = setMemoryFileValue(memoryLoc, "memory.limit_in_bytes", fiveHundredMBInB)
	if err != nil {
		return err
	}

	return nil
}

func setMemoryFileValue(memoryLoc, file, value string) error {
	return ioutil.WriteFile(
		filepath.Join(
			memoryLoc,
			file,
		),
		[]byte(value),
		0700,
	)
}
