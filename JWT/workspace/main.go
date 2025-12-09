package main

import (
	handler "workspjwt/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login",handler.LoginHandler)

	r.Run(":8080")
}
