package cmd

import (
	"fmt"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	"time"
)

func init() {
	Register("rm <vm-name>", "Remove a vm from your channel", remove)
}

func remove(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	vmName := request.Param("vm-name")

	if len(vmName) == 0 {
		response.Reply("Give me a vm name to remove 🙄")
		return
	}
	channel := request.Event.Channel
	user := request.Event.User

	vm := model.VM{TeamID: request.Event.Team, Name: vmName}

	found, _ := database.DB.Get(&vm)
	if !found {
		response.Reply(fmt.Sprintf("I couldn't find vm *%s* on <#%s>.", vmName, channel))
	} else {
		if vm.CreatedByUser != user {
			response.Reply(fmt.Sprintf("Only who created can remove a VM. Please check with <@%s>", vm.CreatedByUser))
		} else {
			if vm.InUseBy != "free" {
				response.Reply(fmt.Sprintf("Can't remove. VM *%s* is being used by <@%s> since %s.", vmName, vm.InUseBy, vm.UpdatedAt.Format(time.RFC1123)))
			} else {
				affected, _ := database.DB.ID(vm.ID).Delete(&vm)
				if affected == 1 {
					response.Reply(fmt.Sprintf("VM *%s* removed.", vmName))
				} else {
					response.Reply("This doesn't smells good")
				}

			}
		}
	}
}
