package cmd

import (
	"fmt"
	"strings"

	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

var (
	maxNameSize         = 22
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
	if len(containerName) > maxNameSize {
		response.Reply(fmt.Sprintf("Try a name up to %d characters.", maxNameSize))
		return
	}

	if containerName == "" {
		response.Reply("Please, give a name for your container. 🤦")
		return
	}

	found, err := database.DB.Get(&model.Container{TeamID: request.Event.Team, Name: containerName})

	if err != nil {
		log.Errorf("Fail to get check if the container exists. %s", err)
		response.Reply("Fail to get check if the container exists.")
		return
	}

	if found {
		response.Reply("There is a container with the same name on this channel. Try a different one.")
		return
	}

	container := new(model.Container)
	container.Name = containerName
	container.TeamID = request.Event.Team
	container.ChannelID = request.Event.Channel
	container.InUseBy = "free"
	container.CreatedByUser = request.Event.User

	affected, err := database.DB.Insert(container)

	if err != nil {
		log.Errorf("Tried to add container `%s` but failed. %s", container.Name, err)
		response.Reply("Fail to add the container.")
		return
	}

	if affected == 1 {
		log.Debugf("Container %s added to channel %s.", container.Name, container.ChannelID)
		response.Reply(fmt.Sprintf("Container `%s` added to channel <#%s>.", container.Name, container.ChannelID))
	} else {
		log.Errorf("`%d` containers were removed when trying to remove container named `%s` on channel `%s` for team `%s`", affected, containerName, request.Event.Channel, request.Event.Team)
		response.Reply("Humm, this looks wrong. Nothing was added. 🤔")
	}
}
