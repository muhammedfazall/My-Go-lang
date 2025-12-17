package handlers

import (
	"log"
	jwtoken "review/pkg"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint
	Email    string
	Password string
	Role     string
}

var mockDBUsers = map[string]User{}

var nextUserId uint = 1

func Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	role := "user"
	if req.Email == "admin@example.com" {
	role = "admin"
	}

	newUser := User{
		ID:       nextUserId,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     role,
	}

	mockDBUsers[req.Email] = newUser
	nextUserId++

	// Users[req.Email] = string(hashedPassword)

	// log.Println(Users)

	c.JSON(201, gin.H{"message": role +" registered succesfully", "Id": newUser.ID})

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

	user, ok := mockDBUsers[req.Email]
	if !ok {
		c.JSON(401, gin.H{"error": "Invalid email"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid password"})
		return
	}

	log.Println("ROLE AT TOKEN GENERATION:", user.Role)
	// ttl := time.Hour * 1
	accessToken, refreshToken, err := jwtoken.GenerateAccessAndRefreshTokens(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cant create token"})
		return
	}

	c.JSON(200, gin.H{"Accesstoken": accessToken, "RefreshToken": refreshToken})
}

func Profile(c *gin.Context) {
	claims := c.MustGet("claims").(*jwtoken.CustomClaims)

	c.JSON(200, gin.H{"message": "Welcome to home",
		"email": claims.Email})
}

func Admin(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome admin mf"})
}
