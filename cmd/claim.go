package cmd

import (
	"fmt"

	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
	"strings"
)

func init() {
	Register("claim <container-name> <reason>", "Claims a container for your use.", claim)
}

func claim(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	isDirect, msg := checkDirect(request.Event.Channel)
	if isDirect {
		response.Reply(msg.Error())
		return
	}

	containerName := request.Param("container-name")

	container, err := model.GetContainer(request.Event.Team, request.Event.Channel, containerName)

	if err != nil {
		log.Errorf("CLAIM. %s", err)
		response.Reply(err.Error())
		return
	}

	if container == (model.Container{}){
		response.Reply(fmt.Sprintf("I couldn't find container `%s` on <#%s>.", containerName, request.Event.Channel))
	} else {
		if container.InUseBy != "" {
			response.Reply(fmt.Sprintf("Container `%s` is already in use, try another one.", containerName))
		} else {
			container.InUseBy = request.Event.User
			container.InUseByReason = getReason(*request)

			err = container.Update()
			if err != nil {
				log.Errorf("Fail to update the container. %s", err)
				response.Reply("Fail to update the container.")
				return
			}

			response.Reply(fmt.Sprintf("Got it. Container `%s` is all yours <@%s>.", containerName, request.Event.User))
		}
	}
}

func getReason(request slacker.Request) string {
	allText := request.Event.Msg.Text
	reasonToClaim := request.Param("reason")
	idx := strings.Index(allText, reasonToClaim)
	if idx == -1 || idx == 0{
		return reasonToClaim
	}

	return allText[idx:]
}