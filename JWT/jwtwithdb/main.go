package main

import (
	"jwtwithdb/handlers"
	"jwtwithdb/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login",handlers.Login)
	r.POST("/register",handlers.Register)

	auth := r.Group("auth")
	auth.Use(middlewares.AuthMiddleware())
	auth.GET("/profile",handlers.Profile)

	r.Run(":8080")
}