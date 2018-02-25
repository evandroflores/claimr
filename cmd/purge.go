package cmd

import (
	"fmt"

	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/messages"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
)

func init() {
	Register("purge", "Purge soft delete models from the database. admin-only", purge)
}

func purge(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()

	event := getEvent(request)
	if !isAdmin(event.User) {
		response.Reply(messages.Messages["admin-only"])
		return
	}

	result := database.DB.Unscoped().Where("deleted_at is not null").Delete(&model.Container{})
	response.Reply(fmt.Sprintf(messages.Messages["x-purged"], result.RowsAffected))
}
