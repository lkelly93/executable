// +build !longTests

package executable_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/lkelly93/executable/pkg/executable"
)

func TestInfiniteRecursion(t *testing.T) {
	t.Parallel()
	var code strings.Builder
	code.WriteString("public static void main(String[] args) {\n")
	code.WriteString("System.out.println(oops(5));\n")
	code.WriteString("}\n")
	code.WriteString("public static int oops(int x){\n")
	code.WriteString("if(x > 0 ){\n")
	code.WriteString("return oops(++x);\n")
	code.WriteString("}else{")
	code.WriteString("return x;\n")
	code.WriteString("}}")

	lang := "java"
	uniqueIdentifier := "TestInfiniteRecursion"
	exec, err := executable.NewExecutable(
		lang,
		code.String(),
		uniqueIdentifier,
	)
	if err != nil {
		t.Error(err)
	}

	_, err = exec.Run()

	if err == nil {
		t.Fatalf("%s Run did not produce an error.", uniqueIdentifier)
	}

	expectedError := &executable.RuntimeError{}
	if !errors.Is(err, expectedError) {
		t.Errorf("Expected %T but got %T", expectedError, err)
	}

	expectedOutput := "Exception in thread \"main\" java.lang.StackOverflowError\n"

	actual := err.Error()
	newLineIndex := strings.Index(actual, "\n") + 1
	actual = actual[:newLineIndex]

	assertEquals(expectedOutput, actual, t)
}

func TestInfiniteLoop(t *testing.T) {
	t.Parallel()
	var code strings.Builder
	code.WriteString("x = 5\n")
	code.WriteString("while(True):\n")
	code.WriteString("\tx+=1\n")

	lang := "python"
	uniqueIdentifier := "TestInfiniteLoop"
	exec, err := executable.NewExecutable(
		lang,
		code.String(),
		uniqueIdentifier,
	)

	if err != nil {
		t.Error(err)
	}

	_, err = exec.Run()

	expectedError := &executable.TimeLimitExceededError{
		MaxTime: executable.MaxExecutableRunTime,
	}
	if !errors.Is(err, expectedError) {
		t.Errorf("Expected %T but got %T", err, expectedError)
	}

	assertEquals(expectedError.Error(), err.Error(), t)
}
