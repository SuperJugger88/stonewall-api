package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"stonewall-api/app/models"
	"stonewall-api/app/models/dto"
	"stonewall-api/app/services"
)

type AuthController struct {
	DB *gorm.DB
}

func (controller AuthController) LoginUser(ctx *gin.Context) {
	err := controller.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	var loginDTO dto.UserDTO

	if err := ctx.ShouldBindJSON(&loginDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.AuthenticateUser(loginDTO.Email, loginDTO.Password, controller.DB, ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	session := sessions.Default(ctx)
	session.Set(loginDTO.Email, token)
	session.Options(sessions.Options{
		MaxAge:   3600,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		Domain:   ".localhost",
	})
	err = session.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"email":   loginDTO.Email,
	})
}
