package handlers

import (
	tokens "review2/pkg/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var UsersDb = map[string]User{}

func Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req) ; err != nil{
		c.JSON(400,gin.H{"error":"invalid json"})
		return
	}

	hashed,err := bcrypt.GenerateFromPassword([]byte(req.Password),bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500,gin.H{"error":"could not hash password"})
		return
	}

	UsersDb[req.Email] = User{
		Email: req.Email,
		Password: string(hashed),
		Role : "user",
	}

	c.JSON(201,gin.H{"message":"registered successfully"})
}

func Login(c * gin.Context)  {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req) ; err != nil{
		c.JSON(400,gin.H{"error":"invalid json"})
		return
	}

	user,ok := UsersDb[req.Email]
	if !ok {
		c.JSON(400,gin.H{"error":"user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password)); err != nil{
		c.JSON(400,gin.H{"error":"invalid password"})
		return
	}

	accessToken,refresToken,err := tokens.GenerateToken(user.Email,user.Role)
	if err != nil{
		c.JSON(500,gin.H{"error":"could not create tokens"})
		return
	}

	c.SetCookie("access",accessToken,600,"/","",false,true)
	c.SetCookie("refresh",refresToken,60 * 60,"/","",false,true)

	c.JSON(200,gin.H{"message":"Logged in succesfully"})
}

func Profile(c *gin.Context)  {
	claims := c.MustGet("claims").(*tokens.CustomClaims)

	c.JSON(200,gin.H{"message":"hey welcome "+ claims.Email}) 
}

func Logout(c *gin.Context)  {
	c.SetCookie("access","",-1,"/","",false,true)
	c.SetCookie("refresh","",-1,"/","",false,true)

	c.JSON(200,gin.H{"message":"Logged out successfully"})
}