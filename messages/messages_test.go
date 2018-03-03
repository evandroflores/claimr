package messages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInvalidKey(t *testing.T) {
	assert.Equal(t, Messages["invalid-message-key"], Get("?"))
}

func TestGetValidKey(t *testing.T) {
	assert.Equal(t, Messages["same-name"], Get("same-name"))
}
