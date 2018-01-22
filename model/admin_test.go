package model

import (
	"testing"

	"github.com/bouk/monkey"
	"github.com/nlopes/slack"
	"github.com/shomali11/slacker"
	"github.com/stretchr/testify/assert"
)

func TestLoadingAdmins(t *testing.T) {

	bot := slacker.NewClient("fake")

	expectedAdmins := []Admin{
		{ID: "U33333", RealName: "Admin 1"},
		{ID: "U44444", RealName: "Owner"},
	}

	patchGetUsers := monkey.Patch(bot.Client.GetUsers,
		func() ([]slack.User, error) {
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

	LoadAdmins(bot)
	assert.Equal(t, expectedAdmins, Admins)
	patchGetUsers.Unpatch()
}
