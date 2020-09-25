package runners

//Runner is an interface that can implement the creation of runner files
type Runner interface {
	CreateFile(string, string) (string, string, error)
}

type pythonRunner struct {
	langCommand string
	className   string
}

type javaRunner struct {
	langCommand string
	className   string
}

var supportedLanguages = map[string]Runner{
	"python": &pythonRunner{
		langCommand: "python3",
		className:   "PythonRunner.py",
	},
	"java": &javaRunner{
		langCommand: "java",
		className:   "JavaRunner.java",
	},
}
