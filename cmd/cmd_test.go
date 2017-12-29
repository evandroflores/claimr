package cmd

import (
	"testing"

	"reflect"

	"github.com/bouk/monkey"
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
	mockResponse, patchReply := createMockReply(t, "No pancakes for you! 🥞")

	notImplemented(new(slacker.Request), mockResponse)

	defer patchReply.Unpatch()
}

func TestCmdCommandList(t *testing.T) {
	usageExpected := []string{
		"add <container-name>",
		"claim <container-name> <reason>",
		"free <container-name>",
		"list",
		"remove <container-name>",
		"show <container-name>",
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
	assert.Error(t, err, "this look like a direct message. Containers are related to a channels")
}

func TestAllCmdsCheckingDirect(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, "this look like a direct message. Containers are related to a channels")
	patchGetEvent := createMockEvent(t, "team", "DIRECT", "user")

	for _, command := range commands {
		command.Handler(new(slacker.Request), mockResponse)
	}

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
}
