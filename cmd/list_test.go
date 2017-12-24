package cmd

import (
	"testing"

	"github.com/shomali11/slacker"
)

func TestTryToListDirect(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, "this look like a direct message. Containers are related to a channels")
	patchGetEvent := createMockEvent(t, "team", "DIRECT", "user")

	list(new(slacker.Request), mockResponse)

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
}
