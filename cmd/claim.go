package cmd

import (
	"fmt"

	"strings"

	"github.com/evandroflores/claimr/messages"
	"github.com/evandroflores/claimr/model"
	"github.com/evandroflores/claimr/utils"
	"github.com/shomali11/slacker"
)

func init() {
	Register("claim <container-name> <reason>", "Claims a container for your use.", claim)
}

func claim(request *slacker.Request, response slacker.ResponseWriter) {
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

	if container == (model.Container{}) {
		response.Reply(fmt.Sprintf(messages.Get("container-not-found-on-channel"), containerName, event.Channel))
		return
	}

	if container.InUseBy != "" {
		inUseMessageKey := utils.IfThenElse(container.InUseBy == event.User, "container-in-use-by-you", "container-in-use")
		response.Reply(fmt.Sprintf(messages.Get(inUseMessageKey.(string)), containerName))
		return
	}

	if container.SetInUse(event.User, getReason(request)) != nil {
		response.Reply(messages.Get("fail-to-update"))
		return
	}

	response.Reply(fmt.Sprintf(messages.Get("container-claimed"), containerName, event.User))
}

func getReason(request *slacker.Request) string {
	if request.Param("reason") == "" {
		return ""
	}
	allText := GetEventText(request)
	reasonToClaim := request.Param("reason")
	idx := strings.Index(allText, reasonToClaim)
	return allText[idx:]
}
