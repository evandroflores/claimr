package cmd

import (
	"fmt"

	"github.com/evandroflores/claimr/messages"
	"github.com/evandroflores/claimr/model"
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
	text += container.InUseText("full")
	response.Reply(text)
}
