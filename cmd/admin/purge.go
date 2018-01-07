package admin

import (
	"fmt"

	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/model"
	"github.com/shomali11/slacker"
)

func init() {
	Register("purge", "Purge soft delete models from the database.", purge)
}

func purge(request *slacker.Request, response slacker.ResponseWriter) {
	result := database.DB.Unscoped().Where("deleted_at is not null").Delete(&model.Container{})
	response.Reply(fmt.Sprintf("%d Container rows purged", result.RowsAffected))
}
