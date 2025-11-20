package main

import "github.com/gin-gonic/gin"

// Struct to receive JSON data for POST /users
type UserInput struct {
	Name string `json:"name"` // maps JSON field "name" to this struct field
}

func main() {
	// Creates a router with default middleware:
	// - Logger (logs every request)
	// - Recovery (handles panics)
	router := gin.Default()

	// Basic GET route for homepage
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "welcome"})
	})

	// Grouping all API routes under /api
	api := router.Group("/api")

	// Simple health-check endpoint
	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	// GET /api/users — returns list of users
	api.GET("/users", func(ctx *gin.Context) {
		// Sending JSON array using []gin.H
		ctx.JSON(200, []gin.H{
			{"id": 1, "name": "Fazal"},
			{"id": 2, "name": "Ali"},
		})
	})

	// POST /api/users — creates a user
	api.POST("/users", func(ctx *gin.Context) {
		var input UserInput

		// BindJSON reads request JSON and fills the struct
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{"error": "invalid input"})
			return
		}

		// Successful response
		ctx.JSON(200, gin.H{
			"status": "user created",
			"name":   input.Name,
		})
	})

	// Start server on port 8080
	router.Run(":8080")
}


// package main

// import (
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	// gin.New() creates a new Gin router WITHOUT any default middleware.
// 	// (Unlike gin.Default() which adds Logger + Recovery automatically)
// 	r := gin.New()

// 	// Adding Logger middleware manually.
// 	// This logs each incoming request to the console.
// 	r.Use(gin.Logger())

// 	fmt.Println("Server Started...")

// 	// Registering routes:
// 	// r.GET("<path>", <handler function>)
// 	r.GET("/home", hey)   // When user visits /home → run hey()
// 	r.GET("/users", users) // When user visits /users → run users()

// 	// Start the server on port 8080
// 	r.Run(":8080")
// }

// // Handler function for GET /home
// func hey(ctx *gin.Context) {
// 	// Responds with JSON object:
// 	// { "message": "hey fazal" }
// 	//
// 	// gin.H is a shortcut for map[string]interface{}
// 	ctx.JSON(200, gin.H{"message": "hey fazal"})
// }

// // Handler function for GET /users
// func users(usr *gin.Context) {

// 	// Creating a map representing multiple users
// 	var u = map[string]string{
// 		"user1": "abdu",
// 		"user2": "abdu",
// 		"user3": "abdu",
// 		"user4": "abdu",
// 		"user5": "abdu",
// 		"user6": "abdu",
// 	}

// 	// Returning the map as JSON
// 	usr.JSON(200, u)
// }
