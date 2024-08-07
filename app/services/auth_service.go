package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
	"stonewall-api/app/middleware"
	"stonewall-api/app/models"
)

func dd(myVar ...interface{}) {
	varDump(myVar...)
	os.Exit(1)
}
func varDump(myVar ...interface{}) {
	fmt.Printf("%v\n", myVar)
}

func AuthenticateUser(email, password string, DB *gorm.DB) (string, error) {

	var user models.User

	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	if err := middleware.VerifyPassword(user.Password, password); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
