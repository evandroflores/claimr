package cmd

import (
	"testing"

	"github.com/evandroflores/claimr/messages"
)

func TestDefaultCommand(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, messages.Messages["command-not-found"])
	mockRequest, _ := createMockRequest(t, nil)

	Default(mockRequest, mockResponse)

	patchReply.Unpatch()
}
