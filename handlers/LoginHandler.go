package handlers

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, Response{
				Success: "failed",
				Message: "Invalid Input",
				Data:    nil,
			})
			return
		}

		var user models.User
		if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, Response{
				Success: "failed",
				Message: "Unauthorized",
				Data:    nil,
			})
			return
		}

		// Pastikan untuk memeriksa kata sandi yang benar di sini

		token, err := CreateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Success: "failed",
				Message: "Internal Server Error",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, Response{
			Success: "success",
			Message: "User Account Login successfully",
			Token:    token,
		})
	}
}
