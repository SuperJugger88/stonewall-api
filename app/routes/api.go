package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"stonewall-api/app/controllers"
)

func HandleAuthentication(db *gorm.DB) {

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.GET("/migration", controllers.MigrationController{DB: db}.MakeMigration)
	}

	router.Run(os.Getenv("API_URL"))

}
