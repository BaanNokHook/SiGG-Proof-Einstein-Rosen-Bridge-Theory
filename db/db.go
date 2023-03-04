package db

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/primasio/wormhole/config"
)

const (
	SQLITE = "sqlite3"
	MYSQL  = "mysql"
)

var instance *gorm.DB
var instanceType string

func GetDb() *gorm.DB {
	return instance
}

func GetDbType() string {
	return instanceType
}

func Init() error {

	c := config.GetConfig()

	instanceType = c.GetString("db.type")
	dbConn := c.GetString("db.connection")

	if dbConn == "" {
		f, err := ioutil.TempFile("", "")
		if err != nil {
			return err
		}
		dbConn := f.Name()
		f.Close()
		os.Remove(dbConn)
	}

	var err error

	instance, err = gorm.Open(instanceType, dbConn)

	if err != nil {
		return err
	}

	instance.Set("gorm:table_options", "charset=utf8mb4")

	return nil
}

func ForUpdate(tx *gorm.DB) *gorm.DB {
	if GetDbType() != SQLITE {
		return tx.Set("gorm:query_option", "FOR UPDATE")
	}
	return tx
}
