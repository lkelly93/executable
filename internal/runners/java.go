package runners

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
)

//pythonRunner creates the runner file that will be used to run
//the executable.
//It takes in the code that you want to run, and the destination where you
//want to put the runner file and returns the command to run the file and
//the name of the file or any errors if it encounters them during the creation
//process.
func (runner *javaRunner) CreateFile(code string, destination string) (string, string, error) {
	outFileName := filepath.Join(destination, runner.className)

	var codeFormatter bytes.Buffer
	runner.insertHeaderCode(&codeFormatter)
	codeFormatter.WriteString(code)
	runner.insertTrailingCode(&codeFormatter)

	err := ioutil.WriteFile(outFileName, codeFormatter.Bytes(), 0755)
	if err != nil {
		return "", "", err
	}
	return runner.langCommand, runner.className, nil
}

func (runner *javaRunner) insertHeaderCode(codeFormatter *bytes.Buffer) {
	codeFormatter.WriteString("import java.util.*;\n")
	codeFormatter.WriteString("import java.lang.*;\n")
	codeFormatter.WriteString("public class JavaRunner{\n")
}

func (runner *javaRunner) insertTrailingCode(codeFormatter *bytes.Buffer) {
	codeFormatter.WriteString("\n}")
}
