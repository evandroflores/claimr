package cmd

import (
	"fmt"

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

	isDirect, msg := checkDirect(request.Event.Channel)
	if isDirect {
		response.Reply(msg.Error())
		return
	}

	containers, err := model.GetContainers(request.Event.Team, request.Event.Channel)

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
			IfThenElse(container.InUseBy != "", "in use", "_available_"),
			IfThenElse(container.InUseByReason != "", fmt.Sprintf("- %s", container.InUseByReason), ""),
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
