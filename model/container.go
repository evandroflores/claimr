package model

import (
	"fmt"
	"time"

	"github.com/evandroflores/claimr/database"
	"github.com/guregu/dynamo"
	log "github.com/sirupsen/logrus"
)

// Container defines the Container information on database.
type Container struct {
	ID            string
	TeamID        string
	ChannelID     string
	Name          string
	InUseBy       string
	InUseByReason string
	UpdatedAt     time.Time
	CreatedByUser string
}

var containerTable dynamo.Table

func init() {
	database.DB2.CreateTable("ClaimrContainer", Container{})
	containerTable = database.DB2.Table("ClaimrContainer")
	log.Debug("___")
	name := fmt.Sprintf("name%s", time.Now().UTC())
	log.Debug(AddContainer("team", "channel", name, "owner"))
	log.Debug("___")
}

func isValidContainerInput(teamID string, channelID string, name string) (bool, error){
	if teamID == "" {
		return false, fmt.Errorf("Cant' continue without a teamID. 🙄")
	}

	if channelID == "" {
		return false, fmt.Errorf("Cant' continue without a channelID. 🙄")
	}

	if name == "" {
		return false, fmt.Errorf("Cant' continue without a container name 🙄")
	}

	return true, nil
}

// GetContainer returns a container for teamID, channelID, and name provided
func GetContainer(teamID string, channelID string, name string) (Container, error) {
	valid, err := isValidContainerInput(teamID, channelID, name)

	if !valid{
		return Container{}, err
	}

	results := []Container{}

	err = containerTable.Scan().
			Filter("TeamID = ? AND ChannelID = ? AND 'Name' = ? ", teamID , channelID, name).
			All(&results)

	if err != nil {
		return Container{}, err
	}

	if len(results) == 0 {
		return Container{}, nil
	}

	return results[0], nil
}

func AddContainer(teamID string, channelID string, name string, userID string) error {
	valid, err := isValidContainerInput(teamID, channelID, name)

	if !valid{
		return err
	}

	container, err := GetContainer(teamID, channelID, name)
	if err != nil {
		return err
	}
	if container != (Container{}) {
		return fmt.Errorf("There is a container with the same name on this channel. Try a different one. 😕")
	}

	now := time.Now().UTC()

	newContainer := Container{
						ID: fmt.Sprintf("%s.%s.%s", teamID, channelID, name),
						TeamID: teamID,
						ChannelID:channelID,
						Name:name,
						InUseBy: "",
						InUseByReason: "",
						UpdatedAt: now,
						CreatedByUser: userID,
					}

	err = containerTable.Put(newContainer).Run()
	if err != nil {
		return err
	}

	return nil
}
