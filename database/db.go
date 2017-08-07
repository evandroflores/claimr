package database

import (
	"github.com/evandroflores/claimr/model"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3" // Engine for database
	log "github.com/sirupsen/logrus"
)

// DB is the orm interface to the database
var DB *xorm.Engine

func init() {
	initDB("./claimr.db")
}

func initDB(dbName string) {
	log.Infof("Initializing database [%s]...", dbName)

	var err error
	DB, err = xorm.NewEngine("sqlite3", dbName)

	if err != nil {
		panic("Couldn't open nor create database")
	}

	DB.ShowSQL(true)
	DB.ShowExecTime(true)

	tableExists, err := DB.IsTableExist(&model.VM{})

	if err != nil {
		panic("Couldn't read database")
	}

	if !tableExists {
		log.Info("Creating tables...")
		DB.CreateTables(&model.VM{})
	}
	log.Info("Done.")
}
