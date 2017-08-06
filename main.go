package main

import (
	"github.com/evandroflores/claimr/cmd"
	"github.com/evandroflores/claimr/database"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
    database.InitDB()
    log.SetLevel(log.DebugLevel)
}

func main() {
	bot := slacker.NewClient("xoxb-221107798822-MhoNS4UseJkvo5azVFKRjpud")

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
