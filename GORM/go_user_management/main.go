package main

import (
	"go-user-management/controllers"
	"go-user-management/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	router := gin.Default()
	router.GET("/users", controllers.GetUsers)

	router.POST("/users", controllers.CreateUser)

	router.GET("/users/:id", controllers.GetUserById)

	router.PUT("/users/:id", controllers.UpdateUser)

	router.DELETE("/users/:id", controllers.DeleteUser)

	log.Println("Server running on :8080")
	router.Run(":8080")

}
