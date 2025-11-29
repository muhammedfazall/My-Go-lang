package main

import (
	"ecom_gorm/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	const DSN = "host=localhost user=postgres password=passsql dbname=ecom_db port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Successfully connected to ecom_db!")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}
	fmt.Println("User table schema verified/migrated successfully!")

	// --- 2. CREATE (Insert) ---
	fmt.Println("\n--- C: Creating User ---")
	newUser := models.User{
		FirstName: "Alice",
		Email:     "alice@example.com",
	}
	// db.Create(&newUser) inserts the user into the 'users' table
	db.Create(&newUser)
	fmt.Printf("Created User (ID: %d, Email: %s)\n", newUser.ID, newUser.Email)

	// --- 3. READ (Retrieve) ---
	fmt.Println("\n--- R: Reading User ---")
	var retrievedUser models.User

	// db.First retrieves the first record that matches the condition (ID=newUser.ID)
	db.First(&retrievedUser, newUser.ID)
	fmt.Printf("Retrieved User: ID %d, Name: %s\n", retrievedUser.ID, retrievedUser.FirstName)

	// --- 4. UPDATE (Modify) ---
	fmt.Println("\n--- U: Updating User ---")
	// Change the email address
	retrievedUser.Email = "alice.new@example.com"
	// db.Save() persists the change (updates the row where ID matches retrievedUser.ID)
	db.Save(&retrievedUser)
	fmt.Printf("Updated User (ID: %d) new Email: %s\n", retrievedUser.ID, retrievedUser.Email)

	// --- 5. DELETE (Remove) ---
	fmt.Println("\n--- D: Deleting User ---")
	// db.Delete() removes the record where ID matches retrievedUser.ID
	// NOTE: GORM uses soft deletes by default (setting DeletedAt), but we will verify removal in DB
	db.Delete(&retrievedUser)
	fmt.Printf("Deleted User (ID: %d)\n", retrievedUser.ID)

	// --- 6. Final Check ---
	fmt.Println("\n--- Final Check (Should show 0 results) ---")
	var deletedCheck models.User
	// db.First will try to find the record; if it fails, it returns a record not found error.
	result := db.First(&deletedCheck, retrievedUser.ID)
	if result.Error == gorm.ErrRecordNotFound {
		fmt.Println("Successfully verified: User record is no longer accessible.")
	} else {
		fmt.Println("Check failed or another error occurred.")
	}
}
