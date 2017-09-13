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
	Register("list", "List all containers.", list)
}

func list(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	containers := make([]model.Container, 0)
	channelContainers := model.Container{TeamID: request.Event.Team, ChannelID: request.Event.Channel}

	err := database.DB.Find(&containers, &channelContainers)

	if err != nil {
		response.Reply("Fail to list containers.")
		log.Error(err)
		return
	}

	if len(containers) == 0 {
		response.Reply("No containers to list.")
		return
	}

	containerList := []string {"Here is a list of containers for this channel:"}
	for _, container := range containers {
		line := fmt.Sprintf("`%s`\t%s %s", container.Name,
			IfThenElse(container.InUseBy == "free", "_available_", "in use"),
			IfThenElse(container.InUseByReason != "free", fmt.Sprintf("- %s", container.InUseByReason), ""),
		)
		containerList = append(containerList, line)
	}
	response.Reply(strings.Join(containerList,"\n"))
}

// IfThenElse as Golang does not have ternary ifelse
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
