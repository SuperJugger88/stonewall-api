package application

import (
	"stonewall-api/migrations"
	"github.com/gin-gonic/gin"
)


func HandleMigration(ctx *gin.Context) {
	migrations.Migrate()

	ctx.JSON(200, gin.H{
		"status": "OK",
	  })
}

