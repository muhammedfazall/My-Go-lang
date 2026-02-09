package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {

	r := gin.Default()

	r.POST("/users", handleCreateUser)
	r.POST("/login", handleLogin)

	auth := r.Group("/auth")
	auth.Use(AuthMiddleware())

	auth.GET("/profile", func(c *gin.Context) {
		email, _ := c.Get("UserEmail")
		c.JSON(200, gin.H{"message": "success", "email": email})
	})

	r.Run(":8080")
}

func handleLogin(c *gin.Context) {
	var req UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json format"})
		return
	}

	if req.Email != "admin@go.com" || req.Password != "password123" {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	accessTkn, err := GenerateToken(req.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not create token"})
		return
	}

	c.JSON(200, gin.H{"token": accessTkn})
}

func handleCreateUser(c *gin.Context) {
	var user UserInput

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Validation failed: "})
		return
	}

	c.JSON(201, gin.H{"message": "User validation passed!", "data": user.Username})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authheader := c.GetHeader("Authorization")

		if authheader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is missing"})
			return
		}

		parts := strings.SplitN(authheader," ",2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer"{
			c.AbortWithStatusJSON(401,gin.H{"error":"Authorization header is missing"})
			return
		}

		tokenString := parts[1]

		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		c.Set("UserEmail", claims.Email)

		c.Next()
	}
}
