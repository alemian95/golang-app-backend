package main

import (
	"fmt"

	"alessandromian.dev/golang-app/app/models/database"
	"alessandromian.dev/golang-app/app/router"
	"alessandromian.dev/golang-app/app/router/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Database *gorm.DB

func main() {
	r := gin.New()
	r.Use(middlewares.CORS())
	r.Use(middlewares.Logger())

	database.ConnectDatabase()

	router.RegisterRoutes(r)

	fmt.Println("Server listening on http://localhost:8080")
	r.Run(":8080")
}
