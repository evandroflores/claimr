package database

import (
	_ "github.com/go-sql-driver/mysql" // MySql driver for database
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// DB is the single database instance
var DB *gorm.DB

func init() {
	log.Info("Initializing database")
	var err error

	DB, err = gorm.Open("mysql", "root@/claimr?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("could not create a database connection - %s", err)
	}
}
