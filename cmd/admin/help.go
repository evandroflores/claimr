package admin

import (
	"fmt"
	"strings"

	"github.com/evandroflores/slacker"
)

func init() {
	Register("command-list", "List admin sub commands list (help is reserved for the main commands)", subCommandHelp)
}

func subCommandHelp(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	help := GenerateCommandHelp()
	response.Reply(help)
}

//GenerateCommandHelp is basically what says it is.
func GenerateCommandHelp() string {
	help := ""
	for _, subcommand := range CommandList() {
		splitted := strings.Split(subcommand.Usage, " ")
		help += fmt.Sprintf("*%s* %s - _%s_\n", splitted[0], strings.Join(splitted[1:], " "), subcommand.Description)
	}
	return help
}
