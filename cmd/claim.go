package cmd

import (
	"fmt"

	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
	"strings"
)

func init() {
	Register("claim <container-name> <reason>", "Claim a container for your use.", claim)
}

func claim(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	containerName := request.Param("container-name")

	if containerName == "" {
		response.Reply("Please, give a container name for search. 🤦")
		return
	}

	container := model.Container{TeamID: request.Event.Team, Name: containerName}

	found, err := database.DB.Get(&container)

	if err != nil {
		log.Errorf("Fail to get the container. %s", err)
		response.Reply("Fail to get the container.")
		return
	}

	if !found {
		response.Reply(fmt.Sprintf("I couldn't find container `%s` on <#%s>.", containerName, request.Event.Channel))
	} else {
		if container.InUseBy != "free" {
			response.Reply(fmt.Sprintf("Container `%s` is already in use, try another one.", containerName))
		} else {
			container.InUseBy = request.Event.User
			container.InUseByReason = getReason(*request)

			affected := int64(0)
			affected, err = database.DB.ID(container.ID).Update(&container)
			if err != nil {
				log.Errorf("Fail to update the container. %s", err)
				response.Reply("Fail to update the container.")
				return
			}

			if affected == 1 {
				response.Reply(fmt.Sprintf("Got it. Container `%s` is all yours <@%s>.", containerName, request.Event.User))
			} else {
				log.Errorf("`%d` containers were update when trying to update container named `%s` on channel `%s` for team `%s`", affected, containerName, request.Event.Channel, request.Event.Team)
				response.Reply("Humm, this looks wrong. 🤔")
			}

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