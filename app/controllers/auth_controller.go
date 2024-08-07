package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"stonewall-api/app/models/dto"
	"stonewall-api/app/services"
)

type AuthController struct {
	DB *gorm.DB
}

func (controller AuthController) Login(ctx *gin.Context) {
	var loginDTO dto.UserDTO

	if err := ctx.ShouldBindJSON(&loginDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.AuthenticateUser(loginDTO.Email, loginDTO.Password, controller.DB)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}