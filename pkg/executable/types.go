package executable

import "github.com/lkelly93/executable/internal/runners"

const (
	//MaxExecutableRunTime is the maximum time that a program is given.
	//If a program takes longer then this to run then the program will be killed.
	MaxExecutableRunTime = 15
)

//Executable represents programs that are ready to execute
type Executable interface {
	Run() (string, error)
}

type executableState struct {
	code             string
	lang             string
	uniqueIdentifier string
	runner           runners.Runner
}
