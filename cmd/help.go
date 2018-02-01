package cmd

import (
	"fmt"
	"strings"

	"github.com/shomali11/slacker"
)

// Help overides the default help function
func Help(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	event := getEvent(request)
	help := GenerateCommandHelp(isAdmin(event.User))
	response.Reply(help)
}

// GenerateCommandHelp is basically what says it is.
func GenerateCommandHelp(showAdminCommands bool) string {
	helpPattern := "*%s* %s - _%s_\n"
	help := fmt.Sprintf(helpPattern, "help", "", "Shows this command list.")
	replacer := strings.NewReplacer("<", "`", ">", "`")
	for _, command := range CommandList() {
		if !showAdminCommands && strings.Contains(command.Description, "admin-only") {
			continue
		}
		splitted := strings.Split(command.Usage, " ")
		help += fmt.Sprintf(helpPattern, splitted[0], replacer.Replace(strings.Join(splitted[1:], " ")), command.Description)
	}
	return help
}
