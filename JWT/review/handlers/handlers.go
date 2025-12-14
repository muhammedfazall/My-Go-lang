package handlers

import (
	"log"
	jwtoken "review/pkg"
	"time"

	"github.com/gin-gonic/gin"
)

var Users = map[string]string{}

func Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}

	Users["email"] = req.Email
	Users["password"] = req.Password

	log.Println(Users)
}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}

	ttl := time.Hour * 1
	token, err := jwtoken.GenerateToken(req.Email, ttl)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cant create token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func Profile(c *gin.Context) {
	claims := c.MustGet("claims")

	c.JSON(200, gin.H{"message": "Welcome to home",
		"email": claims})
}
