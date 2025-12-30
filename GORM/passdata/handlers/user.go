package handlers

import (
	"passingdatawgorm/database"
	"passingdatawgorm/models"

	"github.com/gin-gonic/gin"
)

func UsersPage(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	c.HTML(200, "users.html", gin.H{
		"title": "Users",
		"users": users,
	})
}

func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	user := models.User{Name: name, Email: email}
	database.DB.Create(&user)

	c.Redirect(302, "/users")
}

func AdminDashboard(c *gin.Context) {
	c.HTML(200, "admin.html", gin.H{
		"title": "Admin Dashboard",
	})

}
