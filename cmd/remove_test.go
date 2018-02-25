package cmd

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/messages"
	"github.com/evandroflores/claimr/model"
	"github.com/stretchr/testify/assert"
)

func TestTryToRemoveContainerNotFound(t *testing.T) {
	containerName := "container-not-found"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(messages.Messages["container-not-found-on-channel"], containerName, channelName))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	remove(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestTryToRemoveInUseContainer(t *testing.T) {
	containerName := "container"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"
	anotherUser := "anotherUser"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, InUseBy: anotherUser}
	err := container.Add()

	defer container.Delete()
	assert.NoError(t, err)

	containerFromDB, err := model.GetContainer(teamName, channelName, containerName)
	assert.NoError(t, err)

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(messages.Messages["container-in-use-by-this"],
		containerName, anotherUser, containerFromDB.UpdatedAt.Format(time.RFC1123)))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	remove(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestTryToRemoveAContainerCreatedByAnotherUser(t *testing.T) {
	containerName := "container"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"
	anotherUser := "anotherUser"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, CreatedByUser: anotherUser}
	err := container.Add()
	defer container.Delete()

	assert.NoError(t, err)

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(messages.Messages["only-owner-can-remove"],
		containerName, anotherUser))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	remove(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestRemoveErrorWhenDeleting(t *testing.T) {

	guard := monkey.PatchInstanceMethod(reflect.TypeOf(model.Container{}), "Delete",
		func(container model.Container) error {
			return fmt.Errorf("simulated error")
		})

	containerName := "my-container-delete"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, CreatedByUser: userName}
	err := container.Add()
	defer container.Delete()
	assert.NoError(t, err)

	mockResponse, patchReply := createMockReply(t, "simulated error")
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	remove(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
	guard.Unpatch()
}

func TestRemoving(t *testing.T) {
	containerName := "my-container-delete"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, CreatedByUser: userName}
	err := container.Add()
	assert.NoError(t, err)

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(messages.Messages["container-removed"], containerName))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	remove(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}
