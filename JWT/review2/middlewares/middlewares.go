package middlewares

import (
	tokens "review2/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func Authmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken,err := c.Cookie("access")
		if err != nil{
			c.AbortWithStatusJSON(401,gin.H{"error":"Token not found"})
			return
		}

		if accessToken == ""{
			c.AbortWithStatusJSON(401,gin.H{"error":"Token not found"})
			return
		}

		claims,err := tokens.ValidateToken(accessToken)
		if err != nil{
			c.AbortWithStatusJSON(401,gin.H{"error":"invalid token"})
		}

		c.Set("claims",claims)
		c.Next()
	}
}