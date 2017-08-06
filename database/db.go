package database

import (
	"github.com/evandroflores/claimr/model"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3" // Engine for database
	log "github.com/sirupsen/logrus"
)

// DB is the orm interface to the database
var DB, _ = xorm.NewEngine("sqlite3", "./claimr.db")

// InitDB Initialises the database tables and set debug variables
func InitDB() {
	log.Info("Initializing database...")

	DB.ShowSQL(true)
	DB.ShowExecTime(true)

	tableExists, _ := DB.IsTableExist(&model.VM{})
	if !tableExists {
		log.Info("Creating tables...")
		DB.CreateTables(&model.VM{})
	}
	log.Info("Done.")
}
