package cmd

import (
	"fmt"

	"strings"

	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
	Register("claim <container-name> <reason>", "Claims a container for your use.", claim)
}

func claim(request *slacker.Request, response slacker.ResponseWriter) {
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
		response.Reply(fmt.Sprintf("Container `%s` is already in use, try another one.", containerName))
		return
	}

	container.InUseBy = event.User
	container.InUseForReason = getReason(request)

	err = container.Update()
	if err != nil {
		log.Errorf("Fail to update the container. %s", err)
		response.Reply("Fail to update the container.")
		return
	}

	response.Reply(fmt.Sprintf("Got it. Container `%s` is all yours <@%s>.", containerName, event.User))
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
