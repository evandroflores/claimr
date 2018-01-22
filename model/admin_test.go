package model

import (
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/nlopes/slack"
	"github.com/shomali11/slacker"
	"github.com/stretchr/testify/assert"
)

func TestLoadingAdmins(t *testing.T) {

	expectedAdmins := []Admin{
		{ID: "U33333", RealName: "Admin 1"},
		{ID: "U44444", RealName: "Owner"},
	}

	var mockClient *slack.Client

	patchGetUsers := monkey.PatchInstanceMethod(reflect.TypeOf(mockClient), "GetUsers",
		func(*slack.Client) ([]slack.User, error) {
			users := []slack.User{
				{
					ID:       "U11111",
					RealName: "User 1",
					IsAdmin:  false,
					IsOwner:  false,
				},
				{
					ID:       "U22222",
					RealName: "User 2",
					IsAdmin:  false,
					IsOwner:  false,
				},
				{
					ID:       "U33333",
					RealName: "Admin 1",
					IsAdmin:  true,
					IsOwner:  false,
				},
				{
					ID:       "U44444",
					RealName: "Owner",
					IsAdmin:  true,
					IsOwner:  true,
				},
			}

			return users, nil
		})

	patchNewRTM := monkey.PatchInstanceMethod(reflect.TypeOf(mockClient), "NewRTM", func(*slack.Client) *slack.RTM { return nil })

	bot := slacker.NewClient("fake")
	LoadAdmins(bot)
	assert.Equal(t, expectedAdmins, Admins)

	patchNewRTM.Unpatch()
	patchGetUsers.Unpatch()

}
