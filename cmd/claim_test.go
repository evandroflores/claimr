package cmd

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	"github.com/stretchr/testify/assert"
)

func TestTryToClaimInexistentContainer(t *testing.T) {
	containerName := "container-inexistent"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(Messages["container-not-found-on-channel"], containerName, channelName))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	claim(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestTryToClaimAContainerInUseByAnotherUser(t *testing.T) {
	containerName := "container"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"
	anotherUser := "anotherUser"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, InUseBy: anotherUser}
	err := container.Add()
	defer container.Delete()

	assert.NoError(t, err)

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(Messages["container-in-use"], containerName))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	claim(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestClaimErrorWhenUpdate(t *testing.T) {

	guard := monkey.PatchInstanceMethod(reflect.TypeOf(model.Container{}), "Update",
		func(container model.Container) error {
			return fmt.Errorf("simulated error")
		})

	containerName := "my-container-claim"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName}
	err := container.Add()
	defer container.Delete()
	assert.NoError(t, err)

	mockResponse, patchReply := createMockReply(t, Messages["fail-to-update"])
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName, "reason": ""})

	claim(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
	guard.Unpatch()
}

func TestClaiming(t *testing.T) {
	containerName := "my-container"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName}
	err := container.Add()
	assert.NoError(t, err)
	defer container.Delete()

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(Messages["container-claimed"], containerName, userName))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	claim(mockRequest, mockResponse)

	containerFromDB, err2 := model.GetContainer(teamName, channelName, containerName)
	assert.NoError(t, err2)
	assert.Equal(t, userName, containerFromDB.InUseBy)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestClaimingWithOneWordReason(t *testing.T) {
	containerName := "my-container"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"
	reason := "tests"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName}
	err := container.Add()
	assert.NoError(t, err)
	defer container.Delete()

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(Messages["container-claimed"], containerName, userName))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName, "reason": reason})
	patchGetEventText := monkey.Patch(GetEventText,
		func(request *slacker.Request) string {
			return reason
		})

	claim(mockRequest, mockResponse)

	containerFromDB, err2 := model.GetContainer(teamName, channelName, containerName)
	assert.NoError(t, err2)
	assert.Equal(t, userName, containerFromDB.InUseBy)
	assert.Equal(t, reason, containerFromDB.InUseForReason)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
	patchGetEventText.Unpatch()
}

func TestClaimingWithMultiwordsReason(t *testing.T) {
	containerName := "my-container"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"
	reason := "testing my container"
	firstWord := strings.Split(reason, " ")[0]

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName}
	err := container.Add()
	assert.NoError(t, err)
	defer container.Delete()

	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(Messages["container-claimed"], containerName, userName))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName, "reason": firstWord})

	patchGetEventText := monkey.Patch(GetEventText,
		func(request *slacker.Request) string {
			return reason
		})

	claim(mockRequest, mockResponse)

	containerFromDB, err2 := model.GetContainer(teamName, channelName, containerName)
	assert.NoError(t, err2)
	assert.Equal(t, userName, containerFromDB.InUseBy)
	assert.Equal(t, reason, containerFromDB.InUseForReason)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
	patchGetEventText.Unpatch()
}
