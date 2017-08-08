package model

import "time"

// Container defines the Container information on database.
type Container struct {
	ID            int64
	TeamID        string
	ChannelID     string
	Name          string
	InUseBy       string
	UpdatedAt     time.Time `xorm:"updated"`
	CreatedByUser string
}
