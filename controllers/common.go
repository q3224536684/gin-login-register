package controllers

import (
	"frank/gin-login-register/models"
	"github.com/gin-gonic/gin"
)

func SayHello(ctx *gin.Context) {
	models.Result(200, "SUCCESS", "Hello World!", ctx)
}
