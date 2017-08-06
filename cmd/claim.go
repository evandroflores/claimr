package cmd

import (
	"fmt"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
)

func init() {
	Register("claim <vm-name>", "Claim a vm for your use", claim)
}

func claim(request *slacker.Request, response slacker.ResponseWriter) {
	vmName := request.Param("vm-name")
	channel := request.Event.Channel
	user := request.Event.User

	response.Typing()

	vm := model.VM{TeamID: request.Event.Team, Name: vmName}

	found, _ := database.DB.Get(&vm)
	if !found {
		response.Reply(fmt.Sprintf("I couldn't find vm `%s` on <#%s>.", vmName, channel))
	} else {
		if vm.InUseBy != "free" {
			response.Reply(fmt.Sprintf("VM `%s` is already in use, try another one.", vmName))
		} else {
			vm.InUseBy = user
			database.DB.Id(vm.ID).Update(&vm)
			response.Reply(fmt.Sprintf("Got it. VM `%s` is all yours <@%s>?", vmName, user))
		}
	}
}
