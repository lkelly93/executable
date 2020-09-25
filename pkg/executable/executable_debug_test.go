// +build !debug

package executable_test

import (
	"testing"
)

func TestDebug(t *testing.T) {
	assertEquals("Hello", "Hell", t)
	// lang := "python"
	// code := "print(\"Hello World"
	// uniqueIdentifier := "IntialTester"

	// exe, _ := executable.NewExecutable(lang, code, uniqueIdentifier)

	// _, err := exe.Run()

	// if err != nil {
	// t.Errorf("Error:%s\nError Type:%T\n", err.Error(), err)
	// var expectedErr = &executable.RuntimeError{}
	// t.Error(errors.Is(err, expectedErr))
	// if _, ok := err.(*executable.SystemError); ok {
	// 	t.Errorf("Server Logs:\n%s\n", errors.Unwrap(err).Error())
	// }
	// }

	// t.Errorf("Output:\n%s", out)
}
