package cmd

import (
	"github.com/evandroflores/claimr/messages"
	"github.com/shomali11/slacker"
)

// Default command will be called when a command is not recognized.
func Default(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	response.Reply(messages.Messages["command-not-found"])
}
