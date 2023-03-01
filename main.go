package main

import (
	"context"
	"frank/gin-login-register/config"
	"frank/gin-login-register/middles"
	"frank/gin-login-register/routers"
	"frank/gin-login-register/services"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"log"
)

var ctx context.Context

func main() {
	ctx = context.Background()
	config.InitConfig(".")
	db := config.GetDb()

	userService := services.NewUserServiceImpl(db, ctx)
	userService.InitTable()

	// gin
	server := gin.Default()
	server.Use(requestid.New())
	server.Use(middles.AddCors())

	api := server.Group("/api")
	routers.InitCommonRouter(api)
	routers.InitUserRouter(ctx, userService, api)

	// run
	log.Fatal(server.Run(":" + config.GetConfig().Port))
}
