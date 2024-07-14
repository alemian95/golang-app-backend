package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() {
	connection, dbErr := gorm.Open(postgres.Open("host=localhost user=db_user password=example dbname=golang_db port=5432"), &gorm.Config{})

	if dbErr != nil {
		panic(dbErr)
	} else {
		db = connection
	}
}

func Conn() *gorm.DB {
	return db
}
