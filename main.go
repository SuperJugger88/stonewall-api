package main

import (
	"stonewall-api/app/routes"
	database "stonewall-api/config"
)

func main() {
	db := database.InitDatabaseConnection()

	routes.HandleMigration(db)
}
