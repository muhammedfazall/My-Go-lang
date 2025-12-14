package main

import (
	"review/handlers"
	"review/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/register", handlers.Register)

	r.POST("/login", handlers.Login)
	auth := r.Group("api")

	auth.Use(middlewares.AuthMiddleware())

	auth.GET("/profile", handlers.Profile)

	r.Run(":8080")

}
