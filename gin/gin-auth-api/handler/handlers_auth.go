package handler

import (
	"go-auth-api/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(ctx *gin.Context) {
	var req LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if req.Username == "" || req.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required!"})
		return
	}

	if len(req.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password must be atleast 6 characters!"})
		return
	}

	if db.UserExists(req.Username) {
		ctx.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	if err := db.CreateUser(req.Username, req.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "registered"})

}

func LoginHandler(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if req.Username == "" || req.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	if !db.VerifyUser(req.Username, req.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	sid, err := db.CreateSession(req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create session"})
		return
	}

	ctx.SetCookie(db.SessionCookieName, sid, 3600, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"status": "logged in"})

}

func DashboardHandler(ctx *gin.Context) {
	username := ctx.GetString("username")
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "welcome to dashboard",
		"username": username,
	})
}

func LogoutHandler(ctx *gin.Context) {
	cookie, err := ctx.Cookie(db.SessionCookieName)
	if err == nil {
		db.DeleteSession(cookie)
	}

	ctx.SetCookie(db.SessionCookieName, "", -1, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"status": "logged out"})
}

