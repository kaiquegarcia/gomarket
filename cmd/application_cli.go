package cmd

import (
	"fmt"
	"gomarket/pkg/util"
	"strings"
)

func (app *application) RunCLI() {
	fmt.Println("welcome to gomarket! your market manager made in golang :)")
	commandsList := " (" + strings.Join(availableCommands, "/") + ")"
	command := util.AskCLI("what do you want to do today?" + commandsList)
	for {
		nextCall := "what do you want to do now?" + commandsList
		switch Command(command) {
		case ListProducts:
			app.productUsecasesCLI.List()
		case GetProduct:
			app.productUsecasesCLI.Get()
		case CreateProduct:
			app.productUsecasesCLI.Create()
		case UpdateProduct:
			app.productUsecasesCLI.Update()
		case DeleteProduct:
			app.productUsecasesCLI.Delete()
		case RunWeb:
			app.RunWeb()
			return
		case Exit:
			fmt.Println("ok! bye bye...")
			return
		default:
			nextCall = fmt.Sprintf(
				"invalid command. please send one of the following commands:\n- %s",
				strings.Join(availableCommands, "\n- "),
			)
		}

		command = util.AskCLI(nextCall)
	}
}
