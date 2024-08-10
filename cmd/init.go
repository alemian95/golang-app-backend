package cmd

import (
	"golang-app/app/models/database"
	"log"

	"github.com/joho/godotenv"
)

func InitCmdApplication() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment")
	}

	database.ConnectDatabase()
}
