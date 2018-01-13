package cmd

import (
	"os"
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/evandroflores/slacker"
	"github.com/stretchr/testify/assert"
)

func TestPurgeContainers(t *testing.T) {
	var mockResponse *slacker.Response

	containerName := "container-purge"
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := os.Getenv("CLAIMR_SUPERUSER")

	expected := "10 Container rows purged"

	patchReply := monkey.PatchInstanceMethod(reflect.TypeOf(mockResponse), "Reply",
		func(response *slacker.Response, msg string) {
			assert.Equal(t, expected, msg)
		})

	_ = monkey.PatchInstanceMethod(reflect.TypeOf(mockResponse), "Typing",
		func(response *slacker.Response) {})

	patchGetEvent := createMockEvent(t, teamName, channelName, userName)

	for i := 1; i <= 10; i++ {
		container := model.Container{TeamID: teamName, ChannelID: channelName, Name: containerName, CreatedByUser: userName}
		container.Add()
		container.Delete()
	}
	database.DB.Unscoped().Where("deleted_at is not null and team_id <> 'TestPurge'").Delete(&model.Container{})

	purge(nil, mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
}
