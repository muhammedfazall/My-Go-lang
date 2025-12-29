package main

import (
	handlers "review2/Handlers"
	"review2/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("login", handlers.Login)

	auth := r.Group("auth")

	auth.Use(middlewares.Authmiddleware())
	auth.GET("profile",handlers.Profile)
	auth.GET("logout",handlers.Logout)

	r.Run(":8080")
}
