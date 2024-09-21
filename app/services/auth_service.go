package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"stonewall-api/app/models"
	"stonewall-api/middleware"
)

func AuthenticateUser(email, password string, DB *gorm.DB, ctx *gin.Context) (string, error) {

	var user models.User

	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	if err := middleware.VerifyPassword(user.Password, password); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := middleware.GenerateJWT(user.ID, nil)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	
	return token, nil
}
