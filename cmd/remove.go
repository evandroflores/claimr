package cmd

import (
	"fmt"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	"time"
)

func init() {
	Register("rm <container-name>", "Remove a container from your channel.", remove)
}

func remove(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	containerName := request.Param("container-name")

	if len(containerName) == 0 {
		response.Reply("Give me a container name to remove. 🙄")
		return
	}
	channel := request.Event.Channel
	user := request.Event.User

	container := model.Container{TeamID: request.Event.Team, Name: containerName}

	found, _ := database.DB.Get(&container)
	if !found {
		response.Reply(fmt.Sprintf("I couldn't find container `%s` on <#%s>.", containerName, channel))
	} else {
		if container.CreatedByUser != user {
			response.Reply(fmt.Sprintf("Only who created can remove a container. Please check with <@%s>.", container.CreatedByUser))
		} else {
			if container.InUseBy != "free" {
				response.Reply(fmt.Sprintf("Can't remove. Container `%s` is being used by <@%s> since %s.", containerName, container.InUseBy, container.UpdatedAt.Format(time.RFC1123)))
			} else {
				affected, _ := database.DB.ID(container.ID).Delete(&container)
				if affected == 1 {
					response.Reply(fmt.Sprintf("Container `%s` removed.", containerName))
				} else {
					response.Reply("This doesn't smells good.")
				}

			}
		}
	}
}
