package cmd

import (
	"github.com/shomali11/slacker"
)

func init() {
	Register("admin <sub-command> <sub-command-parameter>", "Administrative set of commands. Available only for admins.", admin)
}

func admin(request *slacker.Request, response slacker.ResponseWriter) {
	response.Reply("Command available only for admins. 🛑")
}
