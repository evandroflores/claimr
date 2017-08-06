package cmd

import "github.com/shomali11/slacker"

var commands []Command

// Command defines a command to be register to slack
type Command struct {
	Usage string
	Description string
	Handler func(request *slacker.Request, response slacker.ResponseWriter)
}

// Register add a command to commands list an prepare to register to slacker
func Register(usage string, description string, handler func(request *slacker.Request, response slacker.ResponseWriter)) {
	commands = append(commands, Command{Usage: usage, Description: description, Handler: handler})
}

// CommandList returns the list of registered commands
func CommandList() []Command {
	return commands
}

func notImplemented(request *slacker.Request, response slacker.ResponseWriter) {
	response.Reply("No pancakes for you! 🥞")
}