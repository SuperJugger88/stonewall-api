package routes

import "github.com/gin-gonic/gin"

func GetMainPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
