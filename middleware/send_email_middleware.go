package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/smtp"
	"stonewall-api/config"
)

func SendEmailMiddleware(email []string, stringEmail string, ctx *gin.Context) {
	env, err := config.LoadConfig(".")
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "failed to receive file0"})
	}

	message := []byte(stringEmail)

	err = smtp.SendMail(env.SmptHost+":"+env.SmptPort, nil, env.FromEmail, email, message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send email"})
		return
	}
}
