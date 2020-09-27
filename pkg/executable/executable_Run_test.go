package executable_test

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lkelly93/executable/pkg/executable"
)

type TestRunData struct {
	lang             string
	code             string
	uniqueIdentifier string
	expected         string
	expectedError    error
}

func TestRunPythonCode(t *testing.T) {
	test := TestRunData{
		lang:             "python",
		code:             "print(\"Hello World\")",
		uniqueIdentifier: "TestRunPythonCode",
		expected:         "Hello World\n",
	}
	test.run(t)
}

func TestRunJavaCode(t *testing.T) {
	test := TestRunData{
		lang:             "java",
		code:             "public static void main(String[] args){System.out.println(\"Hello World\");}",
		uniqueIdentifier: "TestRunJavaCode",
		expected:         "Hello World\n",
	}
	test.run(t)
}

func TestRunPythonCodeLargerFile(t *testing.T) {
	test := TestRunData{
		lang:             "python",
		code:             getFileData("longPythonCode.py", t),
		uniqueIdentifier: "TestRunPythonCodeLargerFile",
		expected:         "Male\n",
	}
	test.run(t)
}

func TestRunJavaCodeLargerFile(t *testing.T) {
	test := TestRunData{
		lang:             "java",
		code:             getFileData("longJavaCode.java", t),
		uniqueIdentifier: "TestRunJavaCodeLargerFile",
		expected:         "NonRecursive\n[0, 1, 0, 0, 1, 0, 1, 0]\n[0, 0, 0, 0, 0, 1, 1, 0]\n",
	}
	test.run(t)
}

func TestRecursion(t *testing.T) {
	test := TestRunData{
		lang:             "java",
		code:             getFileData("recursiveCode.java", t),
		uniqueIdentifier: "TestRecursion",
		expected:         "Recursive\n[0, 1, 0, 0, 1, 0, 1, 0]\n[0, 0, 0, 0, 0, 1, 1, 0]\n",
	}
	test.run(t)
}

func TestRunBadJavaCode(t *testing.T) {
	var expected strings.Builder
	expected.WriteString("JavaRunner.java:4: error: ';' expected\n")
	expected.WriteString("public static void main(String[] args){System.out.println(\"Hello World\")\n")
	expected.WriteString("                                                                        ^\n")
	expected.WriteString("JavaRunner.java:5: error: reached end of file while parsing\n")
	expected.WriteString("}\n")
	expected.WriteString(" ^\n")
	expected.WriteString("2 errors\n")
	expected.WriteString("error: compilation failed\n")

	test := TestRunData{
		lang:             "java",
		code:             "public static void main(String[] args){System.out.println(\"Hello World\")",
		uniqueIdentifier: "TestRunBadJavaCode",
		expected:         "",
		expectedError: &executable.RuntimeError{
			ErrMessage: expected.String(),
		},
	}

	test.run(t)
}

func TestRunBadPythonCode(t *testing.T) {
	var expected strings.Builder
	expected.WriteString("  File \"PythonRunner.py\", line 2\n")
	expected.WriteString("    print(\"Hi\n")
	expected.WriteString("            ^\n")
	expected.WriteString("SyntaxError: EOL while scanning string literal\n")

	test := TestRunData{
		lang:             "python",
		code:             "print(\"Hi",
		uniqueIdentifier: "TestRunBadPythonCode",
		expected:         "",
		expectedError: &executable.RuntimeError{
			ErrMessage: expected.String(),
		},
	}

	test.run(t)

}

func (data *TestRunData) run(t *testing.T) {
	t.Parallel()
	t.Helper()

	exe := getExecutable(data.lang, data.code, data.uniqueIdentifier, t)

	out, err := exe.Run()
	if err != nil {
		if data.expectedError != nil {
			if !errors.Is(err, data.expectedError) {
				t.Errorf("Expected %T but got %T", data.expectedError, err)
				t.Fatalf("Error:%s", err.Error())
			}
			assertEquals(data.expectedError.Error(), err.Error(), t)
		} else {
			t.Errorf("Run() failed with the code %s and it should not have. -- Error:%s", data.code, err.Error())
		}
	}
	assertEquals(out, data.expected, t)
}

func getFileData(fileName string, t *testing.T) string {
	fileLoc := filepath.Join("test_data", fileName)
	bytes, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		t.Errorf("Could not open %s.", fileName)
	}

	return string(bytes)
}
