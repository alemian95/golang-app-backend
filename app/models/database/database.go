package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

/**
 * Connect to database
 */
func ConnectDatabase() {

	var database_host = os.Getenv("DB_HOST")
	var database_user = os.Getenv("DB_USER")
	var datanase_password = os.Getenv("DB_PASSWORD")
	var database_name = os.Getenv("DB_NAME")
	var database_port = os.Getenv("DB_PORT")

	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", database_host, database_user, datanase_password, database_name, database_port)

	connection, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbErr != nil {
		panic(dbErr)
	} else {
		db = connection
	}
}

func Conn() *gorm.DB {
	return db
}
