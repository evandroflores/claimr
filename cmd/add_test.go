package cmd

import (
	"testing"

	"github.com/shomali11/slacker"
)

func TestAddDirect(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, "this look like a direct message. Containers are related to a channels")
	patchGetEvent := createMockEvent(t, "team", "DIRECT", "user")

	add(new(slacker.Request), mockResponse)

	defer patchReply.Unpatch()
	defer patchGetEvent.Unpatch()
}

func TestAddNoName(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, "can not continue without a container name 🙄")
	patchGetEvent := createMockEvent(t, "team", "channel", "user")
	mockRequest, patchParam := createMockRequest(t, map[string]string{"container-name": ""})

	add(mockRequest, mockResponse)

	defer patchReply.Unpatch()
	defer patchGetEvent.Unpatch()
	defer patchParam.Unpatch()
}
