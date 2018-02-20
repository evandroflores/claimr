package cmd

import (
	"fmt"

	"strings"

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
		inUseMessageKey := utils.IfThenElse(container.InUseBy == event.User, "container-in-use-by-you", "container-in-use")
		response.Reply(fmt.Sprintf(Messages[inUseMessageKey.(string)], containerName))
		return
	}

	if container.SetInUse(event.User, getReason(request)) != nil {
		response.Reply(Messages["fail-to-update"])
		return
	}

	response.Reply(fmt.Sprintf(Messages["container-claimed"], containerName, event.User))
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
