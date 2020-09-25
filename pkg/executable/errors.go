package executable

import "fmt"

//UnsupportedLanguageError is the error returned by NewExecutable if
//the language provided is not supported.
type UnsupportedLanguageError struct {
	Lang string
}

func (ule *UnsupportedLanguageError) Error() string {
	return fmt.Sprintf("%s is not a supported language", ule.Lang)
}

//Is Returns true if the given error is a UnsupportedLanguageError
func (ule *UnsupportedLanguageError) Is(err error) bool {
	_, ok := err.(*UnsupportedLanguageError)
	return ok
}

//CompilationError is the error returned by Run() if the provided source code
//could not compile.
type CompilationError struct {
	ErrMessage string
}

func (ce *CompilationError) Error() string {
	return fmt.Sprintf("Error, could not compile source code:\n %s", ce.ErrMessage)
}

//Is Returns true if the given error is a CompilationError
func (ce *CompilationError) Is(err error) bool {
	_, ok := err.(*CompilationError)
	return ok
}

//TimeLimitExceededError is returned if the max the executable took to long to run
type TimeLimitExceededError struct {
	MaxTime int
}

func (tle *TimeLimitExceededError) Error() string {
	return fmt.Sprintf("Time Limit Exceeded %ds", tle.MaxTime)
}

//Is Returns true if the given error is a TimeLimitExceededError
func (tle *TimeLimitExceededError) Is(err error) bool {
	_, ok := err.(*TimeLimitExceededError)
	return ok
}

//RuntimeError is returned if the executable had a runtime error during
//execution.
type RuntimeError struct {
	ErrMessage string
}

func (re *RuntimeError) Error() string {
	return re.ErrMessage
}

//Is Returns true if the given error is a RunTimeError
func (re *RuntimeError) Is(err error) bool {
	_, ok := err.(*RuntimeError)
	return ok
}

//SystemError is returned as a general wrapper for any undefined errors.
type SystemError struct {
	Err error
}

func (se *SystemError) Error() string {
	return "Error during execution of the program, check server logs."
}

//Is Returns true if the given error is a SystemError
func (se *SystemError) Is(err error) bool {
	_, ok := err.(*SystemError)
	return ok
}

func (se *SystemError) Unwrap() error { return se.Err }

type MalformedUniqueIdentifier struct {
	ErrMessage string
}

func (mui *MalformedUniqueIdentifier) Error() string {
	return mui.ErrMessage
}

//Is Returns true if the given error is a MalformedUniqueIdentifier
func (se *MalformedUniqueIdentifier) Is(err error) bool {
	_, ok := err.(*MalformedUniqueIdentifier)
	return ok
}
