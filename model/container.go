package model

import (
	"fmt"

	"strings"

	"github.com/evandroflores/claimr/database"
	"github.com/jinzhu/gorm"
)

func init() {
	database.DB.AutoMigrate(&Container{})
}

// Container defines the Container information on database.
type Container struct {
	gorm.Model
	TeamID         string `gorm:"not null"`
	ChannelID      string `gorm:"not null"`
	Name           string `gorm:"not null"`
	InUseBy        string
	InUseForReason string
	CreatedByUser  string
}

// MaxNameSize is the max number of characters for a container name.
var MaxNameSize = 22

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

	if len(name) > MaxNameSize {
		return false, fmt.Errorf("try a name up to %d characters", MaxNameSize)
	}

	return true, nil
}

// GetContainer returns a container for teamID, channelID, and name provided
func GetContainer(teamID string, channelID string, name string) (Container, error) {
	result := Container{}
	valid, err := isValidContainerInput(teamID, channelID, name)

	if !valid {
		return result, err
	}

	database.DB.Where(&Container{TeamID: teamID, ChannelID: channelID, Name: strings.ToLower(name)}).
		First(&result)

	return result, nil
}

// GetContainers returns a list of containers for the given TeamID and ChannelID
func GetContainers(teamID string, channelID string) ([]Container, error) {
	results := []Container{}
	valid, err := isValidContainerInput(teamID, channelID, ".")

	if !valid {
		return results, err
	}

	database.DB.Where(&Container{TeamID: teamID, ChannelID: channelID}).
		Find(&results)

	return results, nil
}

// Add a given Container to database
func (container Container) Add() error {
	existingContainer, err := GetContainer(container.TeamID, container.ChannelID, container.Name)

	if err != nil {
		return err
	}

	if existingContainer != (Container{}) {
		return fmt.Errorf("there is a container with the same name on this channel. Try a different one 😕")
	}
	container.Name = strings.ToLower(container.Name)
	database.DB.Create(&container)

	return nil
}

// Update a given Container
func (container Container) Update() error {
	existingContainer, err := GetContainer(container.TeamID, container.ChannelID, strings.ToLower(container.Name))

	if err != nil {
		return err
	}

	if existingContainer == (Container{}) {
		return fmt.Errorf("could not find this container on this channel. Can not update 😕")
	}

	existingContainer.InUseBy = container.InUseBy
	existingContainer.InUseForReason = container.InUseForReason

	database.DB.Save(&existingContainer)

	return nil
}

// Delete removes a Container from the database
func (container Container) Delete() error {
	existingContainer, err := GetContainer(container.TeamID, container.ChannelID, strings.ToLower(container.Name))

	if err != nil {
		return err
	}
	if existingContainer == (Container{}) {
		return fmt.Errorf("could not find this container on this channel. Can not delete 😕")
	}

	database.DB.Delete(&existingContainer)

	return nil
}
