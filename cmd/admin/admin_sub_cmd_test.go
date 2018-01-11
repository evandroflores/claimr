package admin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdminCmdCommandList(t *testing.T) {
	usageExpected := []string{
		"command-list",
		"log-level `level`",
		"purge",
	}
	commands := CommandList()

	assert.Len(t, commands, len(usageExpected))

	usageActual := []string{}

	for _, command := range commands {
		usageActual = append(usageActual, command.Usage)
	}

	assert.Subset(t, usageExpected, usageActual)
}
