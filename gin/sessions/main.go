package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

var sessions = map[string]string{}

var sessionCookieName = "session_id"

func generateSessionId(newB int) (string, error) {
	b := make([]byte, newB)
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

func deleteSession(id string)  {
	delete(sessions,id)
}

var demoUser = "fazal"
var demoPass = "fazal123"

func authMiddleware() gin.HandlerFunc  {
	return func(ctx *gin.Context) {

	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome - login to get started!"})
	})

	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	r.POST("/login", func(ctx *gin.Context) {
		var req loginRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		if req.Username != demoUser || req.Password != demoPass {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		ssid, err := createSession(req.Username)
		if err != nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":"could not create session"})
			return
		}

		ctx.SetCookie(sessionCookieName,ssid,3600,"/","",false,true)
		ctx.JSON(http.StatusOK,gin.H{"status":"logged in"})

	})

	r.GET("/dashboard",func(ctx *gin.Context) {

	})

	r.GET("/logout",func(ctx *gin.Context) {
		cookie,err := ctx.Cookie(sessionCookieName)
		if err == nil{
			deleteSession(cookie)
		}

		ctx.SetCookie(sessionCookieName,"",-1,"/","",false,true)
		ctx.JSON(http.StatusOK,gin.H{"status":"logged out"})
	})

	r.Run(":8080")
}
