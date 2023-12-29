package cmd

import (
	"fmt"
	"gomarket/internal/errs"
	"gomarket/internal/repository"
	"gomarket/internal/usecases/productcli"
	"gomarket/pkg/storage"
	"gomarket/pkg/util"
	"os"
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
	var command string
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch Command(command) {
	case ListProducts:
		app.productUsecases.List()
	case GetProduct:
		fmt.Println("unimplemented")
		util.FinishCLI()
	case CreateProduct:
		app.productUsecases.Create()
	case UpdateProduct:
		fmt.Println("unimplemented")
		util.FinishCLI()
	case DeleteProduct:
		fmt.Println("unimplemented")
		util.FinishCLI()
	default:
		fmt.Printf("invalid command. please send one of the following commands:\n- %s\n", strings.Join(availableCommands, "\n- "))
		util.FinishCLI()
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
