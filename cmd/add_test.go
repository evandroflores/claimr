package cmd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	"github.com/stretchr/testify/assert"
)

func TestTryToAddDirect(t *testing.T) {
	// will fail if the message is different
	mockResponse, patchReply := createMockReply(t, "this look like a direct message. Containers are related to a channels")
	patchGetEvent := createMockEvent(t, "team", "DIRECT", "user")

	add(new(slacker.Request), mockResponse)

	defer patchReply.Unpatch()
	defer patchGetEvent.Unpatch()
}

func TestTryToAddNoName(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, "can not continue without a container name 🙄")
	patchGetEvent := createMockEvent(t, "team", "channel", "user")
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": ""})

	add(mockRequest, mockResponse)

	defer patchReply.Unpatch()
	defer patchGetEvent.Unpatch()
	defer patchParam.Unpatch()
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

	mockResponse, patchReply := createMockReply(t, "There is a container with the same name on this channel. Try a different one.")
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	add(mockRequest, mockResponse)

	defer patchReply.Unpatch()
	defer patchGetEvent.Unpatch()
	defer patchParam.Unpatch()
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

	defer patchReply.Unpatch()
	defer patchGetEvent.Unpatch()
	defer patchParam.Unpatch()
}

func TestAddContainer(t *testing.T) {
	containerName := "my-container-ok"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := "user"

	msg := fmt.Sprintf("Container `%s` added to channel <#%s>.", containerName, channelName)
	mockResponse, patchReply := createMockReply(t, msg)
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": containerName})

	defer func() {
		container, _ := model.GetContainer(teamName, channelName, containerName)
		container.Delete()
	}()

	add(mockRequest, mockResponse)

	defer patchReply.Unpatch()
	defer patchGetEvent.Unpatch()
	defer patchParam.Unpatch()
}
