package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/messages"
	"github.com/evandroflores/claimr/model"
)

func TestPurgeContainers(t *testing.T) {
	containerName := "container-purge"
	teamName := "TestPurge"
	channelName := "TestChannel"
	userName := os.Getenv("CLAIMR_SUPERUSER")

	testAmount := 10
	expected := fmt.Sprintf(messages.Get("x-purged"), testAmount)

	mockResponse, patchReply := createMockReply(t, expected)
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)

	for i := 1; i <= testAmount; i++ {
		container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, CreatedByUser: userName}
		container.Add()
		container.Delete()
	}
	database.DB.Unscoped().Where("deleted_at is not null and team_id <> 'TestPurge'").Delete(&model.Container{})

	purge(nil, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
}
