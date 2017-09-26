package model

import (
	"fmt"
	"time"

	"github.com/evandroflores/claimr/database"
	"github.com/guregu/dynamo"
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

var maxNameSize = 22
var containerTable dynamo.Table

func init() {
	Container{}.initTable(database.DB)
}

func (container Container) initTable(db *dynamo.DB) {
	containerTable = db.Table("ClaimrContainer")
}

func isValidContainerInput(teamID string, channelID string, name string) (bool, error) {
	if teamID == "" {
		return false, fmt.Errorf("can not continue without a teamID 🙄")
	}

	if channelID == "" {
		return false, fmt.Errorf("can not continue without a channelID 🙄")
	}

	if name == "" {
		return false, fmt.Errorf("can not continue without a container name 🙄")
	}

	if len(name) > maxNameSize {
		return false, fmt.Errorf("try a name up to %d characters", maxNameSize)
	}

	return true, nil
}

// GetContainer returns a container for teamID, channelID, and name provided
func GetContainer(teamID string, channelID string, name string) (Container, error) {
	valid, err := isValidContainerInput(teamID, channelID, name)

	if !valid {
		return Container{}, err
	}

	results := []Container{}

	err = containerTable.Scan().
		Filter("TeamID = ? AND ChannelID = ? AND 'Name' = ? ", teamID, channelID, name).
		All(&results)

	if err != nil {
		return Container{}, err
	}

	if len(results) == 0 {
		return Container{}, nil
	}

	return results[0], nil
}

// GetContainers returns a list of containers for the given TeamID and ChannelID
func GetContainers(teamID string, channelID string) ([]Container, error) {
	results := []Container{}

	err := containerTable.Scan().
		Filter("TeamID = ? AND ChannelID = ? ", teamID, channelID).
		All(&results)

	if err != nil {
		return []Container{}, err
	}

	if len(results) == 0 {
		return []Container{}, nil
	}

	return results, nil
}

func (container Container) getID() string {
	return fmt.Sprintf("%s.%s.%s", container.TeamID, container.ChannelID, container.Name)
}

func (container Container) Add() error {
	valid, err := isValidContainerInput(container.TeamID, container.ChannelID, container.Name)

	if !valid {
		return err
	}

	existingContainer, err := GetContainer(container.TeamID, container.ChannelID, container.Name)
	if err != nil {
		return err
	}
	if existingContainer != (Container{}) {
		return fmt.Errorf("there is a container with the same name on this channel. Try a different one 😕")
	}

	container.ID = container.getID()
	container.UpdatedAt = time.Now().UTC()

	err = containerTable.Put(container).Run()
	if err != nil {
		return err
	}

	return nil
}

// Update a given Container
func (container Container) Update() error {
	valid, err := isValidContainerInput(container.TeamID, container.ChannelID, container.Name)

	if !valid {
		return err
	}

	existingContainer, err := GetContainer(container.TeamID, container.ChannelID, container.Name)
	if err != nil {
		return err
	}
	if existingContainer == (Container{}) {
		return fmt.Errorf("could not find this container on this channel. Can not update 😕")
	}

	container.ID = container.getID()
	container.UpdatedAt = time.Now().UTC()

	err = containerTable.Put(container).Run()
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a Container from the database
func (container Container) Delete() error {
	valid, err := isValidContainerInput(container.TeamID, container.ChannelID, container.Name)

	if !valid {
		return err
	}

	existingContainer, err := GetContainer(container.TeamID, container.ChannelID, container.Name)
	if err != nil {
		return err
	}
	if existingContainer == (Container{}) {
		return fmt.Errorf("could not find this container on this channel. Can not delete 😕")
	}

	err = containerTable.Delete("ID", container.getID()).Run()
	if err != nil {
		return err
	}

	return nil
}
