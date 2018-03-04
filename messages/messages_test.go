package messages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInvalidKey(t *testing.T) {
	assert.Equal(t, messagesMap["invalid-message-key"], Get("?"))
}

func TestGetValidKey(t *testing.T) {
	assert.Equal(t, messagesMap["same-name"], Get("same-name"))
}
