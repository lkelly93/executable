package executable_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lkelly93/executable/pkg/executable"
)

func TestFileIsDeletedAfterRun(t *testing.T) {
	t.Parallel()
	lang := "python"
	code := "print(\"Hello World!\")"
	uniqueIdentifier := "TestFileIsDeletedAfterRun"

	exe, err := executable.NewExecutable(lang, code, uniqueIdentifier)
	if err != nil {
		t.Errorf("NewExecutable returned on error with the supported language %s", lang)
	}

	fileLocation := filepath.Join("/securefs", uniqueIdentifier)

	_, err = os.Stat(fileLocation)
	if err == nil {
		t.Errorf("%s existed before Run() was called", fileLocation)
	}

	exe.Run()

	_, err = os.Stat(fileLocation)
	if err == nil {
		t.Errorf("%s still exists after Run() was called. It should of been deleted", fileLocation)
	}

}

func assertEquals(expected string, actual string, t *testing.T) {
	t.Helper()
	if actual != expected {
		i := 0
		var expectedChar byte
		var actualChar byte
		for ; i < len(expected) && i < len(actual); i++ {
			if expected[i] != actual[i] {
				expectedChar = expected[i]
				actualChar = actual[i]
				break
			}
		}
		t.Errorf("Expected:\n\"%s\"\nbut got\n\"%s\"", expected, actual)
		t.Errorf("Error at index %d, expected %c but was %c", i, expectedChar, actualChar)
	}
}
