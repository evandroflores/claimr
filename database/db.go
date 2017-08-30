package database

import (
	"reflect"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3" // Engine for database
	log "github.com/sirupsen/logrus"
)

// DB is the orm interface to the database
var DB DbWilson

type DbWilson interface {
	Get(bean interface{}) (bool, error)
	IsTableExist(beanOrTableName interface{}) (bool, error)
	CreateTables(beans ...interface{}) error
	ID(id interface{}) *xorm.Session
	Insert(beans ...interface{}) (int64, error)
	Find(beans interface{}, condiBeans ...interface{}) error
}

func init() {
	log.SetLevel(log.DebugLevel)
	initDB("./claimr.db")
}

func initDB(dbName string) {
	log.Infof("Initializing database [%s]...", dbName)

	xorDB, err := xorm.NewEngine("sqlite3", dbName)

	if err != nil {
		log.Fatalf("Couldn't open nor create database [%s].", dbName)
	}
	DB = xorDB

	xorDB.ShowSQL(true)
	xorDB.ShowExecTime(true)
}

// RegisterModel will check and create the model if does not exists on the databse.
func RegisterModel(model interface{}) {

	modelName := reflect.TypeOf(model).Name()
	log.Debugf("Checking table for model %s", modelName)
	tableExists, err := DB.IsTableExist(model)

	if err != nil {
		log.Fatalf("Fail to check model %s. %s", modelName, err)
	}

	if !tableExists {
		log.Infof("Model %s does not exists on the database. Creating table...", modelName)
		DB.CreateTables(model)
	}

}
