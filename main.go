package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"stonewall-api/config"
	"stonewall-api/routes"
)

func main() {
	config.InitDatabaseConnection()

	router := gin.Default()

	router.GET("/api/v1/test", routes.GetMainPage)

	router.Run(os.Getenv("API_URL"))
}
