package controllers

import(
	"stonewall-api/domain/service"
	"github.com/gin-gonic/gin"
)

func UpMigrateAction(ctx *gin.Context) {
	migrateService.Migrate()

	ctx.JSON(200, gin.H{
		"status": "OK",
	  })
}