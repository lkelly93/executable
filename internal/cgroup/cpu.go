package cgroup

import (
	"os"
	"path/filepath"
)

func setupCPUCGroup(cgroup, rootName string) error {
	cpuLoc := filepath.Join(cgroup, "cpuacct", rootName)
	if err := os.MkdirAll(cpuLoc, 0755); err != nil {
		return nil
	}

	return nil
}
