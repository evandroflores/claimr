package cmd

import (
	"fmt"
	"github.com/shomali11/slacker"
	 "github.com/evandroflores/claimr/model"
	 "github.com/evandroflores/claimr/database"
)

func init(){
    Register("free <vm-name>", "Free a vm from use", free)
}

func free(request *slacker.Request, response slacker.ResponseWriter) {
	vmName := request.Param("vm-name")
	channel := request.Event.Channel
	user := request.Event.User

	response.Typing()

	vm := model.VM{TeamID: request.Event.Team, Name: vmName}

	found, _ := database.DB.Get(&vm)
	if !found {
		response.Reply(fmt.Sprintf("I couldn't find vm `%s` on <#%s>.", vmName, channel))
	} else {
		if vm.InUseBy != user {
			response.Reply(fmt.Sprintf("Humm VM `%s` is not being used by you.", vmName))
		} else {
			vm.InUseBy = "free"
			database.DB.Id(vm.ID).Update(&vm)
			response.Reply(fmt.Sprintf("Got it. VM `%s` is now available", vmName))
		}
	}
}