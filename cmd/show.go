package cmd

import (
	"fmt"
	"time"

	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
	Register("show <container-name>", "Show who is using the container.", show)
}

func show(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	containerName := request.Param("container-name")

	if containerName == "" {
		response.Reply("Give me a container name to find. 🙄")
		return
	}

	container := model.Container{TeamID: request.Event.Team, ChannelID: request.Event.Channel, Name: containerName}

	found, err := database.DB.Get(&container)
	if err != nil {
		response.Reply("Fail to get container to show.")
		log.Errorf("Fail to get container to show. %s", err)
		return
	}

	if !found {
		response.Reply(fmt.Sprintf("I couldn't find the container `%s` on <#%s>.", containerName, request.Event.Channel))
	} else {
		text := fmt.Sprintf("Container `%s`.\nCreated by <@%s>.\n", containerName, container.CreatedByUser)

		if container.InUseBy == "free" {
			text += "_Available_"
		} else {
			text += fmt.Sprintf("In use by <@%s>", container.InUseBy)
		}
		text += fmt.Sprintf(" since _%s_.", container.UpdatedAt.Format(time.RFC1123))
		response.Reply(text)
	}
}
