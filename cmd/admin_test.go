package cmd

import (
	"testing"

	adm "github.com/evandroflores/claimr/cmd/admin"
)

func TestNonAdminTryToCallAdminCommand(t *testing.T) {
	teamName := "TestTeamList"
	channelName := "TestChannel"
	userName := "user"

	mockResponse, patchReply := createMockReply(t, "Command available only for admins. 🛑")
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, nil)

	admin(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestAdminTryToCallAdminCommandWithoutCommand(t *testing.T) {
	teamName := "TestTeamList"
	channelName := "TestChannel"
	userName := "TESTSUPERUSER"

	mockResponse, patchReply := createMockReply(t, "Command not found, Type @claimr admin command-list` for valid admin sub commands.")
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, nil)

	admin(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestAdminCallAdminSubCommandList(t *testing.T) {
	teamName := "TestTeamList"
	channelName := "TestChannel"
	userName := "TESTSUPERUSER"

	mockResponse, patchReply := createMockReply(t, adm.GenerateCommandHelp())
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"sub-command": "command-list"})

	admin(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}
