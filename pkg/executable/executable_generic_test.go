package executable_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lkelly93/executable/pkg/executable"
)

func TestFileSystemAreDeletedAfterRun(t *testing.T) {
	t.Parallel()
	lang := "python"
	code := "print(\"Hello World!\")"
	uniqueIdentifier := "TestFileSystemAreDeletedAfterRun"

	exe := getExecutable(lang, code, uniqueIdentifier, t)

	fileLocation := filepath.Join("/securefs", uniqueIdentifier)

	_, err := os.Stat(fileLocation)
	if err == nil {
		t.Errorf("%s existed before Run() was called", fileLocation)
	}

	exe.Run()

	_, err = os.Stat(fileLocation)
	if err == nil {
		t.Errorf("%s still exists after Run() was called. It should of been deleted", fileLocation)
	}

}
func TestCGroupFilesAreDeletedAfterRun(t *testing.T) {
	t.Parallel()
	lang := "python"
	code := "print(\"Hello World!\")"
	uniqueIdentifier := "TestCGroupFilesAreDeletedAfterRun"

	exe := getExecutable(lang, code, uniqueIdentifier, t)

	fileLocation := filepath.Join("/sys/fs/cgroups/pids", uniqueIdentifier)
	_, err := os.Stat(fileLocation)
	if err == nil {
		t.Errorf("%s existed before Run() was called", fileLocation)
	}

	exe.Run()

	_, err = os.Stat(fileLocation)
	if err == nil {
		t.Errorf("%s still exists after Run() was called. It should of been deleted", fileLocation)
	}

}

func assertEquals(expected, actual string, t *testing.T) {
	t.Helper()
	problemIndex := findError(expected, actual)
	if problemIndex > 0 {
		expectedChar := getChar(expected, problemIndex)
		actualChar := getChar(actual, problemIndex)
		t.Errorf("Expected:\n\"%s\"\nbut got\n\"%s\"", expected, actual)
		t.Errorf(
			"Error at index %d, expected %c but was %c",
			problemIndex,
			expectedChar,
			actualChar,
		)
	}
}

func findError(expected, actual string) int {
	lenExpected := len(expected)
	lenActual := len(actual)

	if lenActual != lenExpected {
		return min(lenActual, lenExpected)
	}

	for i := 0; i < lenExpected && i < lenActual; i++ {
		if expected[i] != actual[i] {
			return i
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getChar(s string, index int) byte {
	if index >= len(s) {
		return 0
	}

	return s[index]
}

func getExecutable(lang, code, uniqueIdentifier string, t *testing.T) executable.Executable {
	t.Helper()
	exe, err := executable.NewExecutable(lang, code, uniqueIdentifier)
	if err != nil {
		t.Errorf("NewExecutable failed returned an error with the given language %s", lang)
	}
	return exe
}
