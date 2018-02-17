package database

import (
	"os"
	"reflect"
	"testing"

	"errors"

	"fmt"

	"github.com/bouk/monkey"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestEnvironmentDatabase(t *testing.T) {
	currentEnv := os.Getenv("CLAIMR_DATABASE")
	os.Unsetenv("CLAIMR_DATABASE")
	defer func() { os.Setenv("CLAIMR_DATABASE", currentEnv) }()

	expectedMsg := "Claimr mysql database string unset. Set CLAIMR_DATABASE to continue."

	mockLogFatal := func(msg ...interface{}) {
		assert.Equal(t, expectedMsg, msg[0])
		panic("log.Fatal called")
	}
	patchLog := monkey.Patch(log.Fatal, mockLogFatal)
	defer patchLog.Unpatch()
	assert.PanicsWithValue(t, "log.Fatal called", initDB, "log.Fatal was not called")
}

func TestDBError(t *testing.T) {
	expectedMsg := "could not create a database connection - [ERROR]"

	mockLogFatal := func(format string, args ...interface{}) {
		assert.Equal(t, expectedMsg, fmt.Sprintf(format, args))
		panic("log.Fatal called")
	}

	mockGormOpen := func(dialect string, args ...interface{}) (db *gorm.DB, err error) {
		return nil, errors.New("ERROR")
	}

	patchLog := monkey.Patch(log.Fatalf, mockLogFatal)
	patchGorm := monkey.Patch(gorm.Open, mockGormOpen)

	defer patchLog.Unpatch()
	defer patchGorm.Unpatch()

	assert.PanicsWithValue(t, "log.Fatal called", initDB, "log.Fatal was not called")
}

func TestDBClose(t *testing.T) {
	mockLogInfo := func(args ...interface{}) {
		assert.Equal(t, "DB Closed", args[0])
	}

	mockDBClose := func(db *gorm.DB) error {
		return nil
	}

	patchLog := monkey.Patch(log.Info, mockLogInfo)
	patchClose := monkey.PatchInstanceMethod(reflect.TypeOf(DB), "Close", mockDBClose)

	CloseDB()

	patchLog.Unpatch()
	patchClose.Unpatch()
}

func TestDBCloseError(t *testing.T) {
	expectedMsg := "Error while closing database. Please check for memory leak [Simulated error]"

	mockLogWarnf := func(format string, args ...interface{}) {
		assert.Equal(t, expectedMsg, fmt.Sprintf(format, args))
	}

	mockDBClose := func(db *gorm.DB) error {
		return fmt.Errorf("Simulated error")
	}

	patchLog := monkey.Patch(log.Warnf, mockLogWarnf)
	patchClose := monkey.PatchInstanceMethod(reflect.TypeOf(DB), "Close", mockDBClose)

	CloseDB()

	patchLog.Unpatch()
	patchClose.Unpatch()
}
