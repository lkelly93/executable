// +build !debug

package executable_test

import (
	"errors"
	"testing"

	"github.com/lkelly93/executable/pkg/executable"
)

func TestDebug(t *testing.T) {
	lang := "python"
	// code := "import os\nos.system(\"bomb() { bomb | bomb & }; bomb\")"
	// code := "import time\ntime.sleep(13)"
	// code := "import os\nos.system(\"cat /sys/fs/cgroup/pids/IntialTester/cgroup.procs\")"
	code := "import os\nos.system(\"print(\"Hello\")"
	// code := "import os\nos.system(\"cat /proc/mounts\")"
	uniqueIdentifier := "IntialTester"

	exe, _ := executable.NewExecutable(lang, code, uniqueIdentifier)

	out, err := exe.Run()

	if err != nil {
		t.Errorf("Error:%s\nError Type:%T\n", err.Error(), err)
		if errors.Is(err, &executable.SystemError{}) {
			t.Errorf("Server Logs:\n%s\n", errors.Unwrap(err).Error())
		}
	}

	if len(out) != 0 {
		t.Errorf("Output:\n%s", out)
	}
}
