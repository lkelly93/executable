package environment

//EnvironmentData is what is returned from the SetupMethod. It holds the data
//needed to tear down the environment after it has been used.
type EnvironmentData struct {
	RootPath string
	rootName string
}
