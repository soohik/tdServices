package main

import (
	"tdapi/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// sqlFactory is receiver for Build method
var sqldb *gorm.DB

// implement Build method for SQL database
func SqlBuild() error {

	config := config.Config.SQLConfig

	db, err := gorm.Open(mysql.Open(config.UrlAddress), &gorm.Config{})
	if err != nil {
		return nil
	}
	// check the connection
	sqlDB, err := db.DB()
	sqlDB.Ping()

	sqldb = db
	return err

}
