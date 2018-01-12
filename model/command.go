package model

import "github.com/evandroflores/slacker"

// Command defines a command to be register to slack
type Command struct {
	Usage       string
	Description string
	Handler     func(request *slacker.Request, response slacker.ResponseWriter)
}
