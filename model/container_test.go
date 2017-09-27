package model

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestGetContainerNoTeam(t *testing.T) {
	container, err := GetContainer("", "", "")

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err, "can not continue without a teamID 🙄")
}

func TestGetContainerNoChannel(t *testing.T) {
	container, err := GetContainer("TeamID", "", "")

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err, "can not continue without a channelID 🙄")
}

func TestGetContainerNoName(t *testing.T) {
	container, err := GetContainer("TeamID", "ChannelID", "")

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err, "can not continue without a container name 🙄")
}

func TestGetContainerBigName(t *testing.T) {
	container, err := GetContainer("TeamID", "ChannelID",
		"LoremIpsumDolorSitAmetConsecteturAdipiscingElit")

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err, "try a name up to 22 characters")
}

func TestGetContainerNotFound(t *testing.T) {
	containerName := "TestDoesNotExist"
	container, err := GetContainer("TeamID", "ChannelID", containerName)

	assert.ObjectsAreEqual(Container{}, container)
	assert.NoError(t, err, fmt.Sprintf("Container %s not found", containerName))
}

func TestAddContainer(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: "Name"}
	err := container.Add()
	assert.NoError(t, err)

	containerFromDB, err2 := GetContainer("TestTeam", "TestChannel", "Name")

	assert.NoError(t, err2)
	assert.ObjectsAreEqual(container, containerFromDB)
}

func TestAddContainerValidateTeamID(t *testing.T) {
	container := Container{TeamID: "", ChannelID: "", Name: ""}
	err := container.Add()
	assert.EqualError(t, err, "can not continue without a teamID 🙄")
}

func TestAddContainerValidateChannelID(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "", Name: ""}
	err := container.Add()
	assert.EqualError(t, err, "can not continue without a channelID 🙄")
}

func TestAddContainerValidateName(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: ""}
	err := container.Add()
	assert.EqualError(t, err, "can not continue without a container name 🙄")
}

func TestAddContainerDuplicate(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: "Name"}
	err := container.Add()
	assert.EqualError(t, err, "there is a container with the same name on this channel. Try a different one 😕")
}

func TestDeleteContainer(t *testing.T) {
	container, err := GetContainer("TestTeam", "TestChannel", "Name")
	assert.NoError(t, err)

	err2 := container.Delete()
	assert.NoError(t, err2)
}

func TestDeleteContainerValidateTeamID(t *testing.T) {
	container := Container{TeamID: "", ChannelID: "", Name: ""}
	err := container.Delete()
	assert.EqualError(t, err, "can not continue without a teamID 🙄")
}

func TestDeleteContainerValidateChannelID(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "", Name: ""}
	err := container.Delete()
	assert.EqualError(t, err, "can not continue without a channelID 🙄")
}

func TestDeleteContainerValidateName(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: ""}
	err := container.Delete()
	assert.EqualError(t, err, "can not continue without a container name 🙄")
}

func TestDeleteContainerInexistent(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: "Name"}
	err := container.Delete()
	assert.EqualError(t, err, "could not find this container on this channel. Can not delete 😕")
}

func TestUpdateContainer(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: "Name"}
	err := container.Add()
	assert.NoError(t, err)

	container.InUseBy = "me"
	err2 := container.Update()
	assert.NoError(t, err2)

	containerUpdated := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: "Name"}

	assert.ObjectsAreEqual(containerUpdated.InUseBy, "me")
	assert.ObjectsAreEqual(container, containerUpdated)
	container.Delete()
}

func TestUpdateContainerValidateTeamID(t *testing.T) {
	container := Container{TeamID: "", ChannelID: "", Name: ""}
	err := container.Update()
	assert.EqualError(t, err, "can not continue without a teamID 🙄")
}

func TestUpdateContainerValidateChannelID(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "", Name: ""}
	err := container.Update()
	assert.EqualError(t, err, "can not continue without a channelID 🙄")
}

func TestUpdateContainerValidateName(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: ""}
	err := container.Update()
	assert.EqualError(t, err, "can not continue without a container name 🙄")
}

func TestUpdateContainerInexistent(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: "Name"}
	err := container.Update()
	assert.EqualError(t, err, "could not find this container on this channel. Can not update 😕")
}

func TestListContainers(t *testing.T) {
	names := [4]string{"A", "B", "C", "D"}
	for _, name := range names {
		container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: name}
		err := container.Add()
		assert.NoError(t, err)
	}

	containers, err2 := GetContainers("TestTeam", "TestChannel")
	assert.NoError(t, err2)

	assert.Len(t,containers, 4)
	for idx, container := range containers {
		assert.ObjectsAreEqual(container.Name, names[idx])
		container.Delete()
	}
}

func TestListContainersValidateTeamID(t *testing.T) {
	containers, err := GetContainers("", "")
	assert.EqualError(t, err, "can not continue without a teamID 🙄")
	assert.ObjectsAreEqual(containers, []Container{})
}

func TestListContainersValidateChannelID(t *testing.T) {
	containers, err := GetContainers("TestTeam", "")
	assert.EqualError(t, err, "can not continue without a channelID 🙄")
	assert.ObjectsAreEqual(containers, []Container{})
}