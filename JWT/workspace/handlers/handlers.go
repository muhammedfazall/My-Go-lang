package handler

import (
	"net/http"
	"time"
	myjwt "workspjwt/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func LoginHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	// simple login validation
	if req.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong email/password"})
		return
	}

	role := "user"
	if req.Email == "admin@example.com" {
		role = "admin"
	}

	// generate JWT (valid for 1 hour)
	ttl := time.Hour * 1
	token, err := myjwt.GenerateToken(role, req.Email, ttl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func ProfileHandler(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome!",
		"role":    claims["role"],
		"email":   claims["email"],
	})
}

func AdminDashboard(c *gin.Context) {
	c.JSON(200, gin.H{"message": "welcome to ad dashbrd"})
}
