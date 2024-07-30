package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"stonewall-api/infrastructure"
	"stonewall-api/application"
)

func main() {
	infrastructure.InitDatabaseConnection()

	router := gin.Default()

	//application.HandleRegistration()
	router.GET("/api/v1/migration", application.HandleMigration)

	router.Run(os.Getenv("API_URL"))
}
