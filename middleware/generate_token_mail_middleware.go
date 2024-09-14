package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func GenerateTokenMailMiddleware(email []string, ctx *gin.Context) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Токен будет действителен 72 часа

	tokenString, err1 := token.SignedString(mySigningKey)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "dab try create token"})
		return ""
	}

	return tokenString
}
