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

	exe, err := executable.NewExecutable(lang, code, uniqueIdentifier)
	if err != nil {
		t.Errorf("NewExecutable failed with the supported language %s", lang)
	}

	_, err = exe.Run()
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

	exe, err := executable.NewExecutable(lang, code, uniqueIdentifier)
	if err != nil {
		t.Errorf("NewExecutable failed with the supported language %s", lang)
	}

	_, err = exe.Run()
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
