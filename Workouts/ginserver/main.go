package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var storedHash []byte

func main() {
	r := gin.New()
	r.Use(MyLogger())

	r.GET("/health", healthFn)
	r.GET("/user/:name", greet)

	r.POST("/api/user", createUser)

	r.POST("/register", registerhandler)
	r.POST("/login", loginHandler)

	auth := r.Group("/auth")
	auth.Use(AuthRequired())

	auth.GET("/dashboard", dashboardHandler)
	auth.GET("/logout", logout)

	if err := r.Run(":8080"); err != nil {
		fmt.Print("could not run server")
		return
	}
}
