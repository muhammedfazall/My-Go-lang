package main

import (
	"go-auth-api/handler"
	"go-auth-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.LoggingMiddleWare())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "welcome"})
	})

	r.POST("/login", middleware.AuthMiddleWare(), handler.LoginHandler)
	r.POST("/register", handler.RegisterHandler)

	r.GET("/dashboard", handler.DashboardHandler)

	r.GET("/logout", middleware.AuthMiddleWare(), handler.LogoutHandler)

	r.Run(":8080")
}
