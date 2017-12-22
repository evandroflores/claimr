package cmd

import (
	"fmt"
	"strings"

	"github.com/shomali11/slacker"
)

var (
	commands            []Command
	directChannelPrefix = "D"
)

// Command defines a command to be register to slack
type Command struct {
	Usage       string
	Description string
	Handler     func(request *slacker.Request, response slacker.ResponseWriter)
}

// Register add a command to commands list an prepare to register to slacker
func Register(usage string, description string, handler func(request *slacker.Request, response slacker.ResponseWriter)) {
	commands = append(commands, Command{Usage: usage, Description: description, Handler: handler})
}

// CommandList returns the list of registered commands
func CommandList() []Command {
	return commands
}

func notImplemented(request *slacker.Request, response slacker.ResponseWriter) {
	response.Reply("No pancakes for you! 🥞")
}

func checkDirect(channelID string) (bool, error) {
	if strings.HasPrefix(channelID, directChannelPrefix) {
		return true, fmt.Errorf("this look like a direct message. Containers are related to a channels")
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
