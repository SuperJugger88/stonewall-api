package routes

import (
	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
	"gorm.io/gorm"
	"net/http"
	"os"
	"stonewall-api/app/controllers"
	"stonewall-api/middleware"
)

func SetupRouter(db *gorm.DB) {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/registration", controllers.RegistrationController{DB: db}.CreateUser)
		api.POST("/login", controllers.AuthController{DB: db}.LoginUser)
		api.POST("/sendMail", controllers.MailController{DB: db}.SendMail)
		api.GET("/welcome", middleware.AuthMiddleware(), welcome)
		api.GET("/verifyMail", middleware.VerifyMailMiddleware(), controllers.ActivateEmailController{DB: db}.ActivateEmail)
	}

	protectedGroup := router.Group("/protected")
	protectedGroup.Use(middleware.AuthMiddleware())
	{
		protectedGroup.GET("/data", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "This is protected data"})
		})
	}

	err := router.Run(os.Getenv("API_URL"))
	if err != nil {
		panic(err)
	}

}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome!"})
}
