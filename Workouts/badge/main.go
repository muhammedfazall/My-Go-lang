package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Customer struct{
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Place string `gorm:"not null"`
}

var DB *gorm.DB

func main() {
	var err error
	dsn := "host=localhost user=postgres password=passsql dbname=demodb port=5432 sslmode=disable"

	DB,err = gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil{
		log.Fatalf("Error conneting database: %v",err)
	}
	
	if err := DB.AutoMigrate(&Customer{}) ; err != nil{
		log.Fatalf("Error migrating database: %v",err)
	}

	r := gin.Default()

	r.Run(":8080")
}

