package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"stonewall-api/infrastructure"
)

func main() {
	infrastructure.InitDatabaseConnection()

	//application.HandleRegistration()

	router := gin.Default()

	router.Run(os.Getenv("API_URL"))
}
