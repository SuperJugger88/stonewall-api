package validate

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"stonewall-api/app/models"
	"stonewall-api/app/models/dto"
)

func ExistUserEmail(DB *gorm.DB, emailDto dto.EmailDto, ctx *gin.Context) {
	var existingUser models.User
	user := DB.Where("email = ?", emailDto.Email).First(&existingUser)
	if user == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User not found"})
		return
	}
}
