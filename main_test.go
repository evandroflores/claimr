package main

import (
	"os"
	"os/exec"
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestMustHaveTokenEnv(t *testing.T) {
	currentToken := os.Getenv("CLAIMR_TOKEN")
	os.Unsetenv("CLAIMR_TOKEN")
	defer func() { os.Setenv("CLAIMR_TOKEN", currentToken) }()

	cmd := exec.Command("make", "run")
	//cmd.Env = append(os.Environ(), "TEST_MUST_ENV=1")
	err := cmd.Run()
	fmt.Print(err.(*exec.ExitError))
	//if e, ok := err.(*exec.ExitError); ok && !e.Success() {
	//	return
	//}
	assert.EqualError(t, err, "Claimr slack bot token unset. Set CLAIMR_TOKEN to continue.")
}
