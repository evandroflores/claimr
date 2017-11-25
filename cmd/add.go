package cmd

import (
	"fmt"

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
	isDirect, msg := checkDirect(event.Channel)
	if isDirect {
		response.Reply(msg.Error())
		return
	}
	containerName := request.Param("container-name")

	container, err := model.GetContainer(event.Team, event.Channel, containerName)

	if err != nil {
		log.Errorf("ADD. [%s, %s, %s] %s", event.Team, event.Channel, containerName, err)
		response.Reply(err.Error())
		return
	}

	if container != (model.Container{}) {
		response.Reply("There is a container with the same name on this channel. Try a different one.")
		return
	}

	err = model.Container{
		TeamID:        event.Team,
		ChannelID:     event.Channel,
		Name:          containerName,
		InUseBy:       "",
		InUseByReason: "",
		CreatedByUser: event.User,
	}.Add()

	if err != nil {
		log.Errorf("ADD. [%s, %s, %s] %s", event.Team, event.Channel, containerName, err)
		response.Reply(err.Error())
		return
	}

	response.Reply(fmt.Sprintf("Container `%s` added to channel <#%s>.", containerName, event.Channel))
}
