package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/smtp"
	"os"
)

func SendEmailMiddleware(email []string, stringEmail string, ctx *gin.Context) {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	emailFrom := os.Getenv("EMAIL_FROM")

	message := []byte(stringEmail)

	err := smtp.SendMail(smtpHost+":"+smtpPort, nil, emailFrom, email, message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send email"})
		return
	}
}
