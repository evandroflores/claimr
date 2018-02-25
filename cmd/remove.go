package cmd

import (
	"fmt"
	"time"

	"github.com/evandroflores/claimr/messages"
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

	checks := []Check{
		{container == (model.Container{}), fmt.Sprintf(messages.Messages["container-not-found-on-channel"], containerName, event.Channel)},
		{container.InUseBy != "", fmt.Sprintf(messages.Messages["container-in-use-by-this"], containerName, container.InUseBy, container.UpdatedAt.Format(time.RFC1123))},
		{container.CreatedByUser != event.User, fmt.Sprintf(messages.Messages["only-owner-can-remove"], containerName, container.CreatedByUser)},
	}

	err = RunChecks(checks)
	if err != nil {
		response.Reply(err.Error())
		return
	}

	err = container.Delete()
	if err != nil {
		response.Reply(err.Error())
		return
	}

	response.Reply(fmt.Sprintf(messages.Messages["container-removed"], containerName))
}
