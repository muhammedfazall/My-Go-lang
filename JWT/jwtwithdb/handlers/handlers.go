package handlers

import (
	"jwtwithdb/model"
	tokens "jwtwithdb/pkg/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var UsersDb = map[string]model.User{}

func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not hashpassword: " + err.Error()})
		return
	}

	UsersDb[req.Email] = model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashed),
		Role:     "user",
	}

	c.JSON(200, gin.H{"message": "registered successfully"})
}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}

	user, ok := UsersDb[req.Email]
	if !ok {
		c.JSON(401, gin.H{"error": "user not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid password"})
		return
	}

	accessToken, refreshToken, err := tokens.GenerateTokens(user.Email, user.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not create tokens "+err.Error()})
		return
	}

	c.SetCookie("access_token",accessToken,900,"/","",false,true)

	c.JSON(200, gin.H{"message": "logged in",
		"access":  accessToken,
		"refresh": refreshToken})
}

func Profile(c *gin.Context) {
	claims := c.MustGet("claims").(*tokens.CustomClaims)
	c.JSON(200, gin.H{"welcome": claims.Email})

}
