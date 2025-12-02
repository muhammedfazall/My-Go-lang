package main

import (
	"review/database"
	"review/model"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	r:= gin.Default()

	r.GET("/customers",getCustomers)

	r.Run(":8080")
}

func getCustomers(c *gin.Context)  {
	var Customers []model.Customer

	database.DB.Find(&Customers)
	c.JSON(200,Customers)
}