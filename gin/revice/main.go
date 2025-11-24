package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var users = map[string]string{}
var sessions = map[string]string{}

const SessionCookieName = "sessions_id"

func CreateUser(username, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	users[username] = string(hashed)
	return nil
}

func UserExists(username string) bool {
	_, ok := users[username]
	return ok
}

func VerifyUser(username, password string) bool {
	hashed, ok := users[username]
	if !ok {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false
	}
	return true
}

func generateSessionId(nByte int) (string, error) {
	b := make([]byte, nByte)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func CreateSession(username string) (string, error) {
	id, err := generateSessionId(16)
	if err != nil {
		return "", err
	}

	sessions[id] = username
	return id, nil
}

func deleteSession(id string) {
	delete(sessions, id)
}

func loggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		method := ctx.Request.Method
		path := ctx.FullPath()
		ip := ctx.ClientIP()

		ctx.Next()
		delay := time.Since(start)

		status := ctx.Writer.Status()

		if path == "/login" || path == "/logout" {
			fmt.Printf("[log] %s %s -> %d from %s in %v\n", method, path, status, ip, delay)
		}
	}
}

func main() {
	r := gin.New()

	r.Use(loggingMiddleware())

	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	r.POST("/register", func(ctx *gin.Context) {
		var req loginRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
			return
		}

		if req.Username == "" || req.Password == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required!"})
			return
		}

		if UserExists(req.Username) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			return
		}

		if err := CreateUser(req.Username, req.Password); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "registered"})
	})

	r.POST("/login", func(ctx *gin.Context) {
		var req loginRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
			return
		}

		if !VerifyUser(req.Username, req.Password) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		sid, err := CreateSession(req.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create session"})
			return
		}

		ctx.SetCookie(SessionCookieName, sid, 3600, "/", "", false, true)
		ctx.JSON(http.StatusOK, gin.H{"status": "logged in"})

	})

	r.GET("/dashboard", func(ctx *gin.Context) {
		username := ctx.GetString("username")
		ctx.JSON(http.StatusOK, gin.H{"message": "welcome",
			"username": username})
	})

	r.GET("/logout", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie(SessionCookieName)
		if err == nil {
			deleteSession(cookie)
		}

		ctx.SetCookie(SessionCookieName, "", -1, "/", "", false, true)
		ctx.JSON(http.StatusOK, gin.H{"message": "loged out"})
	})

	r.Run(":8080")
}
