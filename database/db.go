package database

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
 _ "github.com/go-sql-driver/mysql"
)

// DB is the single database instance
var DB *gorm.DB

func init() {
	log.Info("Initializing database")
	var err error = nil

	DB, err = gorm.Open("mysql", "root@/claimr?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("could not create a database connection - %s", err)
	}
}
