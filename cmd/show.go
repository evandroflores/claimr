package cmd

import (
	"fmt"
	"time"
	"github.com/shomali11/slacker"
	 "github.com/evandroflores/claimr/model"
	 "github.com/evandroflores/claimr/database"
)

func init(){
    Register("show <vm-name>", "Show who is using the vm", show)
}

func show(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	vmName := request.Param("vm-name")
	channel := request.Event.Channel
	vm := model.VM{TeamID: request.Event.Team, Name: vmName}

	found, _ := database.DB.Get(&vm)
	if !found {
		response.Reply(fmt.Sprintf("I couldn't find vm *%s* on <#%s>.", vmName, channel))
	} else {
		text := fmt.Sprintf("VM `%s`. Created by <@%s>\n", vmName, vm.CreatedByUser)

		if vm.InUseBy == "free" {
			text += "_Available_"
		} else {
			text += fmt.Sprintf("Being used by <@%s>", vm.InUseBy)
		}
		text += fmt.Sprintf(" since %s.", vm.UpdatedAt.Format(time.RFC1123))
		response.Reply(text)
	}
}
