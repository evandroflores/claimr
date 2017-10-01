package main

import (
	"os"
	"testing"

	"github.com/bouk/monkey"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestEnvironmentToken(t *testing.T) {
	currentEnv := os.Getenv("CLAIMR_TOKEN")
	os.Unsetenv("CLAIMR_TOKEN")
	defer func() { os.Setenv("CLAIMR_TOKEN", currentEnv) }()

	wantMsg := "Claimr slack bot token unset. Set CLAIMR_TOKEN to continue."

	mockLogFatal := func(msg ...interface{}) {
		assert.Equal(t, wantMsg, msg[0])
		panic("log.Fatal called")
	}
	patchLog := monkey.Patch(log.Fatal, mockLogFatal)
	defer patchLog.Unpatch()
	assert.PanicsWithValue(t, "log.Fatal called", main, "log.Fatal was not called")
}
