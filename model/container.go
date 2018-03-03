package model

import (
	"fmt"

	"strings"

	"github.com/evandroflores/claimr/database"
	"github.com/evandroflores/claimr/messages"
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
const MaxNameSize = 22

func isValidContainerInput(teamID string, channelID string, containerName string) (bool, error) {
	fields := []struct {
		name  string
		value string
	}{
		{"teamID", teamID},
		{"channelID", channelID},
		{"container name", containerName},
	}

	for _, field := range fields {
		err := checkRequired(field.name, field.value)
		if err != nil {
			return false, err
		}
	}

	if len(containerName) > MaxNameSize {
		return false, fmt.Errorf(messages.Messages["field-name-too-big"], MaxNameSize)
	}

	return true, nil
}

func checkRequired(fieldName string, fieldValue string) error {
	if fieldValue == "" {
		return fmt.Errorf(messages.Messages["field-required"], fieldName)
	}
	return nil
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
		return fmt.Errorf(messages.Messages["same-name"])
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
		return fmt.Errorf(messages.Messages["container-not-found"])
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
		return fmt.Errorf(messages.Messages["container-not-found"])
	}

	database.DB.Delete(&existingContainer)

	return nil
}

// ClearInUse removes information InUseBy and InUseForReason, for a given container
func (container Container) ClearInUse() error {
	container.InUseBy = ""
	container.InUseForReason = ""

	return container.Update()
}

// SetInUse sets information InUseBy and InUseForReason, for a given container
func (container Container) SetInUse(by string, reason string) error {
	container.InUseBy = by
	container.InUseForReason = reason

	return container.Update()
}
