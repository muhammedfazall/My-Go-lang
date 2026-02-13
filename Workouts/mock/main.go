package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)


type Input struct{
	Num1 int `json:"n1"`
	Num2 int `json:"n2"`
}

func main() {

	r := gin.New()
	r.Use(Logger())

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

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()


		method := ctx.Request.Method
		path := ctx.FullPath()

		ctx.Next()

		status := ctx.Writer.Status()

		duration := time.Since(start)

		fmt.Printf("[%s] - %s - status:%v / in %v",method,path,status,duration)
	}
}