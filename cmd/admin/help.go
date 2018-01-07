package admin

import (
	"fmt"
	"strings"

	"github.com/shomali11/slacker"
)

func init() {
	Register("command-list", "List admin sub commands list (help is reserved for the main commands)", AdminHelp)
}

//AdminHelp prints the help for Admin subcommands
func AdminHelp(request *slacker.Request, response slacker.ResponseWriter) {
	help := ""
	for _, subcommand := range CommandList() {
		splitted := strings.Split(subcommand.Usage, " ")
		help += fmt.Sprintf("*%s* %s - _%s_\n", splitted[0], strings.Join(splitted[1:], " "), subcommand.Description)
	}
	response.Reply(help)
}
