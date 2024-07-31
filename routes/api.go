package routes

import (
	"stonewall-api/domain/controllers"
	"github.com/gin-gonic/gin"
	"os"

)


func GroupUserUrl() {

	router := gin.Default()


	api := router.Group("/api/v1")
	{
		api.GET("/api/v1/migration", controllers.UpMigrateAction)
	}


	api1 := router.Group("/api/v1")
	{
		api1.GET("/register", controllers.Register)
	}
	
	router.Run(os.Getenv("API_URL"))

}