package main

import (
	"fmt"
	"log"
	"os"

	"golang-app/app/models/database"
	"golang-app/app/models/database/migrations"
	"golang-app/app/router"
	"golang-app/app/router/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment")
	}

	// Initializing router
	r := gin.New()

	// Registering middlewares
	r.Use(middlewares.CORS())
	r.Use(middlewares.Logger())
	r.Use(middlewares.CSRF())

	// Initializing database
	database.ConnectDatabase()
	// Migrating database
	migrations.Migrate()

	// Registering routes
	router.RegisterRoutes(r)

	// Starting server
	fmt.Println(fmt.Sprintf("Server listening on http://localhost:%s", os.Getenv("APP_PORT")))
	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
