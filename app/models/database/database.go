package database

import (
	"alessandromian.dev/golang-app/app/controllers/index_controller"
	"alessandromian.dev/golang-app/app/controllers/user_controller"
	"gorm.io/gorm"
)

func RegisterControllersDatabase(db *gorm.DB) {
	index_controller.RegisterDatabase(db)
	user_controller.RegisterDatabase(db)
}
