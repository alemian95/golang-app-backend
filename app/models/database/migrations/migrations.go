package migrations

import (
	"golang-app/app/models/database"
	"golang-app/app/models/user_model"
	"golang-app/app/utils/auth"
)

/**
 * Migrates the database.
 */
func Migrate() {
	database.Conn().AutoMigrate(&user_model.User{})
}

func Seed() {

	hash, _ := auth.HashPassword("example")

	database.ConnectDatabase()

	user := user_model.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: hash,
	}

	user.Create()
}
