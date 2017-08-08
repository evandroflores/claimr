package cmd

import (
	"fmt"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
	"strings"
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
		response.Reply(fmt.Sprintf("Try a name smaller than %d", maxNameSize))
		return
	}
	if len(containerName) == 0 {
		response.Reply("Please, give a name for your container. 🤦")
		return
	}

	found, _ := database.DB.Get(&model.Container{TeamID: request.Event.Team, Name: containerName})

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
		log.Error(err)
	}
	if affected == 1 {
		response.Reply(fmt.Sprintf("Container `%s` added to channel <#%s>.", container.Name, container.ChannelID))
		log.Debugf("Container %s added to channel %s.", container.Name, container.ChannelID)
	} else {
		log.Errorf("Tried to add container `%s` but failed.", container.Name)
		response.Reply("This doesn't smells good.")
	}
}
