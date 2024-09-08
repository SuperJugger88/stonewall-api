package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"net/http"
	"stonewall-api/app/models"
	"time"
)

type ActivateEmailController struct {
	DB *gorm.DB
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
