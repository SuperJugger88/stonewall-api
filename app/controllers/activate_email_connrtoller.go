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
	"stonewall-api/app/models/validate"
	"stonewall-api/middleware"
	"time"
)

type ActivateEmailController struct {
	DB *gorm.DB
}

func (controller ActivateEmailController) SendMail(ctx *gin.Context) {
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

	validate.ExistUserEmail(controller.DB, emailDto, nil)

	email := []string{emailDto.Email}

	tokenString := middleware.GenerateTokenMailMiddleware(email, nil)
	messageString := fmt.Sprintf("для подтверждения пароля передите по ссылке ниже: \n http://localhost:80/api/v1/verifyMail?token=%s ", tokenString)

	middleware.SendEmailMiddleware(email, messageString, nil)

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}

func (controller ActivateEmailController) ActivateEmail(ctx *gin.Context) {

	tokenString := ctx.Query("token")
	token, _ := jwt.Parse(tokenString, nil)
	claims := token.Claims.(jwt.MapClaims)

	user := models.User{}
	result := controller.DB.Model(&user).Where("email = ?", claims["email"]).Update("ActivatedAt", time.Now())
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, result.Error)
		return
	}
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no user found with email"})
		return
	}
}
