package handler

import (
	"exercises/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var req User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password needed"})
		return
	}

	if database.UserExists(req.Username) {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists!"})
		return
	}

	if err := database.CreateUser(req.Username, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered!"})
}

func Login(c *gin.Context) {
	var r User

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
	}

	if r.Username == "" || r.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password needed"})
		return
	}

	if !database.VerifyUser(r.Username, r.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
		return
	}

	id, err := database.CreateSession(r.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gftf"})
		return
	}

	c.SetCookie(database.Sessionid, id, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "Logged in"})
}


func Logout(c *gin.Context)  {

	sid,err := c.Cookie(database.Sessionid)
	if err == nil {
		database.DeleteSession(sid)
	}
	c.SetCookie(database.Sessionid,"",-1,"/","",false,true)
	c.JSON(http.StatusOK,gin.H{"status":"logged out"})
}