package environment

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func setupAllCGroupsFS(rootName string) error {
	cgroup := "/sys/fs/cgroup"
	if err := setupPidsCgroup(cgroup, rootName); err != nil {
		return err
	}

	return nil
}

func setupPidsCgroup(cgroup, rootName string) error {
	pidsLoc := filepath.Join(cgroup, "pids", rootName)
	if err := os.MkdirAll(pidsLoc, 0755); err != nil {
		return err
	}

	err := ioutil.WriteFile(filepath.Join(pidsLoc, "pids.max"), []byte("75"), 0700)
	if err != nil {
		return err
	}

	_, err = os.Create(filepath.Join(pidsLoc, "cgroup.procs"))
	if err != nil {
		return err
	}

	return nil
}

func tearDownAllCgroupFS(rootName string) error {
	cgroup := "/sys/fs/cgroup"
	if err := os.RemoveAll(filepath.Join(cgroup, "pids", rootName)); err != nil {
		return err
	}

	return nil
}
