package model

import (
	"testing"

	"fmt"
	"github.com/evandroflores/claimr/model"
	"github.com/stretchr/testify/assert"
	"github.com/evandroflores/claimr/database"
)


func TestGetContainerNoTeam(t *testing.T) {
	container, err := model.GetContainer("", "", "")

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err, "Give me a teamID to find. 🙄")
}

func TestGetContainerNoChannel(t *testing.T) {
	container, err := model.GetContainer("TeamID", "", "")

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err, "Give me a channelID to find. 🙄")
}

func TestGetContainerNoName(t *testing.T) {
	container, err := model.GetContainer("TeamID", "ChannelID", "")

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err, "Give me a container name to find. 🙄")
}

func TestGetContainerNotFound(t *testing.T) {
	containerName := "TestDoesNotExist"
	container, err := model.GetContainer("TeamID", "ChannelID", containerName)

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err, fmt.Sprintf("Container %s not found", containerName))
}

//TestGetContainerDBError? How to raise an error?

func TestGetContainerFound(t *testing.T) {
	teamID := "TeamID"
	channelID := "ChannelID"
	containerName := "Test"
	testContainer := Container{TeamID: teamID, ChannelID: channelID, Name:containerName, InUseBy: "free"}
	database.DB.Insert(testContainer)

	container, err := model.GetContainer(teamID, channelID, containerName)

	assert.ObjectsAreEqual(testContainer, container)
	assert.NoError(t, err)
}
