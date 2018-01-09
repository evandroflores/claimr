package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/evandroflores/claimr/cmd"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestEnvironmentToken(t *testing.T) {
	currentEnv := os.Getenv("CLAIMR_TOKEN")
	os.Unsetenv("CLAIMR_TOKEN")
	defer func() { os.Setenv("CLAIMR_TOKEN", currentEnv) }()

	expectedMsg := "Claimr slack bot token unset. Set CLAIMR_TOKEN to continue."

	mockLogFatal := func(msg ...interface{}) {
		assert.Equal(t, expectedMsg, msg[0])
		panic("log.Fatal called")
	}
	patchLog := monkey.Patch(log.Fatal, mockLogFatal)
	defer patchLog.Unpatch()
	assert.PanicsWithValue(t, "log.Fatal called", main, "log.Fatal was not called")
}

func TestSlackerCommands(t *testing.T) {
	var mockSlacker *slacker.Slacker

	var commands []model.Command
	patchCommand := monkey.PatchInstanceMethod(reflect.TypeOf(mockSlacker), "Command",
		func(slacker *slacker.Slacker, usage string, description string, handler func(request *slacker.Request, response slacker.ResponseWriter)) {
			if usage == "help" {
				return //ignoring help slacker command
			}
			commands = append(commands, model.Command{Usage: usage, Description: description, Handler: handler})
		})
	patchListen := monkey.PatchInstanceMethod(reflect.TypeOf(mockSlacker), "Listen", func(*slacker.Slacker) error {
		assert.CallerInfo()
		return nil
	})

	main()

	for i, command := range cmd.CommandList() {
		assert.Equal(t, command.Description, commands[i].Description)
	}
	patchCommand.Unpatch()
	patchListen.Unpatch()
}

func TestSlackerListenError(t *testing.T) {
	var mockSlacker *slacker.Slacker

	patchCommand := monkey.PatchInstanceMethod(reflect.TypeOf(mockSlacker), "Command",
		func(slacker *slacker.Slacker, usage string, description string, handler func(request *slacker.Request, response slacker.ResponseWriter)) {
		})

	patchListen := monkey.PatchInstanceMethod(reflect.TypeOf(mockSlacker), "Listen", func(*slacker.Slacker) error {
		return fmt.Errorf("Simulated Error")
	})

	expectedMsg := fmt.Errorf("Simulated Error")

	mockLogFatal := func(msg ...interface{}) {
		assert.Equal(t, expectedMsg, msg[0])
		panic("log.Fatal called")
	}
	patchLog := monkey.Patch(log.Fatal, mockLogFatal)

	defer patchCommand.Unpatch()
	defer patchListen.Unpatch()
	defer patchLog.Unpatch()

	assert.PanicsWithValue(t, "log.Fatal called", main, "log.Fatal called was not called")
}
