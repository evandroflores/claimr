package cmd

import (
	"fmt"
	"time"

	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
)

func init() {
	Register("remove <container-name>", "Removes a container from your channel.", remove)
}

func remove(request *slacker.Request, response slacker.ResponseWriter) {
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
		response.Reply(fmt.Sprintf(Messages["container-not-found-on-channel"], containerName, event.Channel))
		return
	}

	if container.InUseBy != "" {
		response.Reply(fmt.Sprintf(Messages["container-in-use-by-this"], containerName, container.InUseBy, container.UpdatedAt.Format(time.RFC1123)))
		return
	}

	if container.CreatedByUser != event.User {
		response.Reply(fmt.Sprintf(Messages["only-owner-can-remove"], containerName, container.CreatedByUser))
		return
	}

	err = container.Delete()
	if err != nil {
		response.Reply(err.Error())
		return
	}

	response.Reply(fmt.Sprintf(Messages["container-removed"], containerName))
}
