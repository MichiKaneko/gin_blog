package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var db_err error

func DBConnect(connectionString string) {
	Database, db_err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if db_err != nil {
		log.Fatal(db_err)
	}
	log.Println("Database connected")
}
