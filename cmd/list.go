package cmd

import (
	"fmt"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
	Register("list", "List all containers.", list)
}

func list(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	containers := make([]model.Container, 0)
	channelContainers := model.Container{TeamID: request.Event.Team, ChannelID: request.Event.Channel}

	err := database.DB.Find(&containers, &channelContainers)

	if err != nil {
		response.Reply("Fail to list containers.")
		log.Error(err)
		return
	}

	if len(containers) == 0 {
		response.Reply("No containers to list.")
		return
	}

	text := "Here is a list of all containers\n---\n"
	for _, container := range containers {
		line := fmt.Sprintf("`%s` \t", container.Name)
		if container.InUseBy == "free" {
			line += "_available_"
		} else {
			line += "in use"
		}
		text += fmt.Sprintf("%s\n", line)
	}
	response.Reply(text)
}
