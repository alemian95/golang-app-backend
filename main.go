package main

import (
	"fmt"

	"alessandromian.dev/golang-app/app/models/database"
	"alessandromian.dev/golang-app/app/models/database/migrations"
	"alessandromian.dev/golang-app/app/router"
	"alessandromian.dev/golang-app/app/router/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	// Initializing router
	r := gin.New()

	// Registering middlewares
	r.Use(middlewares.CORS())
	r.Use(middlewares.Logger())

	// Initializing database
	database.ConnectDatabase()
	// Migrating database
	migrations.Migrate()

	// Registering routes
	router.RegisterRoutes(r)

	// Starting server
	fmt.Println("Server listening on http://localhost:8080")
	r.Run(":8080")
}
