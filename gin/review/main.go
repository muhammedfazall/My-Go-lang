package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var sessions = map[string]string{}

const sessionCookieName = "sessions_id"

var demoUser = "fazal"
var demoPass = "fazal123"

func generateSessionId(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func createSession(username string) (string, error) {
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

func LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		method := ctx.Request.Method
		path := ctx.FullPath()
		ip := ctx.ClientIP()

		ctx.Next()

		status := ctx.Writer.Status()
		delay := time.Since(start)

		if path == "/login" || path == "/logout" {
			fmt.Printf("[LOG] %s %s -> %d from %s in %v\n", method, path, status, ip, delay)
		}
	}
}

func main() {
	r := gin.New()

	r.Use(LogMiddleware())

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

		if req.Username != demoUser || req.Password != demoPass {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "inavalid credentials"})
			return
		}

		sId, err := createSession(req.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create"})
			return
		}

		ctx.SetCookie(sessionCookieName, sId, 3600, "/", "", false, true)
		ctx.JSON(http.StatusOK, gin.H{"status": "logged in"})

	})

	r.GET("/dashboard", func(ctx *gin.Context) {
		sid,_ := ctx.Cookie(sessionCookieName)
		// username := ctx.GetString("username")  need to set authmiddleware


		ctx.JSON(http.StatusOK, gin.H{"username":sid})
	})

	r.GET("/logout", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie(sessionCookieName)
		if err == nil {
			deleteSession(cookie)
		}

		ctx.SetCookie(sessionCookieName, "", -1, "/", "", false, true)
		ctx.JSON(http.StatusOK, gin.H{"message": "logged Out"})
	})

	r.Run(":9090")
}
