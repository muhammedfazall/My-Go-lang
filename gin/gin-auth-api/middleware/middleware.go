package middleware

import (
	"fmt"
	"go-auth-api/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleWare() gin.HandlerFunc {
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

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie(db.SessionCookieName)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorised- please login"})
			return
		}

		if username, ok := db.GetUsernameBySession(cookie); ok {
			ctx.Set("username", username)
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
	}
}
