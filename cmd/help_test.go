package cmd

import (
	"reflect"
	"testing"

	"github.com/bouk/monkey"
	"github.com/shomali11/slacker"
	"github.com/stretchr/testify/assert"
)

func TestCallHelp(t *testing.T) {
	var mockResponse *slacker.Response

	expected := GenerateCommandHelp()

	patchReply := monkey.PatchInstanceMethod(reflect.TypeOf(mockResponse), "Reply",
		func(response *slacker.Response, msg string) {
			assert.Equal(t, expected, msg)
		})

	_ = monkey.PatchInstanceMethod(reflect.TypeOf(mockResponse), "Typing",
		func(response *slacker.Response) {})

	Help(nil, mockResponse)

	patchReply.Unpatch()
}
