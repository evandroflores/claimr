package cmd

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/model"
	"github.com/nlopes/slack"
	"github.com/shomali11/slacker"
	"github.com/stretchr/testify/assert"
)

func TestRefreshAdmins(t *testing.T) {
	expected1stRun := []model.Admin{
		{ID: "U44444", RealName: "Owner"},
	}

	expected2ndRun := []model.Admin{
		{ID: "U33333", RealName: "Admin 1"},
		{ID: "U44444", RealName: "Owner"},
	}

	var mockClient *slack.Client

	patchGetUsers1st := monkey.PatchInstanceMethod(reflect.TypeOf(mockClient), "GetUsers",
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
					IsAdmin:  false,
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

	bot := slacker.NewClient("fake")
	model.LoadAdmins(bot)
	assert.Equal(t, expected1stRun, model.Admins)
	patchGetUsers1st.Unpatch()

	patchGetUsers2nd := monkey.PatchInstanceMethod(reflect.TypeOf(mockClient), "GetUsers",
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

	patchGetEvent := createMockEvent(t, "test", "test", os.Getenv("CLAIMR_SUPERUSER"))
	mockResponse, patchReply := createMockReply(t, fmt.Sprintf(Messages["x-admin-loaded"], len(expected2ndRun)))

	refreshAdmins(nil, mockResponse)

	assert.Equal(t, expected2ndRun, model.Admins)

	patchGetEvent.Unpatch()
	patchReply.Unpatch()
	patchNewRTM.Unpatch()
	patchGetUsers2nd.Unpatch()
}
