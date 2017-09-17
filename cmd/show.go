package cmd

import (
	"fmt"
	"time"

	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
	Register("show <container-name>", "Shows a container details.", show)
}

func show(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	containerName := request.Param("container-name")

	container, err := model.GetContainer(request.Event.Team, request.Event.Channel, containerName)

	if err != nil {
		log.Errorf("SHOW. [%s, %s, %s] %s", request.Event.Team, request.Event.Channel, containerName, err)
		response.Reply(err.Error())
		return
	}

	if container == (model.Container{}) {
		response.Reply(fmt.Sprintf("I couldn't find the container `%s` on <#%s>.", containerName, request.Event.Channel))
		return
	}

	text := fmt.Sprintf("Container `%s`.\nCreated by <@%s>.\n", containerName, container.CreatedByUser)

	if container.InUseBy == "" {
		text += "_Available_"
	} else {
		text += fmt.Sprintf("In use by <@%s>", container.InUseBy)
	}
	text += fmt.Sprintf(" since _%s_.", container.UpdatedAt.Format(time.RFC1123))
	response.Reply(text)
}
