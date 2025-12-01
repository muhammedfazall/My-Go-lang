package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm="primaryKey" json="id"`
	Name  string `json="name"`
	Email string `gorm="unique" json="email"`
	Age   int    `json="age"`
}

var DB *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=passsql dbname=users_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database:%v", err)
	}

	sqlDb,err := db.DB()

	if err != nil {
		log.Fatalf("Failed to get generic to the database object: %v",err)
	}

	sqlDb.SetMaxOpenConns(25)
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)

	DB = db

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("failed to migrate from database,%v", err)
	}

	router := gin.Default()
	router.GET("/users", getUsers)

	router.POST("/users", createUser)

	router.GET("/users/:id", getUserById)

	router.PUT("/users/:id", updateUser)

	router.DELETE("/users/:id", deleteUser)

	log.Println("Server running on :8080")
	router.Run(":8080")

}

func getUsers(c *gin.Context) {

	var users []User
	if result := DB.Find(&users); result.Error != nil {
		c.JSON(500, gin.H{"error": "failed to fetch users from database"})
		return
	}
	c.JSON(200, users)
}

func createUser(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "invalid request data"})
		return
	}

	if result := DB.Create(&newUser); result.Error != nil {
		c.JSON(500, gin.H{"error": "failed to create user."})
		return
	}
	c.JSON(201, newUser)
}

func getUserById(c *gin.Context) {
	id := c.Param("id")
	var user User

	if result := DB.First(&user, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to fetch user"})
		return
	}
	c.JSON(200, user)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	if result := DB.First(&user, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to fetch user"})
		return
	}

	if result := DB.Delete(&user, id); result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(201, nil)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var existingUser User
	var updatedUser User

	if result := DB.First(&existingUser, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User Not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to fetch user"})
		return
	}

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": "invalid update data"})
		return
	}

	if result := DB.Model(&existingUser).Updates(&updatedUser); result.Error != nil {
		log.Printf("DB update Error:%v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	DB.First(&existingUser,id)
	c.JSON(200,existingUser)
}
