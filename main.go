package main

import (
	"github.com/evandroflores/claimr/cmd"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetLevel(log.DebugLevel)
	token := os.Getenv("CLAIMR_TOKEN")
	if token == "" {
		log.Fatal("Claimr slack bot token unset. Set CLAIMR_TOKEN to continue.")
	}
	bot := slacker.NewClient(token)

	log.Info("Loading commands...")
	for _, command := range cmd.CommandList() {
		log.Debugf("%s - %s", command.Usage, command.Description)
		bot.Command(command.Usage, command.Description, command.Handler)
	}
	log.Infof("Commands loaded. [%d]", len(cmd.CommandList()))

	err := bot.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
