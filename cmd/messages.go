package cmd

import (
	"fmt"

	"github.com/evandroflores/claimr/model"
)

// Messages is a map of centralized strings to be used on the project
var Messages = map[string]string{
	"not-implemented":                "No pancakes for you! 🥞",
	"direct-not-allowed":             "this look like a direct message. Containers are related to a channels",
	"admin-only":                     "Command available only for admins. ⛔",
	"command-not-found":              "Not sure what you are asking for. Type `@claimr help` for valid commands.",
	"same-name":                      "There is a container with the same name on this channel. Try a different one.",
	"added-to-channel":               "Container `%s` added to channel <#%s>.",
	"field-name-too-big":             fmt.Sprintf("try a smaller container name up to %d characters", model.MaxNameSize),
	"field-name-required":            "can not continue without a container name 🙄",
	"container-not-found-on-channel": "I couldn't find the container `%s` on <#%s>.",
	"container-in-use":               "Container `%s` is already in use, try another one.",
	"fail-to-update":                 "Fail to update the container.",
	"container-claimed":              "Got it. Container `%s` is all yours <@%s>.",
	"container-in-use-by-other":      "Hum Container `%s` is not being used by you.",
	"container-free":                 "Got it. Container `%s` is now available",
	"fail-getting-containers":        "Fail to list containers.",
	"empty-containers-list":          "No containers to list.",
	"containers-list":                "Here is a list of containers for this channel:",
	"no-level-provided":              "No log level provided, keeping in `%s`",
	"same-log-level":                 "Same log level than actual. Nothing change.",
	"invalid-log-level":              "not a valid logrus Level: \"%s\"",
	"level-log-changed":              "Log level changed from `%s` to `%s`",
}
