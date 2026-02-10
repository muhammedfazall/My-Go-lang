package main

import "github.com/gin-gonic/gin"


type Input struct{
	Num1 int `json:"n1"`
	Num2 int `json:"n2"`
}

func main() {

	r := gin.Default()

	r.POST("/add",handleAdd)

	r.Run(":8080")

}

func handleAdd(c *gin.Context){
	var req Input

	if err := c.ShouldBindJSON(&req) ; err != nil{
		c.JSON(400,gin.H{"error":"invalid"})
		return
	}

	sum := req.Num1 + req.Num2

	c.JSON(200,gin.H{"sum":sum})
}