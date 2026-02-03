package main

import (
	handler "exercises/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/register",handler.Register)
	r.POST("/login",handler.Login)

	r.GET("/logout",handler.Logout)
}
