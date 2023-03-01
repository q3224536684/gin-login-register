package routers

import (
	"frank/gin-login-register/controllers"
	"github.com/gin-gonic/gin"
)

func InitCommonRouter(r *gin.RouterGroup) {
	r.GET("/hello", controllers.SayHello)
}
