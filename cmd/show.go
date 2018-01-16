package cmd

import (
	"fmt"
	"time"

	"github.com/evandroflores/claimr/model"
	"github.com/evandroflores/claimr/utils"
	"github.com/shomali11/slacker"
)

func init() {
	Register("show <container-name>", "Shows a container details.", show)
}

func show(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	event := getEvent(request)
	isDirect, msg := checkDirect(event.Channel)
	if isDirect {
		response.Reply(msg.Error())
		return
	}

	containerName := request.Param("container-name")

	container, err := model.GetContainer(event.Team, event.Channel, containerName)

	if err != nil {
		response.Reply(err.Error())
		return
	}

	if container == (model.Container{}) {
		response.Reply(fmt.Sprintf("I couldn't find the container `%s` on <#%s>.", containerName, event.Channel))
		return
	}

	text := fmt.Sprintf("Container `%s`.\nCreated by <@%s>.\n", containerName, container.CreatedByUser)

	if container.InUseBy == "" {
		text += "_Available_"
	} else {
		text += fmt.Sprintf("In use by <@%s>%s", container.InUseBy,
			utils.IfThenElse(container.InUseForReason != "", fmt.Sprintf(" for %s", container.InUseForReason), ""))
	}
	text += fmt.Sprintf(" since _%s_.", container.UpdatedAt.Format(time.RFC1123))
	response.Reply(text)
}
