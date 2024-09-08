package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"net/http"
	"stonewall-api/app/models"
	"stonewall-api/app/models/dto"
	"stonewall-api/middleware"
	"time"
)

type MailController struct {
	DB *gorm.DB
}

var mySigningKey = []byte("my_secret_key")

func (controller MailController) SendMail(ctx *gin.Context) {

	var emailDto dto.EmailDto
	if err := ctx.ShouldBindJSON(&emailDto); err != nil {
		// Обработка ошибок валидации
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range ve {
				ctx.JSON(400, gin.H{
					"error": fmt.Sprintf("Field %s is %s", fieldErr.Field(), fieldErr.Tag()),
				})
				return
			}
		}
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	user := controller.DB.Where("email = ?", emailDto.Email).First(&existingUser)
	if user == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User not found"})
		return
	}

	email := []string{emailDto.Email}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Токен будет действителен 72 часа

	tokenString, err1 := token.SignedString(mySigningKey)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "dab try create token"})
		return
	}
	messageString := fmt.Sprintf("для подтверждения пароля передите по ссылке ниже: \n http://localhost:80/api/v1/verifyMail?token=%s ", tokenString)

	middleware.SendEmailMiddleware(email, messageString, nil)

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
