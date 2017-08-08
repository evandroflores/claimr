package cmd

import (
	"fmt"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	"time"
)

func init() {
	Register("show <container-name>", "Show who is using the container.", show)
}

func show(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	containerName := request.Param("container-name")
	channel := request.Event.Channel
	container := model.Container{TeamID: request.Event.Team, Name: containerName}

	found, _ := database.DB.Get(&container)
	if !found {
		response.Reply(fmt.Sprintf("I couldn't find container `%s` on <#%s>.", containerName, channel))
	} else {
		text := fmt.Sprintf("Container `%s`.\nCreated by <@%s>.\n", containerName, container.CreatedByUser)

		if container.InUseBy == "free" {
			text += "_Available_"
		} else {
			text += fmt.Sprintf("Being used by <@%s>", container.InUseBy)
		}
		text += fmt.Sprintf(" since %s.", container.UpdatedAt.Format(time.RFC1123))
		response.Reply(text)
	}
}
