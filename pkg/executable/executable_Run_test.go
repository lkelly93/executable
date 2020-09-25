package executable_test

import (
	"testing"

	"github.com/lkelly93/executable/pkg/executable"
)

func TestRunPythonCode(t *testing.T) {
	t.Parallel()
	lang := "python"
	code := "print(\"Hello World\")"
	uniqueIdentifier := "TestRunPythonCode"

	exe, err := executable.NewExecutable(lang, code, uniqueIdentifier)
	if err != nil {
		t.Errorf("NewExecutable failed with the supported language %s", lang)
	}

	out, err := exe.Run()

	if err != nil {
		t.Errorf("Run() failed with the code %s and it should not have. -- Error:%s", code, err.Error())
	}

	expected := "Hello World\n"
	assertEquals(out, expected, t)
}

func TestRunJavaCode(t *testing.T) {
	t.Parallel()
	lang := "java"
	code := "public static void main(String[] args){System.out.println(\"Hello World\");}"
	uniqueIdentifier := "TestRunJavaCode"

	exe, err := executable.NewExecutable(lang, code, uniqueIdentifier)
	if err != nil {
		t.Errorf("NewExecutable failed with the supported language %s", lang)
	}

	out, err := exe.Run()

	if err != nil {
		t.Errorf("Run() failed with the code %s and it should not have. -- Error:%s", code, err.Error())
	}

	expected := "Hello World\n"
	assertEquals(out, expected, t)
}
