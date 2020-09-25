// +build !debug

package executable_test

import (
	"errors"
	"testing"

	"github.com/lkelly93/executable/pkg/executable"
)

func TestDebug(t *testing.T) {
	lang := "python"
	code := "print(\"Hello World"
	uniqueIdentifier := "IntialTester"

	exe, _ := executable.NewExecutable(lang, code, uniqueIdentifier)

	out, err := exe.Run()

	if err != nil {
		t.Errorf("Error:\n%s", err.Error())
		if _, ok := err.(*executable.SystemError); ok {
			t.Errorf("Server Logs:\n%s\n", errors.Unwrap(err).Error())
		}
	}

	t.Errorf("Output:\n%s", out)
}
