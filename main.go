package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "WELCOME TO GOLANGG BROOO",
		})
	})

	route.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	route.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Invalid Request Body",
			})
			return
		}

		if loginData.Email == "admin@mail.com" && loginData.Password == "password" {
			c.JSON(200, gin.H{
				"success": true,
				"message": "User logged in successfully",
				"data": gin.H{
					"email": loginData.Email,
				},
			})
		} else {
			c.JSON(401, gin.H{
				"success": false,
				"message": "Invalid Email or Password",
			})
		}
	})

	route.Run(":8080")
}
