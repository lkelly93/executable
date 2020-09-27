package environment

import "github.com/lkelly93/executable/internal/cgroup"

//Setup sets up the running environment for a given executable
func Setup(rootName string) (*EnvironmentData, error) {
	ed := &EnvironmentData{rootName: rootName}
	rootPath, err := setupRunnerFileSystem(rootName)
	ed.RootPath = rootPath
	if err != nil {
		return ed, err
	}

	if err = bindAndCopyRequiredFiles(rootPath); err != nil {
		return ed, err
	}

	if err = cgroup.Setup(rootName); err != nil {
		return ed, err
	}

	return ed, nil
}

//CleanUp removes the filesystem and binds that were setup for an executable.
func (ed *EnvironmentData) CleanUp() error {
	err := unbindAll(ed.RootPath)
	if err != nil {
		return err
	}

	err = removeRunnerFileSystem(ed.RootPath)
	if err != nil {
		return err
	}

	err = cgroup.Cleanup(ed.rootName)
	if err != nil {
		return err
	}

	return nil
}
