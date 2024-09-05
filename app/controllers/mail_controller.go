package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"net/http"
	"net/smtp"
	"time"
)

type MailController struct {
	DB *gorm.DB
}

var mySigningKey = []byte("my_secret_key")

type EmailRequestBody struct {
	Email string
}

func (controller MailController) SendMail(ctx *gin.Context) {

	// Информация об отправителе
	from := "from@gmail.com"
	// smtp сервер конфигурация
	smtpHost := "mailhog"
	smtpPort := "1025"
	// Информация о получателе
	to := []string{
		"sender@example.com",
	}

	var requestBody EmailRequestBody

	if err := ctx.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
	}

	fmt.Println(requestBody.Email)

	//var existingUser models.User

	ctx.JSON(http.StatusOK, gin.H{
		"success": requestBody,
	})

	token := jwt.New(jwt.SigningMethodHS256)

	// Устанавливаем claims (полезные данные)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = "sender@example.com"
	claims["exp"] = time.Now().Add(time.Second * 1).Unix() // Токен будет действителен 72 часа

	// Подписываем токен с помощью секретного ключа
	tokenString, err1 := token.SignedString(mySigningKey)
	if err1 != nil {
		fmt.Println("Error signing token:", err1)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": tokenString,
	})
	// Сообщение.
	message := []byte(fmt.Sprintf("для подтверждения пароля передите по ссылке ниже: \n http://localhost:80/api/v1/verifyMail/%s ", tokenString))

	// Отправка почты.
	err := smtp.SendMail(smtpHost+":"+smtpPort, nil, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Почта отправлена!")

	ctx.JSON(http.StatusOK, gin.H{
		"success": "tru2e",
	})
}
