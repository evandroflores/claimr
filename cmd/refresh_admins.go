package cmd

import (
	"fmt"
	"os"

	"github.com/evandroflores/claimr/messages"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
)

func init() {
	Register("refresh-admins", "Reloads slack admins list. admin-only", refreshAdmins)
}

func refreshAdmins(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	event := getEvent(request)
	if !model.IsAdmin(event.User) {
		response.Reply(messages.Get("admin-only"))
		return
	}

	bot := slacker.NewClient(os.Getenv("CLAIMR_TOKEN")) // This is terrible!
	model.LoadAdmins(bot)
	response.Reply(fmt.Sprintf(messages.Get("x-admin-loaded"), len(model.Admins)))
}
