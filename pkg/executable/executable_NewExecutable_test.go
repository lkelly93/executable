package executable_test

import (
	"errors"
	"testing"

	"github.com/lkelly93/executable/pkg/executable"
)

func TestNewExecutable(t *testing.T) {
	t.Parallel()
	lang := "python"
	code := "print(\"Hello World!\")"
	uniqueIdentifier := "TestNewExecutable"
	exe, err := executable.NewExecutable(lang, code, uniqueIdentifier)
	if err != nil {
		t.Errorf("NewExecutable failed with a supported language.")
	}

	if _, ok := interface{}(exe).(executable.Executable); !ok {
		t.Errorf("Expected an *executable.executableState but got %T", exe)
	}
}

func TestNewExecutableFail(t *testing.T) {
	t.Parallel()
	lang := "Not a Language"
	code := "Not Code"
	uniqueIdentifier := "TestNewExecutableFail"
	_, err := executable.NewExecutable(lang, code, uniqueIdentifier)
	if err == nil {
		t.Errorf("\"%s\" was accepted as a supported langauge and should not be.", lang)
	} else {
		expected := &executable.UnsupportedLanguageError{}
		if !errors.Is(err, expected) {
			t.Errorf("Expected %T but got %T", expected, err)
		}
	}
}

func TestEmptyIdentifier(t *testing.T) {
	t.Parallel()
	lang := "python"
	code := "print(\"Hello World\")"

	_, err := executable.NewExecutable(lang, code, "")

	if err == nil {
		t.Errorf("NewExecutable accepted an empty string as a uniqueIdentifier.")
	}

	expected := &executable.MalformedUniqueIdentifier{}
	if !errors.Is(err, expected) {
		t.Errorf("Expected %T but got %T", expected, err)
	}
}
