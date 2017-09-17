package cmd

import (
	"fmt"
	"time"

	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
	Register("remove <container-name>", "Removes a container from your channel.", remove)
}

func remove(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	containerName := request.Param("container-name")

	if containerName == "" {
		response.Reply("Give me a container name to remove. 🙄")
		return
	}

	container, err := model.GetContainer(request.Event.Team, request.Event.Channel, containerName)

	if err != nil {
		log.Errorf("REMOVE. %s", err)
		response.Reply(err.Error())
		return
	}

	if container == (model.Container{}) {
		response.Reply(fmt.Sprintf("I couldn't find container `%s` on <#%s>.", containerName, request.Event.Channel))
	} else {
		if container.CreatedByUser != request.Event.User {
			response.Reply(fmt.Sprintf("Only who created can remove a container. Please check with <@%s>.", container.CreatedByUser))
		} else {
			if container.InUseBy != "" {
				response.Reply(fmt.Sprintf("Can't remove. Container `%s` is in used by <@%s> since _%s_.", containerName, container.InUseBy, container.UpdatedAt.Format(time.RFC1123)))
			} else {

				err = container.Delete()
				if err != nil {
					log.Errorf("Fail to remove the container. %s", err)
					response.Reply(err.Error())
					return
				}

				response.Reply(fmt.Sprintf("Container `%s` removed.", containerName))

			}
		}
	}
}
