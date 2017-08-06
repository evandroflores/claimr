package cmd

import (
	"fmt"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	"log"
)

var (
	maxNameSize = 22
)

func init() {
	Register("add <vm-name>", "Add a vm to your channel", add)
}

func add(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	vmName := request.Param("vm-name")
	if len(vmName) > maxNameSize {
		response.Reply(fmt.Sprintf("Try a name smaller than %d", maxNameSize))
		return
	}
	if len(vmName) == 0 {
		response.Reply("Please, give a name for your vm 🤦")
		return
	}

	vm := new(model.VM)
	vm.Name = vmName
	vm.TeamID = request.Event.Team
	vm.ChannelID = request.Event.Channel
	vm.InUseBy = "free"
	vm.CreatedByUser = request.Event.User

	affected, err := database.DB.Insert(vm)

	if err != nil {
		log.Println(err)
	}
	if affected == 1 {
		response.Reply(fmt.Sprintf("VM `%s` added to channel <#%s>", vm.Name, vm.ChannelID))
	} else {
		response.Reply("This doesn't smells good")
	}
}
