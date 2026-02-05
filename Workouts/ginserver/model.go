package main

type loginReq struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      int    `json:"age" binding:"required,gte=18"`

	// json:"name" tells Gin to look for "name" in the JSON, not "Name"
	// binding:"required" makes Gin return an error if it's missing
}