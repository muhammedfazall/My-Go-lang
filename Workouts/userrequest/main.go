package main

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null" json:"password" binding:"required,min=6"`
	Role     string `gorm:"default:user"`
	Age      int    `json:"age"`
}

type UserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Product struct {
	ID    uint    `gorm:"primaryKey" json:"id"`
	Name  string  `gorm:"unique;not null" json:"name"`
	Price float64 `gorm:"not null" json:"price"`
}

var DB *gorm.DB

func main() {
	var err error
	dsn := "host=localhost user=postgres password=passsql dbname=demodb port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err := DB.AutoMigrate(&User{}, &Product{}); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// r.POST("/users", handleCreateUser)
	r.POST("/signup", handleSignup)
	r.POST("/login", handleLogin)
	r.POST("createproduct", CreateProduct)
	r.PATCH("updateproduct/:id", updateProduct)
	r.GET("/product/:id", getProduct)

	auth := r.Group("/auth")
	auth.Use(AuthMiddleware())

	auth.GET("/profile", func(c *gin.Context) {
		email, _ := c.Get("UserEmail")
		c.JSON(200, gin.H{"message": "success", "email": email})
	})

	r.Run(":8080")
}

func DeleteP(c *gin.Context){
	id := c.Param("id")

	result := DB.Delete(&Product{},id)
	if result.RowsAffected == 0 {

	}
}

func getProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product

	if err := DB.First(&product, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.JSON(200, gin.H{"product": product})
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product

	if err := DB.First(&product,id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "not found"})
		} else {
			c.JSON(500, gin.H{"error": "server error"})
		}
		return
	}

	var input Product

	if err := c.ShouldBindJSON(&input) ; err != nil{
		c.JSON(400, gin.H{"error": "invalid json format"})
		return
	}

	if err := DB.Model(&product).Updates(&input).Error ;err != nil{
		c.JSON(500, gin.H{"error": "could not update product"})
		return
	}

		c.JSON(200, gin.H{"product": product})


}

func CreateProduct(c *gin.Context) {
	var req Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	if err := DB.Create(&req).Error; err != nil {
		c.JSON(500, gin.H{"error": "could not create product"})
		return
	}

	c.JSON(200, gin.H{"status": "succesfull"})
}

func handleLogin(c *gin.Context) {
	var req UserRequest
	var user User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid json format"})
		return
	}

	if err := DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "invalid email or password!"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(401, gin.H{"error": "invalid email or password!"})
		return
	}

	accessTkn, err := GenerateToken(user.Email, user.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not create token"})
		return
	}

	c.JSON(200, gin.H{"token": accessTkn})
}

// func handleCreateUser(c *gin.Context) {
// 	var user User

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(400, gin.H{"error": "Validation failed: "})
// 		return
// 	}

// 	c.JSON(201, gin.H{"message": "User validation passed!", "data": user.Username})
// }

func handleSignup(c *gin.Context) {
	var req User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := User{
		Email:    req.Email,
		Password: string(hashed),
		Age:      req.Age,
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Registered Successfully",
		"email":   user.Email},
	)

}
