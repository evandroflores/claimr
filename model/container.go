package model

import (
	"fmt"
	"time"

	"github.com/evandroflores/claimr/database"
)

// Container defines the Container information on database.
type Container struct {
	ID            int64
	TeamID        string
	ChannelID     string
	Name          string
	InUseBy       string
	InUseByReason string
	UpdatedAt     time.Time `xorm:"updated"`
	CreatedByUser string
}

func init() {
	database.RegisterModel(Container{})
}

// GetContainer returns a container for teamID, channelID, and name provided
func GetContainer(teamID string, channelID string, name string) (Container, error) {
	if teamID == "" {
		return Container{}, fmt.Errorf("Give me a teamID to find. 🙄")
	}

	if channelID == "" {
		return Container{}, fmt.Errorf("Give me a channelID to find. 🙄")
	}

	if name == "" {
		return Container{}, fmt.Errorf("Give me a container name to find. 🙄")
	}

	container := Container{TeamID: teamID, ChannelID: channelID, Name: name}

	found, err := database.DB.Get(&container)

	if err != nil {
		return Container{}, err
	}

	if !found {
		return Container{}, fmt.Errorf("Container %s not found", name)
	}

	return container, nil
}
