package cmd

import (
	"gomarket/cmd/http/product"
	"gomarket/internal/errs"
	"gomarket/internal/repository"
	"gomarket/internal/usecases/productcli"
	"gomarket/internal/usecases/producthttp"
	"gomarket/pkg/storage"
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
	// RunWeb will start the API procedures of this application, serving it immediately with graceful shutdown
	RunWeb()
}

type application struct {
	rootDir             string
	separator           string
	productUsecasesCLI  productcli.CLI
	productUsecasesHTTP producthttp.HTTP
	productHandlers     product.Handlers
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

	app.loadDirectories()
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

func (app *application) loadDirectories() {
	app.loadStorageDirectory()
}

func (app *application) loadStorageDirectory() {
	storageDir := app.StorageDirectory()
	_, err := os.ReadDir(storageDir)
	if err == nil {
		return
	}

	if os.IsNotExist(err) {
		err = os.Mkdir(storageDir, os.ModeDir)
		if err == nil {
			return
		}

		panic("could not create storage directory: " + err.Error())
	}

	panic("could not detect storage directory: " + err.Error())
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
	app.productUsecasesCLI = productcli.NewCLI(productRepository)
	app.productUsecasesHTTP = producthttp.NewHTTP(productRepository)

	// HTTP Handlers
	app.productHandlers = product.NewHandlers(app.productUsecasesHTTP)
}
