package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Store users in memory
var users = map[string]string{} // username -> password

func main() {
	r := gin.Default()

	// REGISTER
	r.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "" || password == "" {
			c.String(http.StatusBadRequest, "username or password missing")
			return
		}

		if _, exists := users[username]; exists {
			c.String(http.StatusBadRequest, "user already exists")
			return
		}

		users[username] = password
		c.String(http.StatusOK, "registered successfully!")
	})

	// LOGIN
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "" || password == "" {
			c.String(http.StatusBadRequest, "username or password missing")
			return
		}

		if users[username] != password {
			c.String(http.StatusUnauthorized, "invalid username or password")
			return
		}

		// Save logged-in user in cookies
		c.SetCookie("user", username, 3600, "/", "localhost", false, true)

		c.String(http.StatusOK, "login successful!")
	})

	// HOME (shows who is logged in)
	r.GET("/home", func(c *gin.Context) {
		// Read cookie
		username, err := c.Cookie("user")

		if err != nil {
			c.String(http.StatusUnauthorized, "you are not logged in")
			return
		}

		c.String(http.StatusOK, "Welcome to Home, %s!", username)
	})

	// LOGOUT (optional)
	r.GET("/logout", func(c *gin.Context) {
		c.SetCookie("user", "", -1, "/", "localhost", false, true)
		c.String(http.StatusOK, "Logged out!")
	})

	r.Run(":8080")
}
