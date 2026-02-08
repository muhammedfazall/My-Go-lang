package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	ID     uint    `gorm:"primaryKey"`
	Amount float64 `gorm:"not null"`
	UserID uint
}

type User struct {
	ID     uint    `gorm:"primaryKey"`
	Name   string  `json:"name" gorm:"not null"`
	Email  string  `json:"email" gorm:"unique;not null"`
	Age    int     `json:"age"`
	Orders []Order `json:"orders"`
}

var DB *gorm.DB

func createUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON format"})
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not save user!"})
		return
	}
	c.JSON(200, user)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := DB.Preload("Orders").First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(200, user)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	var input User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON format"})
		return
	}

	DB.Model(&user).Updates(input)
	c.JSON(200, user)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	result := DB.Delete(&User{}, id)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, gin.H{"message": "deleted successfullly"})
}

func main() {
	var err error
	dsn := "host=localhost user=postgres password=passsql dbname=users_db port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := DB.AutoMigrate(&User{}, &Order{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	r := gin.Default()

	r.POST("/user", createUser)
	r.GET("/user/:id", getUser)
	r.DELETE("/user/:id", deleteUser)
	r.PATCH("/user/:id", updateUser)

	r.Run(":8080")
}

// db.Model(&User{}).Where("name = ?","Ghopher").Update("Age",35)

// var user User
// result := db.Find(&user,"name = ?","Ghopher")
// if result.Error != nil{
// 	log.Printf("could not find user: %v", result.Error )
// } else {
// 	fmt.Printf("Age of user %s is %v",user.Name,user.Age)
// }

// newVIPUser := &User{
// 	Name:  "Ghopher",
// 	Email: "ghopherking@fake.com",
// 	Age:   30,
// 	Orders: []Order{
// 		{Amount: 150.99},
// 		{Amount: 22.00},
// 		{Amount: 88.99},
// 	},
// }

// err = db.Create(&newVIPUser).Error
// if err != nil {
// 	log.Printf("Error: %v", err)
// } else {
// 	log.Printf("Created User ID: %d with %d orders", newVIPUser.ID, len(newVIPUser.Orders))
// }

// if err := db.Create(&User{
// 	Name: "Fazal", Email: "faz@fake.com", Age: 24,
// }).Error; err != nil {
// 	log.Printf("unique constraint violation: %v", err)
// }

// var retrievedUser User

// result := db.Preload("Orders").First(&retrievedUser, "email = ?","ghopherking@fake.com")

// if result.Error != nil {
// 	log.Printf("could not find user: %v", result.Error )
// } else {
// 	fmt.Println(" order summery ")
// 	fmt.Printf("User : %v\n", retrievedUser.Name)

// 	var totalSpent float64
// 	for _,o := range retrievedUser.Orders{
// 		fmt.Printf("Order ID %v : $%.2f\n",o.ID,o.Amount)
// 		totalSpent += o.Amount
// 	}
// 	fmt.Printf("Total spent : $%.2f\n", totalSpent)
// }
