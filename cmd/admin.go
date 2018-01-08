package cmd

import (
	"os"
	"strings"

	adm "github.com/evandroflores/claimr/cmd/admin"
	"github.com/shomali11/slacker"
)

func init() {
	Register("admin <sub-command> <sub-command-parameter>", "Administrative set of commands. Available only for admins.", admin)
}

func admin(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	event := getEvent(request)
	if event.User != os.Getenv("CLAIMR_SUPERUSER") {
		response.Reply("Command available only for admins. 🛑")
		return
	}

	adminSubCommand := request.Param("sub-command")
	found := false
	for _, subcommand := range adm.CommandList() {
		subcommandName := strings.Split(subcommand.Usage, " ")[0]
		if subcommandName == adminSubCommand {
			found = true
			subcommand.Handler(request, response)
		}
	}
	if !found {
		response.Reply("Command not found, Type `@claimr admin command-list` for valid admin sub commands.")

	}
}
