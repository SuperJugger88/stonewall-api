package controllers

import(
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func Register(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)

	ctx.JSON(200, gin.H{
		"status": "OK",
		"body": body,
	  })
}