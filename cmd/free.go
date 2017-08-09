package cmd

import (
	"fmt"

	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
	Register("free <container-name>", "Free a container from use.", free)
}

func free(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	containerName := request.Param("container-name")

	if containerName == "" {
		response.Reply("Please, give a container name for search. 🤦")
		return
	}

	container := model.Container{TeamID: request.Event.Team, Name: containerName}

	found, err := database.DB.Get(&container)

	if err != nil {
		log.Errorf("Fail to get the container. %s", err)
		response.Reply("Fail to get the container.")
		return
	}

	if !found {
		response.Reply(fmt.Sprintf("I couldn't find container `%s` on <#%s>.", containerName, request.Event.Channel))
	} else {
		if container.InUseBy != request.Event.User {
			response.Reply(fmt.Sprintf("Humm Container `%s` is not being used by you.", containerName))
		} else {
			container.InUseBy = "free"

			affected := int64(0)
			affected, err = database.DB.ID(container.ID).Update(&container)
			if err != nil {
				log.Errorf("Fail to update the container. %s", err)
				response.Reply("Fail to update the container.")
				return
			}

			if affected == 1 {
				response.Reply(fmt.Sprintf("Got it. Container `%s` is now available", containerName))
			} else {
				log.Errorf("`%d` containers were update when trying to update container named `%s` on channel `%s` for team `%s`", affected, containerName, request.Event.Channel, request.Event.Team)
				response.Reply("Humm, this looks wrong. 🤔")
			}
		}
	}
}
