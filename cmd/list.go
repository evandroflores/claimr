package cmd

import (
	"fmt"
	"log"
	"github.com/shomali11/slacker"
	 "github.com/evandroflores/claimr/model"
	 "github.com/evandroflores/claimr/database"
)

func init(){
    Register("list <filter>", "List all VMs or filtered by public, private or available", list)
}

func list(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	vms := make([]model.VM, 0)
	channelVMs := model.VM{TeamID: request.Event.Team, ChannelID: request.Event.Channel}
	publicVMs := model.VM{TeamID: request.Event.Team, ChannelID: "*"}

	err := database.DB.Find(&vms, &channelVMs)
	err = database.DB.Find(&vms, &publicVMs)

	if err != nil {
		log.Println(err)
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