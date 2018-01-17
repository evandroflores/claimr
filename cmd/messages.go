package cmd

import (
	"fmt"

	"github.com/evandroflores/claimr/model"
)

// Messages is a map of centralized strings to be used on the project
var Messages = map[string]string{
	"same-name":                      "There is a container with the same name on this channel. Try a different one.",
	"added-to-channel":               "Container `%s` added to channel <#%s>.",
	"name-too-big":                   fmt.Sprintf("try a smaller container name up to %d characters", model.MaxNameSize),
	"container-not-found-on-channel": "I couldn't find the container `%s` on <#%s>.",
	"container-in-use":               "Container `%s` is already in use, try another one.",
	"fail-to-update":                 "Fail to update the container.",
	"container-claimed":              "Got it. Container `%s` is all yours <@%s>.",
}
