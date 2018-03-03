package cmd

import (
	"fmt"

	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/messages"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

func init() {
	Register("log-level <level>", "Change the current log level. admin-only", changeLogLevel)
}

func changeLogLevel(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	event := getEvent(request)
	if !isAdmin(event.User) {
		response.Reply(messages.Get("admin-only"))
		return
	}

	currentLogLevel := log.GetLevel()
	newLogLevel := request.Param("level")

	if newLogLevel == "" {
		response.Reply(fmt.Sprintf(messages.Get("no-level-provided"), currentLogLevel))
		return
	}

	logrusLevel, err := log.ParseLevel(newLogLevel)

	if err != nil {
		response.Reply(err.Error())
		return
	}

	if currentLogLevel == logrusLevel {
		response.Reply(messages.Get("same-log-level"))
		return
	}

	database.DB.LogMode(logrusLevel == log.DebugLevel)

	log.SetLevel(logrusLevel)
	log.Printf("Log Test [%s -> %s]", currentLogLevel, newLogLevel)
	response.Reply(fmt.Sprintf(messages.Get("level-log-changed"), currentLogLevel, newLogLevel))
}
