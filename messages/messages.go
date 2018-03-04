package messages

var messagesMap = map[string]string{
	"invalid-message-key":            "🛑 Invalid message key! 🛑",
	"direct-not-allowed":             "This look like a direct message. Containers are related to a channels.",
	"admin-only":                     "Command available only for admins. ⛔",
	"command-not-found":              "Not sure what you are asking for. Type `@claimr help` for valid commands.",
	"same-name":                      "There is a container with the same name on this channel. Try a different one.",
	"container-not-found":            "Unable to find a container with this name on this channel.",
	"added-to-channel":               "Container `%s` added to channel <#%s>.",
	"field-name-too-big":             "Try a smaller container name up to %d characters",
	"field-required":                 "Can not continue without a %s. 🙄",
	"container-not-found-on-channel": "I couldn't find the container `%s` on <#%s>.",
	"container-in-use":               "Container `%s` is already in use, try another one.",
	"fail-to-update":                 "Fail to update the container.",
	"container-claimed":              "Got it. Container `%s` is all yours <@%s>.",
	"container-in-use-by-other":      "Hum Container `%s` is not being used by you.",
	"container-in-use-by-this":       "Can't remove. Container `%s` is in used by <@%s> since _%s_.",
	"container-in-use-by-you":        "Container `%s` is already in use by you.",
	"container-free":                 "Got it. Container `%s` is now available.",
	"fail-getting-containers":        "Fail to list containers.",
	"empty-containers-list":          "No containers to list.",
	"containers-list":                "Here is a list of containers for this channel:",
	"no-level-provided":              "No log level provided, keeping in `%s`",
	"same-log-level":                 "Same log level than actual. Nothing change.",
	"invalid-log-level":              "not a valid logrus Level: \"%s\"",
	"level-log-changed":              "Log level changed from `%s` to `%s`",
	"x-purged":                       "%d Container rows purged.",
	"only-owner-can-remove":          "Only who created the container `%s` can remove it. Please check with <@%s>.",
	"container-removed":              "Container `%s` removed.",
	"container-created-by":           "Container `%s`.\nCreated by <@%s>.\n",
	"container-in-use-by-w-reason":   "In use by <@%s>%s since _%s_.",
	"shouldnt-mention-user":          "This message shouldn't contain user mentions.",
	"shouldnt-mention-channel":       "This message shouldn't contain channel mentions.",
	"x-admin-loaded":                 "%d admins loaded.",
	"in-use-text-invalid":            "Invalid Format for InUseText",
}

// Get returns a message related to the given key for a map of centralized strings to be used on the project
func Get(key string) string {
	msg := messagesMap[key]

	if msg == "" {
		return messagesMap["invalid-message-key"]
	}
	return msg
}
