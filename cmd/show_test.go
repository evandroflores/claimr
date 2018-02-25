package cmd

import (
	"fmt"
	"testing"
	"time"

	"github.com/evandroflores/claimr/messages"
	"github.com/evandroflores/claimr/model"
	"github.com/stretchr/testify/assert"
)

func TestTryToShowContainerNotFound(t *testing.T) {
	containerName := "container-not-found"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(messages.Messages["container-not-found-on-channel"], containerName, channelName))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	show(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestTryToShowInUse(t *testing.T) {
	containerName := "container-in-use"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"
	anotherUser := "anotherUser"
	reason := "testing"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, InUseBy: anotherUser, InUseForReason: reason, CreatedByUser: userName}
	err := container.Add()

	defer container.Delete()
	assert.NoError(t, err)

	containerFromDB, err := model.GetContainer(teamName, channelName, containerName)
	assert.NoError(t, err)

	text := fmt.Sprintf(messages.Messages["container-created-by"], containerName, container.CreatedByUser)
	text += fmt.Sprintf(messages.Messages["container-in-use-by-w-reason"], anotherUser, " for "+reason, containerFromDB.UpdatedAt.Format(time.RFC1123))

	mockResponse, patchReply := createMockReply(t, text)
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	show(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestTryToShowAvailable(t *testing.T) {
	containerName := "container-available"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, CreatedByUser: userName}
	err := container.Add()

	defer container.Delete()
	assert.NoError(t, err)

	text := fmt.Sprintf("Container `%s`.\nCreated by <@%s>.\n_Available_",
		containerName, container.CreatedByUser)

	mockResponse, patchReply := createMockReply(t, text)
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	show(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}
