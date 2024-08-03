package routes

import (
	"github.com/gin-gonic/gin"
	"os"
	"stonewall-api/app/controllers"
)

func HandleAuthentication() {

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.GET("/migration", controllers.MakeMigration)
		api.GET("/register", controllers.RegisterUser)
	}

	router.Run(os.Getenv("API_URL"))

}
