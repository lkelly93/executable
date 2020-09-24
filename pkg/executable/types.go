package executable

//Executable represents programs that are ready to execute
type Executable interface {
	Run() (string, error)
}

//Settings represents all the settings for a given executable.
type Settings struct {
	Imports      string
	ClassName    string
	TrailingCode string
	//UniqueIdentifier is used to prevent data races. This package will not
	//check if the Identifier is unique that is on the user of the package,
	//an unique identifier can cause data races with unknown results.
	UniqueIdentifier string
}

type executableState struct {
	code        string
	lang        string
	settings    *Settings
	fileCreator fileCreationFunction
}

type fileCreationFunction func(string, *Settings) (string, string, error)

var supportedLanguages = map[string]fileCreationFunction{
	// "python": createRunnerFilePython,
}
