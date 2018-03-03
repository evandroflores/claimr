package cmd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/messages"
	"github.com/evandroflores/claimr/model"
	"github.com/stretchr/testify/assert"
)

func TestTryToFreeContainerNotFound(t *testing.T) {
	containerName := "container-not-found"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(messages.Get("container-not-found-on-channel"), containerName, channelName))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	free(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestTryToFreeAContainerInUseByAnotherUser(t *testing.T) {
	containerName := "container"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"
	anotherUser := "anotherUser"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, InUseBy: anotherUser}
	err := container.Add()
	defer container.Delete()

	assert.NoError(t, err)

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(messages.Get("container-in-use-by-other"), containerName))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	free(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestFreeError(t *testing.T) {

	guard := monkey.PatchInstanceMethod(reflect.TypeOf(model.Container{}), "Update",
		func(container model.Container) error {
			return fmt.Errorf("simulated error")
		})

	containerName := "my-container-free"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, InUseBy: userName}
	err := container.Add()
	defer container.Delete()
	assert.NoError(t, err)

	mockResponse, patchReply := createMockReply(t, messages.Get("fail-to-update"))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	free(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
	guard.Unpatch()
}

func TestFreeing(t *testing.T) {
	containerName := "my-container-free"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, InUseBy: userName}
	err := container.Add()
	assert.NoError(t, err)
	defer container.Delete()

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(messages.Get("container-free"), containerName))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	free(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}
