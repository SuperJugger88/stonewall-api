package controllers

import (
	"github.com/gin-gonic/gin"
	database "stonewall-api/services"
)

func MakeMigration(ctx *gin.Context) {
	database.InitDatabaseConnection()

	ctx.JSON(200, gin.H{
		"status": "OK",
	})
}
