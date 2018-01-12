package admin

import (
	"fmt"

	"github.com/evandroflores/claimr/database"
	log "github.com/sirupsen/logrus"

	"github.com/evandroflores/slacker"
)

func init() {
	Register("log-level `level`", "Change the current log level", changeLogLevel)
}

func changeLogLevel(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	currentLogLevel := log.GetLevel()
	newLogLevel := request.Param("sub-command-parameter")

	if newLogLevel == "" {
		response.Reply(fmt.Sprintf("No log level provided, keeping in `%s`", currentLogLevel))
		return
	}

	logrusLevel, err := log.ParseLevel(newLogLevel)

	if err != nil {
		response.Reply(err.Error())
		return
	}

	database.DB.LogMode(logrusLevel == log.DebugLevel)

	log.SetLevel(logrusLevel)
	log.Printf("Log Test [%s -> %s]", currentLogLevel, newLogLevel)
	response.Reply(fmt.Sprintf("Log level changed from `%s` to `%s`", currentLogLevel, newLogLevel))
}
