package routers

import (
	"context"
	"frank/gin-login-register/controllers"
	"frank/gin-login-register/middles"
	"frank/gin-login-register/services"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(ctx context.Context, userService services.UserService, router *gin.RouterGroup) {
	controllersImpl := controllers.NewUserControllersImpl(ctx, userService)
	router.POST("/register", controllersImpl.Register)
	authMiddleware, err := middles.InitAuthMiddlewares(controllersImpl)
	if err != nil {
		panic(err)
	}
	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/refresh_token", authMiddleware.RefreshHandler)

	{
		auth := router.Use(authMiddleware.MiddlewareFunc())
		auth.POST("/logout", authMiddleware.LogoutHandler)
		auth.GET("/list", controllersImpl.UserList)
	}
}
