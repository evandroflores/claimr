package cmd

import (
	"fmt"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
	"strings"
)

var (
	maxNameSize         = 22
	directChannelPrefix = "D"
)

func init() {
	Register("add <vm-name>", "Add a vm to your channel", add)
}

func add(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	if strings.HasPrefix(request.Event.Channel, directChannelPrefix) {
		response.Reply("This look like a direct message. VMs are related to a channel.")
		return
	}

	vmName := request.Param("vm-name")
	if len(vmName) > maxNameSize {
		response.Reply(fmt.Sprintf("Try a name smaller than %d", maxNameSize))
		return
	}
	if len(vmName) == 0 {
		response.Reply("Please, give a name for your vm 🤦")
		return
	}

	found, _ := database.DB.Get(&model.VM{TeamID: request.Event.Team, Name: vmName})

	if found {
		response.Reply("There is a VM with the same name on this channel. Try a different one.")
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
		log.Error(err)
	}
	if affected == 1 {
		response.Reply(fmt.Sprintf("VM `%s` added to channel <#%s>", vm.Name, vm.ChannelID))
		log.Debugf("VM %s added to channel %s", vm.Name, vm.ChannelID)
	} else {
		log.Errorf("Tried to add vm %s but failed.", vm.Name)
		response.Reply("This doesn't smells good")
	}
}
