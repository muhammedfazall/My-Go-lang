package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("session_id")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Login Required"})
			return
		}
		c.Next()
	}
}

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		// path := c.Request.URL.Path
		path := c.FullPath()
		method := c.Request.Method
		ip := c.ClientIP()

		c.Next()

		status := c.Writer.Status()
		delay := time.Since(start)

		fmt.Printf("[%v] - %v - status: %v - ip: %v - duration : %v \n", method, path, status, ip, delay)

	}
}

func VIPonly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole,err := c.Cookie("user_role")
		if err != nil{
			c.AbortWithStatusJSON(401,gin.H{"error":"Access denied!"})
			return
		}

		if userRole != "admin"{
			c.AbortWithStatusJSON(403,gin.H{"error":"Access denied!"})
			return
		}
		c.Next()
	}
}