package main

import (
	handler "workspjwt/handlers"
	"workspjwt/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Public route
	r.POST("/login", handler.LoginHandler)

	// Protected group
	auth := r.Group("/api")
	auth.Use(middlewares.JWTAuthMiddleware())
	auth.GET("/profile", handler.ProfileHandler)

	admin := auth.Group("/admin")
	admin.Use(middlewares.RoleMiddleware())
	admin.GET("", handler.AdminDashboard)

	r.Run(":8080")
}
