package cmd

import (
	"testing"

	"github.com/shomali11/slacker"
)

func TestAddDirect(t *testing.T) {
	req := new(slacker.Request)

	mockResponse, patchReply := createMockReply(t, "this look like a direct message. Containers are related to a channels")
	patchGetEvent := createMockEvent(t, "team", "DIRECT", "user")
	patchParam := createMockParam(t, "container-name", "test")

	add(req, mockResponse)

	defer patchReply.Unpatch()
	defer patchGetEvent.Unpatch()
	defer patchParam.Unpatch()
}
