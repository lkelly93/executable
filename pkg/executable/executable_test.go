package executable_test

import (
	"errors"
	"testing"

	"github.com/lkelly93/pkg/executable"
)

func Test(t *testing.T) {
	lang := "python"
	code := "print(\"Hello World\")"
	settings := executable.Settings{
		UniqueIdentifier: "InitalTester",
	}

	exe, _ := executable.NewExecutable(lang, code, &settings)

	_, err := exe.Run()

	if err != nil {
		t.Error(err)
		if _, ok := err.(*executable.SystemError); ok {
			t.Errorf("Server Logs:\n%s\n", errors.Unwrap(err).Error())
		}
	}
}
