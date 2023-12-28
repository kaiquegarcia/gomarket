package cmd

import (
	"gomarket/internal/errs"
	"path/filepath"
	"runtime"
	"strings"
)

// Application has all the necessary data to run every command of this program
type Application interface {
	// RootDirectory returns the path to the project files, always ending the path with the FileSeparator
	RootDirectory() string
	// Separator returns the path separator of the current operational system
	Separator() string
	// RunCLI will start the CLI procedures of this application
	RunCLI()
}

type application struct {
	rootDir   string
	separator string
}

// NewApp initializes an implementation of Application interface
func NewApp() Application {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic(errs.SystemLoadingErr)
	}

	separator := string(filepath.Separator)
	dirname := filepath.Dir(filepath.Dir(filename) + separator + ".." + separator)
	if !strings.HasPrefix(dirname, separator) {
		dirname += separator
	}

	return application{
		rootDir:   dirname,
		separator: separator,
	}
}

func (app application) RootDirectory() string {
	return app.rootDir
}

func (app application) Separator() string {
	return app.separator
}

func (app application) RunCLI() {
	// TODO
}
