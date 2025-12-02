package handlers

import (
	"test/database"
	"test/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
	}

	database.DB.Create(user)
	c.JSON(201, user)
}

func Login(c *gin.Context) {
	var req struct {
		Name  string `json="name"`
		Email string `json="email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	user := &model.User{}

	err := database.DB.Where("email=?",req.Email).First(user).Error
	if err != nil{
		c.JSON(500,err)
		return
	}
	c.JSON(200,user)
}

func Home(c *gin.Context)  {
	var Users []model.User

	database.DB.Find(&Users)
	c.JSON(200,Users)
}