package cmd

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	"github.com/stretchr/testify/assert"
)

func TestTryToRemoveDirect(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, "this look like a direct message. Containers are related to a channels")
	patchGetEvent := createMockEvent(t, "team", "DIRECT", "user")

	remove(new(slacker.Request), mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
}

func TestTryToRemoveNoName(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, "can not continue without a container name 🙄")
	patchGetEvent := createMockEvent(t, "team", "channel", "user")
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": ""})

	remove(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestTryToRemoveInexistentContainer(t *testing.T) {
	containerName := "container-inexistent"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf("I couldn't find container `%s` on <#%s>.", containerName, channelName))
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

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf("Can't remove. Container `%s` is in used by <@%s> since _%s_.",
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

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf("Only who created the container `%s` can remove it. Please check with <@%s>.",
		containerName, anotherUser))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	remove(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestRemoveError(t *testing.T) {

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

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf("Container `%s` removed.", containerName))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	remove(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}
