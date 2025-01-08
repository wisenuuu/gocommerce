package handlers

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User

		// Validasi input dari request body
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Success: "failed",
				Message: "Invalid input data",
				Data:    nil,
			})
			return
		}

		// Periksa apakah username sudah ada
		var existingUser models.User
		if err := db.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Success: "failed",
				Message: "Username already exists",
				Data:    nil,
			})
			return
		}

		// Periksa apakah email sudah ada
		if err := db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Success: "failed",
				Message: "Email already exists",
				Data:    nil,
			})
			return
		}

		// Hash password
		hashedPassword, err := HashPassword(input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Error while hashing password",
				Data:    nil,
			})
			return
		}

		// Buat user baru
		newUser := models.User{
			Username: input.Username,
			Email:    input.Email,
			Password: hashedPassword,
		}

		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Failed to create user",
				Data:    nil,
			})
			return
		}

		// Respon sukses
		c.JSON(http.StatusOK, models.Response{
			Success: "success",
			Message: "User registered successfully",
			Data: map[string]interface{}{
				"username": newUser.Username,
				"email":    newUser.Email,
			},
		})
	}
}

// Fungsi untuk hashing password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
