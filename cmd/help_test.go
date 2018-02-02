package cmd

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/bouk/monkey"
	"github.com/shomali11/slacker"
	"github.com/stretchr/testify/assert"
)

func TestCallHelpAsRegularUser(t *testing.T) {
	var mockResponse *slacker.Response

	expected := GenerateCommandHelp(false)

	assert.False(t, strings.Contains(expected, "admin-only"), expected)

	patchGetEvent := createMockEvent(t, "team", "channel", "notAdmin")
	patchReply := monkey.PatchInstanceMethod(reflect.TypeOf(mockResponse), "Reply",
		func(response *slacker.Response, msg string) {
			assert.Equal(t, expected, msg)
		})

	_ = monkey.PatchInstanceMethod(reflect.TypeOf(mockResponse), "Typing",
		func(response *slacker.Response) {})

	Help(nil, mockResponse)

	patchGetEvent.Unpatch()
	patchReply.Unpatch()
}

func TestCallHelpAsAdmin(t *testing.T) {
	var mockResponse *slacker.Response

	expected := GenerateCommandHelp(true)

	assert.True(t, strings.Contains(expected, "admin-only"), expected)

	patchGetEvent := createMockEvent(t, "team", "channel", os.Getenv("CLAIMR_SUPERUSER"))
	patchReply := monkey.PatchInstanceMethod(reflect.TypeOf(mockResponse), "Reply",
		func(response *slacker.Response, msg string) {
			assert.Equal(t, expected, msg)
		})

	_ = monkey.PatchInstanceMethod(reflect.TypeOf(mockResponse), "Typing",
		func(response *slacker.Response) {})

	Help(nil, mockResponse)

	patchGetEvent.Unpatch()
	patchReply.Unpatch()
}
