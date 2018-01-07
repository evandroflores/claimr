package admin

import (
	"github.com/shomali11/slacker"
)

var adminCommands []SubCommand

// SubCommand is a second level of command
type SubCommand struct {
	Usage       string
	Description string
	Handler     func(request *slacker.Request, response slacker.ResponseWriter)
}

// Register add a command to commands list an prepare to register to slacker
func Register(usage string, description string, handler func(request *slacker.Request, response slacker.ResponseWriter)) {
	adminCommands = append(adminCommands, SubCommand{Usage: usage, Description: description, Handler: handler})
}

// CommandList returns the list of registered admin sub commands
func CommandList() []SubCommand {
	return adminCommands
}
