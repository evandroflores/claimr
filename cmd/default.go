package cmd

import "github.com/evandroflores/slacker"

//Default command will be called when a command is not recognized.
func Default(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	response.Reply("Not sure what you are asking for. Type `@claimr help` for valid commands.")
}
