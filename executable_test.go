package executable_test

import (
	"testing"

	"github.com/lkelly93/executable"
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
	}
}
