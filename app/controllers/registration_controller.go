package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"stonewall-api/app/models"
	"stonewall-api/app/models/dto"
	"stonewall-api/middleware"
)

type RegistrationController struct {
	DB *gorm.DB
}

func (controller RegistrationController) CreateUser(ctx *gin.Context) {
	err := controller.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	var userDTO dto.UserDTO

	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := controller.DB.Where("email = ?", userDTO.Email).First(&existingUser).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, err := middleware.HashPassword(userDTO.Password)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	// Преобразование DTO в модель пользователя
	user := &models.User{
		Email:    userDTO.Email,
		Password: hashedPassword,
	}

	// Сохранение пользователя в базе данных
	if err := controller.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправка ответа клиенту
	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
		"email":   userDTO.Email,
	})
}
