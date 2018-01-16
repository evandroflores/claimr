package cmd

import (
	"fmt"

	"strings"

	"github.com/evandroflores/claimr/model"
	"github.com/evandroflores/claimr/utils"
	"github.com/shomali11/slacker"
)

func init() {
	Register("list", "List all containers.", list)
}

func list(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	event := getEvent(request)
	isDirect, msg := checkDirect(event.Channel)
	if isDirect {
		response.Reply(msg.Error())
		return
	}

	containers, err := model.GetContainers(event.Team, event.Channel)

	if err != nil {
		response.Reply("Fail to list containers.")
		return
	}

	if len(containers) == 0 {
		response.Reply("No containers to list.")
		return
	}

	containerList := []string{"Here is a list of containers for this channel:"}
	for _, container := range containers {
		line := fmt.Sprintf("`%s`\t%s %s", container.Name,
			utils.IfThenElse(container.InUseBy != "", "in use", "_available_"),
			utils.IfThenElse(container.InUseForReason != "", fmt.Sprintf("- %s", container.InUseForReason), ""),
		)
		containerList = append(containerList, line)
	}
	response.Reply(strings.Join(containerList, "\n"))
}
