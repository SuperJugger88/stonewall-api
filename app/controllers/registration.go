package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"stonewall-api/app/models"
)

type RegistrationController struct {
	DB *gorm.DB
}

func (controller RegistrationController) CreateUser(ctx *gin.Context) {
	var userDTO models.UserDTO

	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := controller.DB.Where("email = ?", userDTO.Email).First(&existingUser).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	hashedPassword, err := models.HashPassword(userDTO.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	// Преобразование DTO в модель пользователя
	user := &models.User{
		Password: hashedPassword,
		Email:    userDTO.Email,
	}

	// Сохранение пользователя в базе данных
	if err := controller.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Отправка ответа клиенту
	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
