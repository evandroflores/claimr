package cmd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/model"
	"github.com/stretchr/testify/assert"
)

func TestTryToAddBigName(t *testing.T) {

	containerName := "lorem-ipsum-container-big-name"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	mockResponse, patchReply := createMockReply(t, Messages["field-name-too-big"])
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	add(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestTryToAddExistentContainer(t *testing.T) {
	containerName := "my-container-exists"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName}
	err := container.Add()
	defer container.Delete()

	assert.NoError(t, err)

	mockResponse, patchReply := createMockReply(t, Messages["same-name"])
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	add(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestAddError(t *testing.T) {

	guard := monkey.PatchInstanceMethod(reflect.TypeOf(model.Container{}), "Add",
		func(container model.Container) error {
			return fmt.Errorf("simulated error")
		})
	defer guard.Unpatch()

	containerName := "my-container-ok"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	mockResponse, patchReply := createMockReply(t, "simulated error")
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	add(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestAddContainer(t *testing.T) {
	containerName := "my-container-ok"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	msg := fmt.Sprintf(Messages["added-to-channel"], containerName, channelName)
	mockResponse, patchReply := createMockReply(t, msg)
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	defer func() {
		container, _ := model.GetContainer(teamName, channelName, containerName)
		container.Delete()
	}()

	add(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestAddContainerWithUserMention(t *testing.T) {
	containerName := "<@user>"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	msg := Messages["shouldnt-mention-user"]
	mockResponse, patchReply := createMockReply(t, msg)
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	defer func() {
		container, _ := model.GetContainer(teamName, channelName, containerName)
		container.Delete()
	}()

	add(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestAddContainerWithChannelMention(t *testing.T) {
	containerName := "<#CHANNEL>"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	msg := Messages["shouldnt-mention-channel"]
	mockResponse, patchReply := createMockReply(t, msg)
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	defer func() {
		container, _ := model.GetContainer(teamName, channelName, containerName)
		container.Delete()
	}()

	add(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}
