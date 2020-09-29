package executable_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/lkelly93/executable/pkg/executable"
)

func TestPutFileInBin(t *testing.T) {
	t.Parallel()
	lang := "python"
	code := "import os\nos.system(\"mkdir /bin/THIS_SHOULD_NOT_BE_HERE\")"
	uniqueIdentifier := "TestPutFileInBin"

	exe := getExecutable(lang, code, uniqueIdentifier, t)

	_, err := exe.Run()
	if err == nil {
		t.Fatalf("%s should of failed but it did not", uniqueIdentifier)
	}

	expectedError := &executable.RuntimeError{
		ErrMessage: "mkdir: cannot create directory ‘/bin/THIS_SHOULD_NOT_BE_HERE’: Read-only file system\n",
	}
	if !errors.Is(err, expectedError) {
		t.Errorf("Expected %T but got %T", expectedError, err)
	}
	assertEquals(expectedError.Error(), err.Error(), t)
}

func TestForkBomb(t *testing.T) {
	t.Parallel()
	lang := "python"
	code := "import os\nos.system(\"bomb() { bomb | bomb & }; bomb\")"
	uniqueIdentifier := "TestForkBomb"

	exe := getExecutable(lang, code, uniqueIdentifier, t)

	_, err := exe.Run()
	if err == nil {
		t.Fatalf("%s should of failed but it did not", uniqueIdentifier)
	}

	expectedError := &executable.RuntimeError{}
	if !errors.Is(err, expectedError) {
		t.Errorf("Expected %T but got %T", expectedError, err)
	}

	expectedSubstring := "Cannot fork"
	actual := err.Error()
	if !strings.Contains(actual, expectedSubstring) {
		t.Errorf(
			"Expected a string that contained %s but got %s",
			expectedSubstring,
			actual,
		)
	}
}

func TestCGroupVisibility(t *testing.T) {
	t.Parallel()
	lang := "python"
	code := "import os\nos.system(\"ls /sys/fs/cgroup\")"
	uniqueIdentifier := "TestCGroupVisibility"

	exe := getExecutable(lang, code, uniqueIdentifier, t)

	actual, err := exe.Run()
	if err != nil {
		t.Fatalf("Run() failed for %s, and should not have.", uniqueIdentifier)
	}

	expected := ""
	assertEquals(expected, actual.Output, t)
}
