package executable

import "fmt"

//UnsupportedLanguageError is the error returned by NewExecutable if
//the language provided is not supported.
type UnsupportedLanguageError struct {
	lang string
}

func (ule *UnsupportedLanguageError) Error() string {
	return fmt.Sprintf("%s is not a supported language", ule.lang)
}

//CompilationError is the error returned by Run() if the provided source code
//could not compile.
type CompilationError struct {
	errMessage string
}

func (ce *CompilationError) Error() string {
	return fmt.Sprintf("Error, could not compile source code:\n %s", ce.errMessage)
}

//TimeLimitExceededError is returned if the max the exectuable took to long to run
type TimeLimitExceededError struct {
	maxTime int
}

func (tle *TimeLimitExceededError) Error() string {
	return fmt.Sprintf("Time Limit Exceeded %ds", tle.maxTime)
}

//RuntimeError is returned if the executable had a runtime erorr during
//execution.
type RuntimeError struct {
	errMessage string
}

func (re *RuntimeError) Error() string {
	return re.errMessage
}

//SystemError is returned as a general wrapper for any undefined errors.
type SystemError struct{}

func (se *SystemError) Error() string {
	return "Error during execution of the program, check server logs"
}
