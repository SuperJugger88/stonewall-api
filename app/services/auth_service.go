package services

import (
	"errors"
	"gorm.io/gorm"
	"stonewall-api/app/models"
	middleware2 "stonewall-api/middleware"
)

func AuthenticateUser(email, password string, DB *gorm.DB) (string, error) {

	var user models.User

	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	if err := middleware2.VerifyPassword(user.Password, password); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := middleware2.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
