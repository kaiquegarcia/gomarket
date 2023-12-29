package cmd

type Command string

const (
	ListProducts  Command = "list"
	GetProduct    Command = "get"
	CreateProduct Command = "create"
	UpdateProduct Command = "update"
	DeleteProduct Command = "delete"
	Exit          Command = "exit"
)

var availableCommands = []string{
	string(ListProducts),
	string(GetProduct),
	string(CreateProduct),
	string(UpdateProduct),
	string(DeleteProduct),
	string(Exit),
}
