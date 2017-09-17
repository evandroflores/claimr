package cmd

import (
	"fmt"
	"strings"

	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

var (
	directChannelPrefix = "D"
)

func init() {
	Register("add <container-name>", "Add a container to your channel.", add)
}

func add(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	if strings.HasPrefix(request.Event.Channel, directChannelPrefix) {
		response.Reply("This look like a direct message. Containers are related to a channel.")
		return
	}

	containerName := request.Param("container-name")

	container, err := model.GetContainer(request.Event.Team, request.Event.Channel, containerName)

	if err != nil {
		log.Errorf("ADD. [%s, %s, %s] %s", request.Event.Team, request.Event.Channel, containerName, err)
		response.Reply(err.Error())
		return
	}

	if container != (model.Container{}) {
		response.Reply("There is a container with the same name on this channel. Try a different one.")
		return
	}

	err = model.Container{
			TeamID: request.Event.Team,
			ChannelID: request.Event.Channel,
			Name: containerName,
			InUseBy: "",
			InUseByReason: "",
			CreatedByUser: request.Event.User,
		}.Add()

	if err != nil {
		log.Errorf("ADD. [%s, %s, %s] %s", request.Event.Team, request.Event.Channel, containerName, err)
		response.Reply(err.Error())
		return
	}

	response.Reply(fmt.Sprintf("Container `%s` added to channel <#%s>.", containerName, request.Event.Channel))
}
