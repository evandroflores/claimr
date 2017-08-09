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
	Register("rm <container-name>", "Remove a container from your channel.", remove)
}

func remove(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	containerName := request.Param("container-name")

	if containerName == "" {
		response.Reply("Give me a container name to remove. 🙄")
		return
	}

	container := model.Container{TeamID: request.Event.Team, Name: containerName}

	found, err := database.DB.Get(&container)

	if err != nil {
		log.Errorf("Fail to get container to remove. %s", err)
		response.Reply("Fail to get container to remove.")
		return
	}

	if !found {
		response.Reply(fmt.Sprintf("I couldn't find container `%s` on <#%s>.", containerName, request.Event.Channel))
	} else {
		if container.CreatedByUser != request.Event.User {
			response.Reply(fmt.Sprintf("Only who created can remove a container. Please check with <@%s>.", container.CreatedByUser))
		} else {
			if container.InUseBy != "free" {
				response.Reply(fmt.Sprintf("Can't remove. Container `%s` is in used by <@%s> since _%s_.", containerName, container.InUseBy, container.UpdatedAt.Format(time.RFC1123)))
			} else {
				affected := int64(0)
				affected, err = database.DB.ID(container.ID).Delete(&container)
				if err != nil {
					log.Errorf("Fail to remove the container. %s", err)
					response.Reply("Fail to remove the container.")
				}
				if affected == 1 {
					response.Reply(fmt.Sprintf("Container `%s` removed.", containerName))
				} else {
					log.Errorf("%d containers were removed when trying to remove container named `%s` on channel `%s` for team `%s`", affected, containerName, request.Event.Channel, request.Event.Team)
					response.Reply("Humm, this looks wrong. 🤔")
				}
			}
		}
	}
}
