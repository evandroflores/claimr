package cmd

import (
	"fmt"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
	Register("list <filter>", "List all VMs or filtered by public, private or available", list)
}

func list(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	vms := make([]model.VM, 0)
	channelVMs := model.VM{TeamID: request.Event.Team, ChannelID: request.Event.Channel}

	err := database.DB.Find(&vms, &channelVMs)

	if err != nil {
		response.Reply("Fail to list VMs")
		log.Error(err)
		return
	}

	if len(vms) == 0 {
		response.Reply("No VMs to list")
		return
	}

	text := "Here is a list of all VMs\n---\n"
	for _, vm := range vms {
		line := fmt.Sprintf("`%s` \t", vm.Name)
		if vm.InUseBy == "free" {
			line += "_available_"
		} else {
			line += "in use"
		}
		text += fmt.Sprintf("%s\n", line)
	}
	response.Reply(text)
}
