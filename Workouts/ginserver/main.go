package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gte=18"`

	// json:"name" tells Gin to look for "name" in the JSON, not "Name"
    // binding:"required" makes Gin return an error if it's missing
}

func main() {
	r := gin.Default()

	r.GET("/health", healthFn)
	r.GET("/user/:name", greet)

	r.POST("/api/user", createUser)

	if err := r.Run(":8080"); err != nil {
		fmt.Print("could not run server")
		return
	}
}

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
