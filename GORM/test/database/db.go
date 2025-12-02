package database

import (
	"log"
	"test/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=passsql dbname=users_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database:%v", err)
	}

	DB = db

	err = db.AutoMigrate(model.User{})
	if err != nil {
		log.Fatalf("Failed to connect to Auto migrate:%v", err)
	}
	log.Println("Connected Sucessfully")
}
