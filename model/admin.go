package model

import (
	"github.com/shomali11/slacker"
	log "github.com/sirupsen/logrus"
)

// Admins is a list of admins for this Team
var Admins []Admin

// Admin is the model representing a Team admin
type Admin struct {
	ID       string
	RealName string
}

// LoadAdmins will load the current slack team admins to be used on admin-only commands
func LoadAdmins(bot *slacker.Slacker) {
	log.Info("Loading admins...")
	users, err := bot.Client.GetUsers()

	if err != nil {
		log.Errorf("Error while loading admins from slack %s", err)
		return
	}

	Admins = []Admin{}
	for _, user := range users {
		log.Debugf("%s %s isAdmin [%t]", user.ID, user.RealName, (user.IsAdmin || user.IsOwner))
		if user.IsAdmin || user.IsOwner {
			Admins = append(Admins, Admin{ID: user.ID, RealName: user.RealName})
		}
	}
	log.Infof("%d admins loaded.", len(Admins))
}
