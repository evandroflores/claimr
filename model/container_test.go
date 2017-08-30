package model

import (
	"testing"

	"fmt"
	"github.com/evandroflores/claimr/model"
	"github.com/stretchr/testify/assert"
	"github.com/evandroflores/claimr/database"
	"github.com/go-xorm/xorm"
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

type MockDB struct {

}

func (MockDB) Get(bean interface{}) (bool, error){
	return false, fmt.Errorf("attempt to write a readonly database")
}

func (MockDB) IsTableExist(beanOrTableName interface{}) (bool, error) {
	return false, nil
}
func (MockDB) CreateTables(beans ...interface{}) error {
	return nil
}

func (MockDB) ID(id interface{}) *xorm.Session{
	return nil
}
func (MockDB) Insert(beans ...interface{}) (int64, error){
	return int64(0), nil
}
func (MockDB) Find(beans interface{}, condiBeans ...interface{}) error{
	return nil
}

func TestGetContainerDBError(t *testing.T){
	_db := database.DB
	defer func() {
		database.DB = _db
	}()

	database.DB = new(MockDB)

	container, err := model.GetContainer("TeamID", "ChannelID", "ShouldReturnAnError")

	assert.ObjectsAreEqual(Container{}, container)
	assert.EqualError(t, err,"attempt to write a readonly database")

}

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
