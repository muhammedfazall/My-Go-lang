package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func healthFn(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP"})
}

func greet(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{"message": "hello " + name})
}

func createUser(c *gin.Context) {
	var req User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	// // 	Struct Tags: They are a powerful Go feature that lets you tell Gin
	// // 	exactly how to handle your data without writing all those if checks.

	// // In Go, you can add "metadata" to struct fields using backticks.
	// // Gin uses the binding tag to perform validation automatically.

	// if req.Name == "" || req.Age == 0 {
	// 	c.JSON(400,gin.H{"error":"name and age required!"})
	// 	return
	// }

	c.JSON(201, gin.H{"message": "User" + req.Name + "created!"})
}

func loginHandler(c *gin.Context) {
	var req loginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	err := bcrypt.CompareHashAndPassword(storedHash, []byte(req.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	c.SetCookie("session_id", "secret-token", 3600, "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "login Successfull"})
}

func logout(c *gin.Context) {
	c.SetCookie("session_id", "secret-token", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "logged out"})
}

func dashboardHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome to dashboard!"})
}

func registerhandler(c *gin.Context) {
	var req User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "error hashing password!"})
		return
	}

	storedHash = hashed
}
