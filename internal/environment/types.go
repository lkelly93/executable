package environment

//EnvironmentData is what is returned from the SetupMethod. It holds the data
//needed to tear down the environment after it has been used.
type EnvironmentData struct {
	RootPath string
	rootName string
}

//UsedCGroups is all of the cgroups that this package setups up.
//This array can be helpful for teardown an environment.
var UsedCGroups = []string{
	"pids",
	"memory",
}
