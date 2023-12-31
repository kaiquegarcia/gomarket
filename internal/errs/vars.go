package errs

import (
	"fmt"
	"gomarket/internal/enum"
	"strings"
)

var (
	SystemLoadingErr                  = fmt.Errorf("could not initialize the system")
	DependenciesLoadingErr            = fmt.Errorf("could not initialize dependencies")
	ApplicationNotPresentInContextErr = fmt.Errorf("the application is not present in current context")
	InvalidDataStoredInContextErr     = fmt.Errorf("the data stored in the current context doesn't have the expected format")
	RegistryNotFoundErr               = fmt.Errorf("registry not found")
	MaterialNotFoundErr               = fmt.Errorf("material not found")
	InvalidCommandErr                 = fmt.Errorf("invalid command")
	InvalidUnitValidationErr          = fmt.Errorf("invalid unit, please inform one of the following values as material unit: '%s'", strings.Join(enum.UnitKinds, "' / '"))
	MaterialCodeCantBeProductCodeErr  = fmt.Errorf("invalid material code, please inform a code different than the product's code")
)
