package main

import (
	"fmt"

	"alessandromian.dev/golang-app/app/models/database"
	"alessandromian.dev/golang-app/app/models/user_model"
	"alessandromian.dev/golang-app/app/router"
	"alessandromian.dev/golang-app/app/router/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func main() {
	r := gin.New()
	r.Use(middlewares.CORS())
	r.Use(middlewares.Logger())

	Database, dbErr := gorm.Open(postgres.Open("host=localhost user=db_user password=example dbname=golang_db port=5432"), &gorm.Config{})

	if dbErr != nil {
		panic(dbErr)
	} else {
		Database.AutoMigrate(&user_model.User{})
		database.RegisterControllersDatabase(Database)
	}

	router.RegisterRoutes(r)

	fmt.Println("Server listening on http://localhost:8080")
	r.Run(":8080")
}
