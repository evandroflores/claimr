package cmd

import (
	"fmt"
	"time"

	"github.com/evandroflores/claimr/model"
	"github.com/evandroflores/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
	Register("remove <container-name>", "Removes a container from your channel.", remove)
}

func remove(request *slacker.Request, response slacker.ResponseWriter) {
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

	if container.InUseBy != "" {
		response.Reply(fmt.Sprintf("Can't remove. Container `%s` is in used by <@%s> since _%s_.", containerName, container.InUseBy, container.UpdatedAt.Format(time.RFC1123)))
		return
	}

	if container.CreatedByUser != event.User {
		response.Reply(fmt.Sprintf("Only who created the container `%s` can remove it. Please check with <@%s>.", containerName, container.CreatedByUser))
		return
	}

	err = container.Delete()
	if err != nil {
		log.Errorf("Fail to remove the container. %s", err)
		response.Reply(err.Error())
		return
	}

	response.Reply(fmt.Sprintf("Container `%s` removed.", containerName))
}
