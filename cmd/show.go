package cmd

import (
	"fmt"
	"time"

	"github.com/evandroflores/claimr/messages"
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
	if direct, err := isDirect(event.Channel); direct {
		response.Reply(err.Error())
		return
	}

	containerName := request.Param("container-name")

	container, err := model.GetContainer(event.Team, event.Channel, containerName)

	if err != nil {
		response.Reply(err.Error())
		return
	}

	if container == (model.Container{}) {
		response.Reply(fmt.Sprintf(messages.Get("container-not-found-on-channel"), containerName, event.Channel))
		return
	}

	text := fmt.Sprintf(messages.Get("container-created-by"), containerName, container.CreatedByUser)

	if container.InUseBy == "" {
		text += "_Available_"
	} else {
		text += fmt.Sprintf(
			messages.Get("container-in-use-by-w-reason"),
			container.InUseBy,
			utils.IfThenElse(container.InUseForReason != "", fmt.Sprintf(" for %s", container.InUseForReason), ""),
			container.UpdatedAt.Format(time.RFC1123))
	}
	response.Reply(text)
}
