package executable_test

import (
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

	expected := " mkdir: cannot create directory ‘/bin/THIS_SHOULD_NOT_BE_HERE’: Read-only file system\n"
	assertEquals(expected, err.Error(), t)
}
