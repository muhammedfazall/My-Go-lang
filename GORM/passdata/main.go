package main

import (
	"passingdatawgorm/database"
	"passingdatawgorm/handlers"
	"passingdatawgorm/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/users",handlers.UsersPage)
	r.POST("/users",handlers.CreateUser)
	r.GET("/admin",handlers.AdminDashboard)

	r.Run(":8080")

}
