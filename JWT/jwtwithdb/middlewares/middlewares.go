package middlewares

import (
	tokens "jwtwithdb/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authheader := ctx.GetHeader("Authorization")

		if authheader == ""{
			ctx.AbortWithStatusJSON(401,gin.H{"error":"token missing"})
			return
		}

		parts := strings.SplitN(authheader," ",2)

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer"{
			ctx.AbortWithStatusJSON(401,gin.H{"error":"invalid token"})
			return
		}

		tokenString := parts[1]

		claims,err := tokens.ValidateToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(401,gin.H{"error":"invalid token"})
			return
		}

		ctx.Set("claims",claims)
		ctx.Next()
	}
}
