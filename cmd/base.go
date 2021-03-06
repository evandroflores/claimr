package cmd

import (
	"fmt"
	"strings"

	"github.com/evandroflores/claimr/messages"
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

func validateInput(channelID string, message string) error {
	if direct, err := isDirect(channelID); direct {
		return err
	}

	if hasUser, err := hasUserOnText(message); hasUser {
		return err
	}

	if hasChannel, err := hasChannelOnText(message); hasChannel {
		return err
	}

	return nil
}

func isDirect(channelID string) (bool, error) {
	if strings.HasPrefix(strings.ToUpper(channelID), directMessagePrefix) {
		return true, fmt.Errorf(messages.Get("direct-not-allowed"))
	}
	return false, nil
}

func hasUserOnText(message string) (bool, error) {
	if strings.Contains(strings.ToUpper(message), userPrefix) {
		return true, fmt.Errorf(messages.Get("shouldnt-mention-user"))
	}
	return false, nil
}

func hasChannelOnText(message string) (bool, error) {
	if strings.Contains(strings.ToUpper(message), channelPrefix) {
		return true, fmt.Errorf(messages.Get("shouldnt-mention-channel"))
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
