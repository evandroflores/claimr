package model

import "time"

// VM defines the VM on database.
type VM struct {
	ID            int64
	TeamID        string
	ChannelID     string
	Name          string
	InUseBy       string
	UpdatedAt     time.Time `xorm:"updated"`
	CreatedByUser string
}