package model

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/nlopes/slack"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
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

func TestLoadingAdminsError(t *testing.T) {
	var mockClient *slack.Client

	expectedMsg := "Simulated Error"

	patchGetUsers := monkey.PatchInstanceMethod(reflect.TypeOf(mockClient), "GetUsers",
		func(*slack.Client) ([]slack.User, error) {
			return nil, fmt.Errorf(expectedMsg)
		})

	mockLogErrorf := func(format string, args ...interface{}) {
		assert.Equal(t, expectedMsg, fmt.Sprintf(format, args))
	}

	patchLog := monkey.Patch(log.Warnf, mockLogErrorf)

	bot := slacker.NewClient("fake")
	LoadAdmins(bot)

	patchLog.Unpatch()
	patchGetUsers.Unpatch()
}
