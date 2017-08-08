package cmd

import (
	"fmt"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
)

func init() {
	Register("claim <container-name>", "Claim a container for your use", claim)
}

func claim(request *slacker.Request, response slacker.ResponseWriter) {
	containerName := request.Param("container-name")
	channel := request.Event.Channel
	user := request.Event.User

	response.Typing()

	container := model.Container{TeamID: request.Event.Team, Name: containerName}

	found, _ := database.DB.Get(&container)
	if !found {
		response.Reply(fmt.Sprintf("I couldn't find container `%s` on <#%s>.", containerName, channel))
	} else {
		if container.InUseBy != "free" {
			response.Reply(fmt.Sprintf("Container `%s` is already in use, try another one.", containerName))
		} else {
			container.InUseBy = user
			database.DB.Id(container.ID).Update(&container)
			response.Reply(fmt.Sprintf("Got it. Container `%s` is all yours <@%s>.", containerName, user))
		}
	}
}
