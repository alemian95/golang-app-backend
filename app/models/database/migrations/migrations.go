package migrations

import (
	"golang-app/app/models/database"
	"golang-app/app/models/forgot_password_token_model.go"
	"golang-app/app/models/user_model"
	"golang-app/app/utils/auth"
)

/**
 * Migrates the database.
 */
func Migrate() {
	database.Conn().AutoMigrate(&user_model.User{})
	database.Conn().AutoMigrate(&forgot_password_token_model.ForgotPasswordToken{})
}

func Seed() {

	hash, _ := auth.HashPassword("example")

	user := user_model.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: hash,
	}

	user.Create()
}
