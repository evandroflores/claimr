package database

import (
	"os"
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

	wantMsg := "Claimr mysql database string unset. Set CLAIMR_DATABASE to continue."

	mockLogFatal := func(msg ...interface{}) {
		assert.Equal(t, wantMsg, msg[0])
		panic("log.Fatal called")
	}
	patchLog := monkey.Patch(log.Fatal, mockLogFatal)
	defer patchLog.Unpatch()
	assert.PanicsWithValue(t, "log.Fatal called", initDB, "log.Fatal was not called")
}

func TestDBErrpr(t *testing.T) {
	wantMsg := "could not create a database connection - [ERROR]"

	mockLogFatal := func(format string, args ...interface{}) {
		assert.Equal(t, wantMsg, fmt.Sprintf(format, args))
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
