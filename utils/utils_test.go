package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfThenElseBooleanTrue(t *testing.T) {
	theTruth := "I am the way and the truth and the life"
	notTrue := "Fake news"

	returned := IfThenElse(true, theTruth, notTrue)

	assert.Equal(t, theTruth, returned)
}

func TestIfThenElseBooleanFalse(t *testing.T) {
	theTruth := "I am the way and the truth and the life"
	notTrue := "Fake news"

	returned := IfThenElse(false, theTruth, notTrue)

	assert.Equal(t, notTrue, returned)
}
