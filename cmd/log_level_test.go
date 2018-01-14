package cmd

import (
	"fmt"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestTryToChangeLogLevelWithoutParameter(t *testing.T) {
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := os.Getenv("CLAIMR_SUPERUSER")

	currentLogLevel := log.GetLevel().String()
	message := fmt.Sprintf("No log level provided, keeping in `%s`", currentLogLevel)
	mockResponse, patchReply := createMockReply(t, message)
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)

	mockRequest, patchParam := createMockRequest(t, map[string]string{"level": ""})

	changeLogLevel(mockRequest, mockResponse)

	assert.Equal(t, currentLogLevel, log.GetLevel().String())

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}

func TestTryToChangeLogLevelToUnknownLevel(t *testing.T) {
	teamName := "TestTeam"
	channelName := "TestChannel"
	userName := os.Getenv("CLAIMR_SUPERUSER")

	currentLogLevel := log.GetLevel().String()
	unknownLogLevel := "unknown"
	message := fmt.Sprintf("not a valid logrus Level: \"%s\"", unknownLogLevel)
	mockResponse, patchReply := createMockReply(t, message)
	patchGetEvent := createMockEvent(t, teamName, channelName, userName)

	mockRequest, patchParam := createMockRequest(t, map[string]string{"level": unknownLogLevel})

	changeLogLevel(mockRequest, mockResponse)

	assert.Equal(t, currentLogLevel, log.GetLevel().String())

	patchReply.Unpatch()
	patchGetEvent.Unpatch()
	patchParam.Unpatch()
}
