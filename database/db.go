package database

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3" // Engine for database
	"log"
	"github.com/evandroflores/claimr/model"
)

// DB is the orm interface to the database
var DB, _  = xorm.NewEngine("sqlite3", "./claimr.db")

// InitDB Initialises the database tables and set debug variables
func InitDB() {
	log.Println("Initializing database...")
	DB.ShowSQL(true)
	DB.ShowExecTime(true)
		
    tableExists, _ := DB.IsTableExist(&model.VM{})
    if !tableExists {
		log.Println("Creating tables...")
        DB.CreateTables(&model.VM{})
	}
	log.Println("Done.")
}