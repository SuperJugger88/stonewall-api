package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"stonewall-api/app/models"
)

type MigrationController struct {
	DB *gorm.DB
}

func (controller MigrationController) MakeMigration(ctx *gin.Context) {
	err := controller.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, gin.H{
		"status": "OK",
	})
}
