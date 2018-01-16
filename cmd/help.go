package cmd

import (
	"fmt"
	"strings"

	"github.com/shomali11/slacker"
)

// Help overides the default help function
func Help(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	help := GenerateCommandHelp()
	response.Reply(help)
}

// GenerateCommandHelp is basically what says it is.
func GenerateCommandHelp() string {
	helpPattern := "*%s* %s - _%s_\n"
	help := fmt.Sprintf(helpPattern, "help", "", "Shows this command list.")
	replacer := strings.NewReplacer("<", "`", ">", "`")
	for _, subcommand := range CommandList() {
		splitted := strings.Split(subcommand.Usage, " ")
		help += fmt.Sprintf(helpPattern, splitted[0], replacer.Replace(strings.Join(splitted[1:], " ")), subcommand.Description)
	}
	return help
}
