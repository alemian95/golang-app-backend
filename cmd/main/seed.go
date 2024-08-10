package main

import (
	"golang-app/app/models/database/migrations"
	"golang-app/cmd"
)

func main() {
	cmd.InitCmdApplication()
	migrations.Seed()
}
