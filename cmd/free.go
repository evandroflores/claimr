package cmd

import (
	"fmt"

	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
	Register("free <container-name>", "Makes a container available for use.", free)
}

func free(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	isDirect, msg := checkDirect(request.Event.Channel)
	if isDirect {
		response.Reply(msg.Error())
		return
	}

	containerName := request.Param("container-name")

	container, err := model.GetContainer(request.Event.Team, request.Event.Channel, containerName)

	if err != nil {
		log.Errorf("FREE. %s", err)
		response.Reply(err.Error())
		return
	}

	if container == (model.Container{}) {
		response.Reply(fmt.Sprintf("I couldn't find container `%s` on <#%s>.", containerName, request.Event.Channel))
	} else {
		if container.InUseBy != request.Event.User {
			response.Reply(fmt.Sprintf("Humm Container `%s` is not being used by you.", containerName))
		} else {
			container.InUseBy = ""
			container.InUseByReason = ""

			err = container.Update()
			if err != nil {
				log.Errorf("Fail to update the container. %s", err)
				response.Reply("Fail to update the container.")
				return
			}

			response.Reply(fmt.Sprintf("Got it. Container `%s` is now available", containerName))
		}
	}
}
