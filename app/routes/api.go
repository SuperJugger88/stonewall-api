package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"stonewall-api/app/controllers"
)

func SetupRouter(db *gorm.DB) {

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.GET("/migration", controllers.MigrationController{DB: db}.MakeMigration)
	}

	err := router.Run(os.Getenv("API_URL"))
	if err != nil {
		panic(err)
	}

}
