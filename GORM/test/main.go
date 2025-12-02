package main

import (
	"test/database"
	"test/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	r := gin.Default()

	r.POST("/register",handlers.Register)
	r.POST("/login",handlers.Login)
	r.GET("/home",handlers.Home)

	r.Run(":8080")

}