package cmd

import "testing"

func TestDefaultCommand(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, Messages["command-not-found"])
	mockRequest, _ := createMockRequest(t, nil)

	Default(mockRequest, mockResponse)

	patchReply.Unpatch()
}
