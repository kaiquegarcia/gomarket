package errs

import "fmt"

var (
	SystemLoadingErr                  = fmt.Errorf("could not initialize the system")
	DependenciesLoadingErr            = fmt.Errorf("could not initialize dependencies")
	ApplicationNotPresentInContextErr = fmt.Errorf("the application is not present in current context")
	InvalidDataStoredInContextErr     = fmt.Errorf("the data stored in the current context doesn't have the expected format")
	RegistryNotFoundErr               = fmt.Errorf("registry not found")
	InvalidCommandErr                 = fmt.Errorf("invalid command")
)
