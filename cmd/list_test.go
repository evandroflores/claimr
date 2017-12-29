package cmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/model"
	"github.com/stretchr/testify/assert"
)

func TestListError(t *testing.T) {

	guard := monkey.Patch(model.GetContainers,
		func(Team string, Channel string) ([]model.Container, error) {
			return []model.Container{}, fmt.Errorf("simulated error")
		})

	teamName := "TestTeamList"
	channelName := "TestChannel"
	userName := "user"

	mockResponse, patchReply := createMockReply(t, "Fail to list containers.")
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, nil)

	list(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
	guard.Unpatch()
}

func TestListNoContainers(t *testing.T) {

	teamName := "TestTeamList"
	channelName := "TestChannel"
	userName := "user"

	mockResponse, patchReply := createMockReply(t, "No containers to list.")
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, nil)

	list(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestListAvailableContainers(t *testing.T) {
	teamName := "TestTeamList"
	channelName := "TestChannel"
	userName := "user"

	expected := []string{"Here is a list of containers for this channel:"}

	for i := 0; i < 5; i++ {
		containerName := fmt.Sprintf("container_%d", i)
		container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName}
		err := container.Add()
		defer container.Delete()
		assert.NoError(t, err)

		line := fmt.Sprintf("`%s`\t%s %s", container.Name, "_available_", "")
		expected = append(expected, line)
	}

	mockResponse, patchReply := createMockReply(t, strings.Join(expected, "\n"))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, nil)

	list(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestListInUseContainers(t *testing.T) {
	teamName := "TestTeamList"
	channelName := "TestChannel"
	userName := "user"

	expected := []string{"Here is a list of containers for this channel:"}

	for i := 0; i < 5; i++ {
		containerName := fmt.Sprintf("container_%d", i)
		container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, InUseBy: userName}
		err := container.Add()
		defer container.Delete()
		assert.NoError(t, err)

		line := fmt.Sprintf("`%s`\t%s %s", container.Name, "in use", "")
		expected = append(expected, line)
	}

	mockResponse, patchReply := createMockReply(t, strings.Join(expected, "\n"))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, nil)

	list(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestListInUseWithReasonContainers(t *testing.T) {
	teamName := "TestTeamList"
	channelName := "TestChannel"
	userName := "user"
	reason := "tests"
	expected := []string{"Here is a list of containers for this channel:"}

	for i := 0; i < 5; i++ {
		containerName := fmt.Sprintf("container_%d", i)
		container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, InUseBy: userName, InUseForReason: reason}
		err := container.Add()
		defer container.Delete()
		assert.NoError(t, err)

		line := fmt.Sprintf("`%s`\t%s %s", container.Name, "in use", fmt.Sprintf("- %s", reason))
		expected = append(expected, line)
	}

	mockResponse, patchReply := createMockReply(t, strings.Join(expected, "\n"))
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)
	mockRequest, patchParam := createMockRequest(t, nil)

	list(mockRequest, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}
