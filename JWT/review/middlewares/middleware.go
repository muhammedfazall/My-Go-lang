package middlewares

import (
	"log"
	jwtoken "review/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authheader := c.GetHeader("Authorization")
		if authheader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "token missing"})
			return
		}

		log.Println(authheader)

		parts := strings.SplitN(authheader, " ", 2)

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		tokenString := parts[1]

		claims,err := jwtoken.ValidateToken(tokenString) 
		if err != nil{
			c.AbortWithStatusJSON(401,gin.H{"error":err.Error()})
			return
		}
		
		c.Set("claims",claims)
		c.Next()
	}
}

func RoleMiddleWare(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims := ctx.MustGet("claims").(*jwtoken.CustomClaims)

		if claims.Role != requiredRole {
			ctx.AbortWithStatusJSON(403,gin.H{"error":"Access Denied: You are not "+ requiredRole})
			return
		}

		ctx.Next()
	}
}
