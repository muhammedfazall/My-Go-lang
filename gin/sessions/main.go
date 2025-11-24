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

var demoHash []byte

var sessions = map[string]string{}

const sessionCookieName = "session_id"

func init() {
	h, err := bcrypt.GenerateFromPassword([]byte(demoPass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	demoHash = h
}

func generateSessionID(nBytes int) (string, error) {
	b := make([]byte, nBytes)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func createSession(username string) (string, error) {
	id, err := generateSessionID(16)
	if err != nil {
		return "", err
	}
	sessions[id] = username
	return id, nil
}

func getUsernameBySession(id string) (string, bool) {
	username, ok := sessions[id]
	return username, ok
}

func deleteSession(id string) {
	delete(sessions, id)
}

const demoUser = "admin"
const demoPass = "password123"

func loggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		method := ctx.Request.Method
		path := ctx.FullPath()
		ip := ctx.ClientIP()
		t := time.Now()
		ctx.Next()
		delay := time.Since(t)
		status := ctx.Writer.Status()

		if path == "/login" || path == "/logout" {
			fmt.Printf("%s %s -> %d from %s - %d\n", method, path, status, ip, delay)
		}

	}
}

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie(sessionCookieName)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unAuthorised,please login"})
			return
		}
		if username, ok := getUsernameBySession(cookie); ok {
			ctx.Set("username", username)
			ctx.Next()
			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
	}
}

func main() {
	r := gin.New()

	r.Use(loggingMiddleware())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello - visit/login to get started"})
	})

	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	r.POST("/login", func(ctx *gin.Context) {

		var req loginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
			return
		}

		if req.Username == "" || req.Password == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "usernamea and password are required"})
			return
		}

		if len(req.Password) < 6 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 6 characters"})
			return
		}

		if req.Username != demoUser {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		if err := bcrypt.CompareHashAndPassword(demoHash, []byte(req.Password)); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		sid, err := createSession(req.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create session"})
			return
		}

		ctx.SetCookie(sessionCookieName, sid, 3600, "/", "", false, true)
		ctx.JSON(http.StatusOK, gin.H{"status": "logged in"})
	})

	r.GET("/dashboard", authMiddleware(), func(ctx *gin.Context) {
		username := ctx.GetString("username")
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "welcome to dashboard",
			"username": username,
		})
	})

	r.GET("/logout", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie(sessionCookieName)

		if err == nil {
			deleteSession(cookie)
		}

		ctx.SetCookie(sessionCookieName, "", -1, "/", "", false, true)
		ctx.JSON(http.StatusOK, gin.H{"status": "logged out"})

	})
	r.Run(":8080")
}
