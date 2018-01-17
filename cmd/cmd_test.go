package cmd

import (
	"fmt"
	"strings"
	"testing"

	"reflect"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/model"
	"github.com/nlopes/slack"
	"github.com/shomali11/proper"
	"github.com/shomali11/slacker"
	"github.com/stretchr/testify/assert"
)

func createMockReply(t *testing.T, expectedMsg string) (*slacker.Response, *monkey.PatchGuard) {
	var mockResponse *slacker.Response

	patchReply := monkey.PatchInstanceMethod(reflect.TypeOf(mockResponse), "Reply",
		func(response *slacker.Response, msg string) {
			assert.Equal(t, expectedMsg, msg)
		})

	_ = monkey.PatchInstanceMethod(reflect.TypeOf(mockResponse), "Typing",
		func(response *slacker.Response) {})

	return mockResponse, patchReply
}

func createMockRequest(t *testing.T, params map[string]string) (*slacker.Request, *monkey.PatchGuard) {
	var mockRequest *slacker.Request

	patchParam := monkey.PatchInstanceMethod(reflect.TypeOf(mockRequest), "Param",
		func(r *slacker.Request, key string) string {
			if params == nil {
				return ""
			}
			return params[key]
		})
	return mockRequest, patchParam
}

func createMockEvent(t *testing.T, team string, channel string, user string) *monkey.PatchGuard {
	patchGetEvent := monkey.Patch(getEvent,
		func(request *slacker.Request) ClaimrEvent {
			return ClaimrEvent{team, channel, user}
		})
	return patchGetEvent
}

func TestCmdNotImplemented(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, Messages["not-implemented"])

	notImplemented(new(slacker.Request), mockResponse)

	patchReply.Unpatch()
}

func TestCmdCommandList(t *testing.T) {
	usageExpected := []string{
		"add <container-name>",
		"claim <container-name> <reason>",
		"free <container-name>",
		"list",
		"remove <container-name>",
		"show <container-name>",
		"log-level <level>",
		"purge",
	}
	commands := CommandList()

	assert.Len(t, commands, len(usageExpected))

	usageActual := []string{}

	for _, command := range commands {
		usageActual = append(usageActual, command.Usage)
	}

	assert.Subset(t, usageExpected, usageActual)
}

func TestCmdNotDirect(t *testing.T) {
	isDirect, err := checkDirect("CHANNEL")
	assert.False(t, isDirect)
	assert.NoError(t, err)
}

func TestCmdDirect(t *testing.T) {
	isDirect, err := checkDirect("DIRECT")
	assert.True(t, isDirect)
	assert.Error(t, err, Messages["direct-not-allowed"])
}

func TestAllCmdsCheckingDirect(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, Messages["direct-not-allowed"])
	patchGetEvent := createMockEvent(t, "team", "DIRECT", "user")

	for _, command := range commands {
		if !strings.Contains(command.Description, "admin-only") {
			command.Handler(nil, mockResponse)
		}
	}

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
}

func TestAllCmdsCheckingNoName(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(Messages["field-container-name"]))
	patchGetEvent := createMockEvent(t, "team", "channel", "user")
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": ""})

	for _, command := range CommandList() {
		if strings.Contains(command.Usage, "<container-name>") {
			command.Handler(mockRequest, mockResponse)
		}
	}

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestAllCmdsErrorWhenGettingFromDB(t *testing.T) {

	guard := monkey.Patch(model.GetContainer,
		func(Team string, Channel string, Name string) (model.Container, error) {
			return model.Container{}, fmt.Errorf("simulated error")
		})

	teamName := "TestTeamList"
	channelName := "TestChannel"
	userName := "user"

	mockResponse, patchReply := createMockReply(t, "simulated error")
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, nil)

	for _, command := range commands {
		if strings.Contains(command.Usage, "<container-name>") {
			command.Handler(mockRequest, mockResponse)
		}
	}

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
	guard.Unpatch()

}

func TestNilGetEvent(t *testing.T) {
	event := getEvent(nil)
	assert.ObjectsAreEqual(ClaimrEvent{}, event)
}

func TestGetEventFromNSLopesEvent(t *testing.T) {
	var message slack.MessageEvent
	message.Team = "Team"
	message.Channel = "Channel"
	message.User = "User"
	request := slacker.NewRequest(nil, &message, &proper.Properties{})
	event := getEvent(request)
	assert.ObjectsAreEqual(ClaimrEvent{}, event)
}

func TestGetEventText(t *testing.T) {
	text := "Text"
	var message slack.MessageEvent
	message.Team = "Team"
	message.Channel = "Channel"
	message.User = "User"
	message.Text = text
	request := slacker.NewRequest(nil, &message, &proper.Properties{})

	assert.Equal(t, text, GetEventText(request))
}

func TestNonAdminTryAccessAdminOnlyCmds(t *testing.T) {
	teamName := "TestPurge"
	channelName := "TestChannel"
	userName := "NotAAdmin"

	mockResponse, patchReply := createMockReply(t, Messages["admin-only"])
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, _ := createMockRequest(t, nil)

	for _, command := range commands {
		if strings.Contains(command.Description, "admin-only") {
			command.Handler(mockRequest, mockResponse)
		}
	}

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
}
