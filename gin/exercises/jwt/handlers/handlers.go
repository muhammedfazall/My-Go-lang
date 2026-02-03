package handlers

import (
	"jwt/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var UsersDB = map[string]User{}

func Registerh(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "jhsbahv"})
		return
	}

	UsersDB[req.Email] = User{
		Email:    req.Email,
		Password: req.Password,
		Role:     "user",
	}
}

func Loginh(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	user, ok := UsersDB[req.Email]
	if !ok {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	accessToken, err := token.GenerateToken(user.Email,user.Role)
}
