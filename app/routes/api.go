package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"stonewall-api/app/controllers"
	config "stonewall-api/config"
	"stonewall-api/middleware"
)

func SetupRouter(db *gorm.DB) {
	router := gin.Default()

	env, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	api := router.Group("/api/v1")
	{
		api.POST("/registration", controllers.RegistrationController{DB: db}.CreateUser)
		api.POST("/login", controllers.AuthController{DB: db}.LoginUser)
		api.POST("/sendMail", controllers.ActivateEmailController{DB: db}.SendMail)
		api.POST("/resetPasswordMail", controllers.ResetPasswordController{DB: db}.SendMail)
		api.POST("/resetPassword", middleware.VerifyMailMiddleware())
		api.POST("/updatePassword", controllers.ResetPasswordController{DB: db}.UpdatePassword)
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

	err = router.Run(env.AppUrl)
	if err != nil {
		panic(err)
	}

}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome!"})
}
