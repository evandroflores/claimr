package cmd

import (
	"fmt"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
)

func init() {
	Register("free <container-name>", "Free a container from use", free)
}

func free(request *slacker.Request, response slacker.ResponseWriter) {
	containerName := request.Param("container-name")
	channel := request.Event.Channel
	user := request.Event.User

	response.Typing()

	container := model.Container{TeamID: request.Event.Team, Name: containerName}

	found, _ := database.DB.Get(&container)
	if !found {
		response.Reply(fmt.Sprintf("I couldn't find container `%s` on <#%s>.", containerName, channel))
	} else {
		if container.InUseBy != user {
			response.Reply(fmt.Sprintf("Humm Container `%s` is not being used by you.", containerName))
		} else {
			container.InUseBy = "free"
			database.DB.Id(container.ID).Update(&container)
			response.Reply(fmt.Sprintf("Got it. Container `%s` is now available", containerName))
		}
	}
}
