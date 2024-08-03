package controllers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func RegisterUser(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, gin.H{
		"status": "OK",
		"body":   body,
	})
}
