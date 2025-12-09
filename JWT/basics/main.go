package main

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Very simple example user (in real world -> DB)
var demoUser = struct {
	Role string
	Email    string
	Password string
}{
	Role: "user",
	Email:    "test@example.com",
	Password: "1234",
}

var demoAdmin = struct {
	Role string
	Email    string
	Password string
}{
	Role: "admin",
	Email:    "admin@example.com",
	Password: "1234",
}

var secretKey = []byte("secret")

func main() {
	r := gin.Default()

	// Public route
	r.POST("/login", loginHandler)

	// Protected group
	auth := r.Group("/api")
	auth.Use(JWTAuthMiddleware())
	auth.GET("/profile", profileHandler)
	
	admin := auth.Group("/admin")
	admin.Use(RoleMiddleware())
	admin.GET("",adminDashboard)

	r.Run(":8080")
}

// ------------------- LOGIN ---------------------

func loginHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	// simple login validation
	if req.Password != demoUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong email/password"})
		return
	}

	role := "user"
	if req.Email == "admin@example.com" {
		role = "admin"
	}

	// generate JWT (valid for 1 hour)
	ttl := time.Hour * 1
	token, err := GenerateToken(role,req.Email, ttl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ------------------- MIDDLEWARE ---------------------

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		// Expect header in format: "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// store claims for next handlers
		c.Set("claims", claims)
		c.Next()
	}
}

func RoleMiddleware()gin.HandlerFunc{
	return func(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)
		
	role := claims["role"]

	if role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error": "You are not an admin"})
		return 
	}

	c.Next()
	}
}

// ------------------- PROTECTED ENDPOINT ---------------------

func profileHandler(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome!",
		"role": claims["role"],
		"email":   claims["email"],
	})
}

func adminDashboard(c *gin.Context)  {
	c.JSON(200,gin.H{"message":"welcome to ad dashbrd"})
}

// ------------------- JWT HELPERS ---------------------

func GenerateToken(role, email string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"role": role,
		"email":   email,
		"exp":     time.Now().Add(ttl).Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(secretKey)
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		// Ensure that the token's signing method is HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// simple exp check
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("invalid exp in token")
	}
	if time.Now().Unix() > int64(exp) {
		return nil, errors.New("expired token")
	}

	return claims, nil
}
