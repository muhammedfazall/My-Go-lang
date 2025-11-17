package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New() // to create a new gin engine
	r.Use(gin.Logger()) // to use middleware
	fmt.Println("Server Started...")

	r.GET("/home", hey) //path , handler
	r.GET("/users", users)

	r.Run(":8080") // To run server
}

func hey(ctx *gin.Context) {  					
	ctx.JSON(200, gin.H{"message": "hey fazal"})  // gin.H - to convert to hashmap
}

func users(usr *gin.Context) {
	var u = map[string]string{
		"user1": "abdu",
		"user2": "abdu",
		"user3": "abdu",
		"user4": "abdu",
		"user5": "abdu",
		"user6": "abdu",
	}
	usr.JSON(200, u)
}


