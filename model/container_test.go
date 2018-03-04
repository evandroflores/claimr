package model

import (
	"testing"

	"fmt"

	"strings"

	"github.com/evandroflores/claimr/messages"
	"github.com/stretchr/testify/assert"
)

func TestGetContainerNoTeam(t *testing.T) {
	container, err := GetContainer("", "", "")

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "teamID"))
}

func TestGetContainerNoChannel(t *testing.T) {
	container, err := GetContainer("TestTeam", "", "")

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "channelID"))
}

func TestGetContainerNoName(t *testing.T) {
	container, err := GetContainer("TestTeam", "TestChannel", "")

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "container name"))
}

func TestGetContainerBigName(t *testing.T) {
	container, err := GetContainer("TestTeam", "TestChannel",
		"LoremIpsumDolorSitAmetConsecteturAdipiscingElit")

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-name-too-big"), MaxNameSize))
}

func TestGetContainerNotFound(t *testing.T) {
	containerName := "TestDoesNotExist"
	container, err := GetContainer("TestTeam", "TestChannel", containerName)

	assert.ObjectsAreEqual(Container{}, container)
	assert.NoError(t, err, fmt.Sprintf("Container %s not found", containerName))
}

func TestGetContainerIgnoreCase(t *testing.T) {
	containerName := "UPPERCASE_NAME"
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: containerName}
	err := container.Add()
	assert.NoError(t, err)

	containerFromDB, err2 := GetContainer("TestTeam", "TestChannel", containerName)

	assert.NoError(t, err2)
	assert.Equal(t, containerFromDB.Name, strings.ToLower(containerName))

	container.Delete()
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
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "teamID"))
}

func TestAddContainerValidateChannelID(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "", Name: ""}
	err := container.Add()
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "channelID"))
}

func TestAddContainerValidateName(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: ""}
	err := container.Add()
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "container name"))
}

func TestAddContainerDuplicate(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: "Name"}
	err := container.Add()
	assert.EqualError(t, err, messages.Get("same-name"))
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
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "teamID"))
}

func TestDeleteContainerValidateChannelID(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "", Name: ""}
	err := container.Delete()
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "channelID"))
}

func TestDeleteContainerValidateName(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: ""}
	err := container.Delete()
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "container name"))
}

func TestDeleteContainerNotFound(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: "Name"}
	err := container.Delete()
	assert.EqualError(t, err, messages.Get("container-not-found"))

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
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "teamID"))
}

func TestUpdateContainerValidateChannelID(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "", Name: ""}
	err := container.Update()
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "channelID"))
}

func TestUpdateContainerValidateName(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: ""}
	err := container.Update()
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "container name"))
}

func TestUpdateContainerNotFound(t *testing.T) {
	container := Container{TeamID: "TestTeam", ChannelID: "TestChannel", Name: "Name"}
	err := container.Update()
	assert.EqualError(t, err, messages.Get("container-not-found"))
}

func TestListContainers(t *testing.T) {
	names := [4]string{"A", "B", "C", "D"}
	for _, name := range names {
		container := Container{TeamID: "TestList", ChannelID: "TestChannel", Name: name}
		err := container.Add()
		assert.NoError(t, err)
	}

	containers, err2 := GetContainers("TestList", "TestChannel")
	assert.NoError(t, err2)

	assert.Len(t, containers, len(names))
	for idx, container := range containers {
		assert.ObjectsAreEqual(container.Name, names[idx])
		container.Delete()
	}
}

func TestListContainersValidateTeamID(t *testing.T) {
	containers, err := GetContainers("", "")
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "teamID"))
	assert.ObjectsAreEqual(containers, []Container{})
}

func TestListContainersValidateChannelID(t *testing.T) {
	containers, err := GetContainers("TestTeam", "")
	assert.EqualError(t, err, fmt.Sprintf(messages.Get("field-required"), "channelID"))
	assert.ObjectsAreEqual(containers, []Container{})
}

func TestRemoveInUseData(t *testing.T) {
	team := "TestTeam"
	channel := "TestChannel"
	containerName := "name"
	user := "User"
	reason := "testing"

	container := Container{TeamID: team, ChannelID: channel, Name: containerName, InUseBy: user, InUseForReason: reason}
	err := container.Add()
	assert.NoError(t, err)

	err2 := container.ClearInUse()
	assert.NoError(t, err2)

	containerExpected := Container{TeamID: team, ChannelID: channel, Name: containerName, InUseBy: "", InUseForReason: ""}
	containerFromDB, err3 := GetContainer(team, channel, containerName)
	assert.NoError(t, err3)

	// Ignoring the difference for this fields
	containerExpected.ID = containerFromDB.ID
	containerExpected.CreatedAt = containerFromDB.CreatedAt
	containerExpected.UpdatedAt = containerFromDB.UpdatedAt
	containerExpected.DeletedAt = containerFromDB.DeletedAt

	assert.Empty(t, containerFromDB.InUseBy)
	assert.Empty(t, containerFromDB.InUseForReason)
	assert.Equal(t, containerExpected, containerFromDB)
	container.Delete()
}

func TestSetInUseData(t *testing.T) {
	team := "TestTeam"
	channel := "TestChannel"
	containerName := "name"
	user := "User"
	reason := "testing"

	container := Container{TeamID: team, ChannelID: channel, Name: containerName, InUseBy: "", InUseForReason: ""}
	err := container.Add()
	assert.NoError(t, err)

	err2 := container.SetInUse(user, reason)
	assert.NoError(t, err2)

	containerExpected := Container{TeamID: team, ChannelID: channel, Name: containerName, InUseBy: user, InUseForReason: reason}
	containerFromDB, err3 := GetContainer(team, channel, containerName)
	assert.NoError(t, err3)

	// Ignoring the difference for this fields
	containerExpected.ID = containerFromDB.ID
	containerExpected.CreatedAt = containerFromDB.CreatedAt
	containerExpected.UpdatedAt = containerFromDB.UpdatedAt
	containerExpected.DeletedAt = containerFromDB.DeletedAt

	assert.Equal(t, containerExpected, containerFromDB)
	container.Delete()
}

func TestInUseTextInvalid(t *testing.T) {
	container := Container{}
	assert.Equal(t, messages.Get("in-use-text-invalid"), container.InUseText("whatever"))
}

func TestInUseTextSimpleAvailable(t *testing.T) {
	container := Container{InUseBy: ""}
	assert.Equal(t, available, container.InUseText("simple"))
}
