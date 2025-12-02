package controllers

import (
	"go-user-management/database"
	"go-user-management/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



// ---------------- GET ALL USERS ----------------
func GetUsers(c *gin.Context) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(200, users)
}

// ---------------- CREATE USER ----------------
func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Create user
	if err := database.DB.Create(&newUser).Error; err != nil {
		log.Println("database.DB Error:", err)

		// if database.IsDuplicateEmailError(err) {
		// 	c.JSON(409, gin.H{"error": "Email already exists"})
		// 	return
		// }

		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(201, newUser)
}

// ---------------- GET USER BY ID ----------------
func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to fetch user"})
		return
	}

	c.JSON(200, user)
}

// ---------------- UPDATE USER ----------------
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var existingUser models.User
	var updatedUser models.User

	// Step 1: Check if user exists
	if err := database.DB.First(&existingUser, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to fetch user"})
		return
	}

	// Step 2: Read request body
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid update data"})
		return
	}

	// Step 3: Update
	if err := database.DB.Model(&existingUser).Updates(updatedUser).Error; err != nil {
		log.Println("Update Error:", err)
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	// Fetch updated data
	database.DB.First(&existingUser, id)

	c.JSON(200, existingUser)
}

// ---------------- DELETE USER ----------------
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to fetch user"})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
