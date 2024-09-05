package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

var mySigningKey = []byte("my_secret_key")

func VerifyMailMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Query("token")

		// Проверяем, на наличии токена в запросе
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is empty"})
			ctx.Abort()
			return
		}

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return mySigningKey, nil
		})

		//if len(err) == 0 {
		//	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Error parsing token"})
		//	ctx.Abort()
		//	return
		//}

		// Проверяем, является ли токен действительным
		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			ctx.Abort()
			return
		}

		// Проверяем, есть ли у токена почта
		claims := token.Claims.(jwt.MapClaims)
		if claims["email"] != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "the token is not tied to mail"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
