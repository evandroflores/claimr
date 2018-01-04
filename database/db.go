package database

import (
	"os"

	_ "github.com/go-sql-driver/mysql" // MySql driver for database
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// DB is the single database instance
var DB *gorm.DB

func init() {
	initDB()
}

func initDB() {
	log.Info("Initializing database")
	var err error

	dbStringConnection := os.Getenv("CLAIMR_DATABASE")
	if dbStringConnection == "" {
		log.Fatal("Claimr mysql database string unset. Set CLAIMR_DATABASE to continue.")
		return
	}

	DB, err = gorm.Open("mysql", dbStringConnection)

	if err != nil {
		log.Fatalf("could not create a database connection - %s", err)
		return
	}
	DB.LogMode(true)
}
