package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

var sessions = map[string]string{}

const sessionCookieName = "session_id"

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
	r := gin.Default()

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

		if req.Username != demoUser || req.Password != demoPass{
			ctx.JSON(http.StatusUnauthorized,gin.H{"error":"inavlid credentials"})
			return
		}

		sid,err := createSession(req.Username)
		if err != nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":"could not create session"})
			return
		}

		ctx.SetCookie(sessionCookieName,sid,3600,"/","",false,true)
		ctx.JSON(http.StatusOK,gin.H{"status":"logged in"})
	})

	r.GET("/dashboard",authMiddleware(),func(ctx *gin.Context) {
		username:= ctx.GetString("username")
		ctx.JSON(http.StatusOK,gin.H{
			"message": "welcome to dashboard",
			"username": username,
		})
	})

	r.GET("/logout", func(ctx *gin.Context) {
		cookie,err := ctx.Cookie(sessionCookieName)

		if err == nil{
			deleteSession(cookie)
		}

		ctx.SetCookie(sessionCookieName,"",-1,"/","",false,true)
		ctx.JSON(http.StatusOK,gin.H{"status":"logged out"})

	})
	r.Run(":8080")
}
