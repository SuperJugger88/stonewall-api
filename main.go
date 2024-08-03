package main

import (
	"stonewall-api/app/routes"
	database "stonewall-api/services"
)

func main() {
	database.InitDatabaseConnection()

	routes.HandleAuthentication()

}
