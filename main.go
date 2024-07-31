package main

import (
	"stonewall-api/infrastructure"
	"stonewall-api/routes"
)

func main() {
	infrastructure.InitDatabaseConnection()

	routes.GroupUserUrl()

}
