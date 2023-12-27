package errs

import "fmt"

var (
	SystemLoadingErr                  = fmt.Errorf("could not initialize the system")
	ApplicationNotPresentInContextErr = fmt.Errorf("the application is not present in current context")
	InvalidDataStoredInContextErr     = fmt.Errorf("the data stored in the current context doesn't have the expected format")
)
