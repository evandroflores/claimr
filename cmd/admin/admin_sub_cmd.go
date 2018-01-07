package admin

import (
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
)

var adminCommands []model.Command

// Register add a command to commands list an prepare to register to slacker
func Register(usage string, description string, handler func(request *slacker.Request, response slacker.ResponseWriter)) {
	adminCommands = append(adminCommands, model.Command{Usage: usage, Description: description, Handler: handler})
}

// CommandList returns the list of registered admin sub commands
func CommandList() []model.Command {
	return adminCommands
}
