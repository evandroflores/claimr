package cmd

import (
	"testing"

	"reflect"

	"github.com/bouk/monkey"
	"github.com/shomali11/slacker"
	"github.com/stretchr/testify/assert"
)

func TestCmdNotImplemented(t *testing.T) {
	expectedMsg := "No pancakes for you! 🥞"

	mockResponse := func(response *slacker.Response, msg string) {
		assert.Equal(t, expectedMsg, msg)
	}

	var response *slacker.Response
	patchResponse := monkey.PatchInstanceMethod(reflect.TypeOf(response), "Reply", mockResponse)

	notImplemented(nil, response)

	defer patchResponse.Unpatch()
}

func TestCmdCommandList(t *testing.T) {
	usageExpected := []string{
		"add <container-name>",
		"claim <container-name> <reason>",
		"free <container-name>",
		"list",
		"remove <container-name>",
		"show <container-name>",
	}
	commands := CommandList()

	assert.Len(t, commands, len(usageExpected))

	usageActual := []string{}

	for _, command := range commands {
		usageActual = append(usageActual, command.Usage)
	}

	assert.Subset(t, usageExpected, usageActual)
}

func TestCmdNotDirect(t *testing.T) {
	isDirect, err := checkDirect("CHANNEL")
	assert.False(t, isDirect)
	assert.NoError(t, err)
}

func TestCmdDirect(t *testing.T) {
	isDirect, err := checkDirect("DIRECT")
	assert.True(t, isDirect)
	assert.Error(t, err, "this look like a direct message. Containers are related to a channels")
}
