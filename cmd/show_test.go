package cmd

import (
	"fmt"
	"testing"
	"time"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	"github.com/stretchr/testify/assert"
)

func TestTryToShowDirect(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, "this look like a direct message. Containers are related to a channels")
	patchGetEvent := createMockEvent(t, "team", "DIRECT", "user")

	show(new(slacker.Request), mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
}

func TestTryToShowNoName(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, "can not continue without a container name 🙄")
	patchGetEvent := createMockEvent(t, "team", "channel", "user")
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": ""})

	show(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestShowError(t *testing.T) {

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

	show(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
	guard.Unpatch()
}

func TestTryToShowInexistentContainer(t *testing.T) {
	containerName := "container-inexistent"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf("I couldn't find the container `%s` on <#%s>.", containerName, channelName))
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

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, InUseBy: anotherUser, InUseForReason: reason}
	err := container.Add()

	defer container.Delete()
	assert.NoError(t, err)

	containerFromDB, err := model.GetContainer(teamName, channelName, containerName)
	assert.NoError(t, err)

	text := fmt.Sprintf("Container `%s`.\nCreated by <@%s>.\nIn use by <@%s> for %s since _%s_.",
		containerName, container.CreatedByUser, anotherUser, reason, containerFromDB.UpdatedAt.Format(time.RFC1123))

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

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName}
	err := container.Add()

	defer container.Delete()
	assert.NoError(t, err)

	containerFromDB, err := model.GetContainer(teamName, channelName, containerName)
	assert.NoError(t, err)

	text := fmt.Sprintf("Container `%s`.\nCreated by <@%s>.\n_Available_ since _%s_.",
		containerName, container.CreatedByUser, containerFromDB.UpdatedAt.Format(time.RFC1123))

	mockResponse, patchReply := createMockReply(t, text)
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	show(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}
