package cmd

import (
	"fmt"

	"github.com/evandroflores/claimr/messages"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
)

func init() {
	Register("free <container-name>", "Makes a container available for use.", free)
}

func free(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	event := getEvent(request)

	containerName := request.Param("container-name")

	err := validateInput(event.Channel, containerName)
	if err != nil {
		response.Reply(err.Error())
		return
	}

	container, err := model.GetContainer(event.Team, event.Channel, containerName)

	if err != nil {
		response.Reply(err.Error())
		return
	}

	checks := []Check{
		{container == (model.Container{}), fmt.Sprintf(messages.Get("container-not-found-on-channel"), containerName, event.Channel)},
		{container.InUseBy != event.User, fmt.Sprintf(messages.Get("container-in-use-by-other"), containerName)},
	}

	err = RunChecks(checks)
	if err != nil {
		response.Reply(err.Error())
		return
	}

	err = container.ClearInUse()
	if err != nil {
		response.Reply(messages.Get("fail-to-update"))
		return
	}

	response.Reply(fmt.Sprintf(messages.Get("container-free"), containerName))
}
