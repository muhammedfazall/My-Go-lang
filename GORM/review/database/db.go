package database

import (
	"log"
	"review/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=passsql dbname=users_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to db:",err)
	}

	DB = db

	err = db.AutoMigrate(model.Customer{})
	if err != nil{
		log.Fatalf("FAiled to auto migrate:",err)
	}

	log.Println("Successful")
}