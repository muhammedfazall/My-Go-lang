package database

import (
	model "gormbasic/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect()  {
	dsn := "host=localhost user=postgres password=passsql dbname=exer port=5432 sslmode=disable"

	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil{
		log.Fatalf("failed to connect:",err)
	}

	DB = db

	err = db.AutoMigrate(model.Customer{})
	if err != nil{
		log.Fatalf("failed to auto",err)

	}
	log.Println("successful")
}