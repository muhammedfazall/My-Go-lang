package handler

import (
	"log"
	"net/http"
	"time"
	"workspjwt/pkg/jwt"

	"github.com/gin-gonic/gin"
)

var demoUser = struct {
	Email    string
	Password string
}{
	Email:    "test@example.com",
	Password: "1234",
}

func LoginHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if demoUser.Email != req.Email || demoUser.Password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid cred"})
		return
	}

	ttl := time.Hour * 1
	log.Println("emil", req.Email)
	token, err := jwt.GenerateToken("123", req.Email, ttl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
