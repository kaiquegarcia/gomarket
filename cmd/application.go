package cmd

import (
	"fmt"
	"gomarket/internal/errs"
	"gomarket/internal/repository"
	"gomarket/internal/usecases/productcli"
	"gomarket/pkg/storage"
	"gomarket/pkg/util"
	"path/filepath"
	"runtime"
	"strings"
)

// Application has all the necessary data to run every command of this program
type Application interface {
	// RootDirectory returns the path to the project files, always ending the path with the FileSeparator
	RootDirectory() string
	// StorageDirectory returns the path to the storage folder in the project files, always ending the path with the FileSeparator
	StorageDirectory() string
	// Separator returns the path separator of the current operational system
	Separator() string
	// RunCLI will start the CLI procedures of this application
	RunCLI()
}

type application struct {
	rootDir         string
	separator       string
	productUsecases productcli.CLI
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

	app := &application{
		rootDir:   dirname,
		separator: separator,
	}

	app.loadDependencies()
	return app
}

func (app *application) RootDirectory() string {
	return app.rootDir
}

func (app *application) StorageDirectory() string {
	return app.rootDir + "storage" + app.separator
}

func (app *application) Separator() string {
	return app.separator
}

func (app *application) RunCLI() {
	fmt.Println("welcome to gomarket! your market manager made in golang :)")
	commandsList := " (" + strings.Join(availableCommands, "/") + ")"
	command := util.AskCLI("what do you want to do today?" + commandsList)
	for {
		switch Command(command) {
		case ListProducts:
			app.productUsecases.List()
			command = util.AskCLI("what do you want to do now?" + commandsList)
		case GetProduct:
			app.productUsecases.Get()
			command = util.AskCLI("what do you want to do now?" + commandsList)
		case CreateProduct:
			app.productUsecases.Create()
			command = util.AskCLI("what do you want to do now?" + commandsList)
		case UpdateProduct:
			command = util.AskCLI("this command is not implemented yet, can you try something else?" + commandsList)
		case DeleteProduct:
			command = util.AskCLI("this command is not implemented yet, can you try something else?" + commandsList)
		case Exit:
			fmt.Println("ok! bye bye")
			util.FinishCLI()
			return
		default:
			command = util.AskCLI(
				fmt.Sprintf("invalid command. please send one of the following commands:\n- %s", strings.Join(availableCommands, "\n- ")),
			)
		}
	}
}

func (app *application) loadDependencies() {
	js := storage.NewJsonStorage(app.StorageDirectory())
	// Collections
	productCollection, err := storage.NewCollection(js, "products")
	if err != nil {
		panic(errs.DependenciesLoadingErr)
	}

	// Repositories
	productRepository := repository.NewProductRepository(productCollection)

	// Usecases
	app.productUsecases = productcli.NewCLI(productRepository)
}
