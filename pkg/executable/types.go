package executable

import "github.com/lkelly93/executable/internal/runners"

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
