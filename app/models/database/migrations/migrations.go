package migrations

import (
	"alessandromian.dev/golang-app/app/models/database"
	"alessandromian.dev/golang-app/app/models/user_model"
)

/**
 * Migrates the database.
 */
func Migrate() {
	database.Conn().AutoMigrate(&user_model.User{})
}
