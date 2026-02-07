package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	ID     uint    `gorm:"primaryKey"`
	Amount float64 `gorm:"not null"`
	UserID uint
}

type User struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"not null"`
	Email  string `gorm:"unique;not null"`
	Age    int
	Orders []Order
}

func main() {
	dsn := "host=localhost user=postgres password=passsql dbname=users_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&User{}, &Order{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	db.Model(&User{}).Where("name = ?","Ghopher").Update("Age",35)

	var user User
	result := db.Find(&user,"name = ?","Ghopher")
	if result.Error != nil{
		log.Printf("could not find user: %v", result.Error )
	} else {
		fmt.Printf("Age of user %s is %v",user.Name,user.Age)
	}

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
}
