package cmd

import (
	"fmt"

	"strings"

	"github.com/evandroflores/claimr/messages"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
)

func init() {
	Register("list", "List all containers.", list)
}

func list(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	event := getEvent(request)
	if direct, err := isDirect(event.Channel); direct {
		response.Reply(err.Error())
		return
	}

	containers, err := model.GetContainers(event.Team, event.Channel)

	if err != nil {
		response.Reply(messages.Get("fail-getting-containers"))
		return
	}

	if len(containers) == 0 {
		response.Reply(messages.Get("empty-containers-list"))
		return
	}

	containerList := []string{messages.Get("containers-list")}
	for _, container := range containers {
		line := fmt.Sprintf("`%s`\t%s", container.Name, container.InUseText("simple"))
		containerList = append(containerList, line)
	}
	response.Reply(strings.Join(containerList, "\n"))
}
