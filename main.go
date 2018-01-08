package main

import (
	"os"

	"github.com/evandroflores/claimr/cmd"
	"github.com/evandroflores/claimr/database"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)
	token := os.Getenv("CLAIMR_TOKEN")
	if token == "" {
		log.Fatal("Claimr slack bot token unset. Set CLAIMR_TOKEN to continue.")
	}
	bot := slacker.NewClient(token)

	log.Debug("Loading commands...")
	for _, command := range cmd.CommandList() {
		log.Infof("%s - %s", command.Usage, command.Description)
		bot.Command(command.Usage, command.Description, command.Handler)
	}
	bot.Default(func(request *slacker.Request, response slacker.ResponseWriter) {
		response.Reply("Not sure what you are asking for. Type `@claimr help` for valid commands.")
	})

	log.Infof("Commands loaded. (%d)", len(cmd.CommandList()))

	err := bot.Listen()
	if err != nil {
		log.Fatal(err)
	}
	defer database.DB.Close()
}
