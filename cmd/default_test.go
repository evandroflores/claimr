package cmd

import "testing"

func TestDefaultCommand(t *testing.T) {
	mockResponse, patchReply := createMockReply(t, "Not sure what you are asking for. Type `@claimr help` for valid commands.")
	mockRequest, _ := createMockRequest(t, nil)

	Default(mockRequest, mockResponse)

	patchReply.Unpatch()
}
