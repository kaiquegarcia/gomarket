package main

import (
	"fmt"
	"gomarket/cmd"
	"os"
)

const (
	exitWithSuccess = iota
	exitWithGenericError
)

func main() {
	app := cmd.NewApp()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("panic triggered:\n> %s\n", err)
			os.Exit(exitWithGenericError)
		}

		os.Exit(exitWithSuccess)
	}()

	app.RunCLI()
}
