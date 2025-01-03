package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"os"
	"stonewall-api/app/controllers"
	"stonewall-api/middleware"
	"strconv"
)

func SetupRouter(db *gorm.DB) {
	router := gin.Default()

	cookieStore := cookie.NewStore([]byte(strconv.Itoa(rand.Int())))
	router.Use(sessions.Sessions("stonewall_session", cookieStore))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost", "https://stonewall.io"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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

	err := router.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}

}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome!"})
}
