package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
)

var commands []model.Command

const (
	directMessagePrefix = "D"
	channelPrefix       = "<#C"
	userPrefix          = "<@U"
)

// Register add a command to commands list an prepare to register to slacker
func Register(usage string, description string, handler func(request *slacker.Request, response slacker.ResponseWriter)) {
	commands = append(commands, model.Command{Usage: usage, Description: description, Handler: handler})
}

// CommandList returns the list of registered commands
func CommandList() []model.Command {
	return commands
}

func notImplemented(request *slacker.Request, response slacker.ResponseWriter) {
	response.Reply(Messages["not-implemented"])
}

func isDirect(channelID string) (bool, error) {
	if strings.HasPrefix(channelID, directMessagePrefix) {
		return true, fmt.Errorf(Messages["direct-not-allowed"])
	}
	return false, nil
}

func hasUserOnText(message string) (bool, error) {
	if strings.HasPrefix(message, userPrefix) {
		return true, fmt.Errorf(Messages["shouldnt-mention-user"])
	}
	return false, nil
}

func hasChannelOnText(message string) (bool, error) {
	if strings.HasPrefix(message, channelPrefix) {
		return true, fmt.Errorf(Messages["shouldnt-mention-channel"])
	}
	return false, nil
}

func getEvent(request *slacker.Request) ClaimrEvent {
	if request == nil {
		return ClaimrEvent{}
	}
	return ClaimrEvent{
		Team:    request.Event.Team,
		Channel: request.Event.Channel,
		User:    request.Event.User,
	}
}

// ClaimrEvent is a struct to simplify the usage of request.Event (and help testing)
type ClaimrEvent struct {
	Team    string
	Channel string
	User    string
}

// GetEventText exists to help testing event message
func GetEventText(request *slacker.Request) string {
	return request.Event.Msg.Text
}

func isAdmin(userName string) bool {
	return userName == os.Getenv("CLAIMR_SUPERUSER")
}
