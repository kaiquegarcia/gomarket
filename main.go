package main

import (
	"fmt"
	"gomarket/cmd"
	"gomarket/pkg/util"
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
			util.FinishCLI()
			os.Exit(exitWithGenericError)
		}

		os.Exit(exitWithSuccess)
	}()

	if len(os.Args) > 1 && os.Args[1] == "serve" {
		app.RunWeb()
	} else {
		app.RunCLI()
	}
}
