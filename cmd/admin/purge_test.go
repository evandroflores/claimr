package admin

import (
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	"github.com/stretchr/testify/assert"
)

func TestPurgeContainers(t *testing.T) {
	var mockResponse *slacker.Response

	expected := "10 Container rows purged"

	patchReply := monkey.PatchInstanceMethod(reflect.TypeOf(mockResponse), "Reply",
		func(response *slacker.Response, msg string) {
			assert.Equal(t, expected, msg)
		})

	_ = monkey.PatchInstanceMethod(reflect.TypeOf(mockResponse), "Typing",
		func(response *slacker.Response) {})

	for i := 1; i <= 10; i++ {
		container := model.Container{TeamID: "TestPurge", ChannelID: "TestChannel", Name: "container", CreatedByUser: "user"}
		container.Add()
		container.Delete()
	}
	database.DB.Unscoped().Where("deleted_at is not null and team_id <> 'TestPurge'").Delete(&model.Container{})

	purge(nil, mockResponse)

	patchReply.Unpatch()
}
