package cmd

import (
	"fmt"

	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
	Register("free <container-name>", "Makes a container available for use.", free)
}

func free(request *slacker.Request, response slacker.ResponseWriter) {
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

	if container.InUseBy != event.User {
		response.Reply(fmt.Sprintf(Messages["container-in-use-by-other"], containerName))
		return
	}

	container.InUseBy = ""
	container.InUseForReason = ""

	err = container.Update()
	if err != nil {
		log.Errorf(Messages["fail-to-update"]+"%s", err)
		response.Reply(Messages["fail-to-update"])
		return
	}

	response.Reply(fmt.Sprintf(Messages["container-free"], containerName))
}
