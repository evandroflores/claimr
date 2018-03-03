package cmd

import (
	"fmt"

	"github.com/evandroflores/claimr/messages"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
	Register("add <container-name>", "Adds a container to your channel.", add)
}

func add(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	event := getEvent(request)

	if direct, err := isDirect(event.Channel); direct {
		response.Reply(err.Error())
		return
	}

	containerName := request.Param("container-name")

	if hasUserOrChannel, err := hasUserOrChannelOnText(containerName); hasUserOrChannel {
		response.Reply(err.Error())
		return
	}

	container, err := model.GetContainer(event.Team, event.Channel, containerName)

	if err != nil {
		response.Reply(err.Error())
		return
	}

	if container != (model.Container{}) {
		response.Reply(messages.Get("same-name"))
		return
	}

	err = model.Container{
		TeamID: event.Team, ChannelID: event.Channel, Name: containerName, InUseBy: "", InUseForReason: "", CreatedByUser: event.User,
	}.Add()

	if err != nil {
		log.Errorf("ADD. [%s, %s, %s] %s", event.Team, event.Channel, containerName, err)
		response.Reply(err.Error())
		return
	}

	response.Reply(fmt.Sprintf(messages.Get("added-to-channel"), containerName, event.Channel))
}
