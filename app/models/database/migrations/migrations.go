package migrations

import (
	"golang-app/app/models/database"
	"golang-app/app/models/user_model"
)

/**
 * Migrates the database.
 */
func Migrate() {
	database.Conn().AutoMigrate(&user_model.User{})
}
