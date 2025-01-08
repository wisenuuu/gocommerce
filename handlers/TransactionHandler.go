package handlers

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler untuk mendapatkan transaksi beserta item terkait
func GetTransactionWithItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var transaction models.Transaction
		if err := db.Preload("Items").First(&transaction, id).Error; err != nil {
			c.JSON(http.StatusNotFound, models.Response{
				Success: "failed",
				Message: "Transaction not found",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Success: "success",
			Message: "Transaction retrieved successfully",
			Data:    transaction,
		})
	}
}

// Handler untuk membuat transaksi baru
func CreateTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Transaction

		// Validasi input JSON
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Success: "failed",
				Message: "Invalid input",
				Data:    nil,
			})
			return
		}

		// Validasi setiap item dalam transaksi
		for _, item := range input.Items {
			var product models.Product
			if err := db.First(&product, item.ProductID).Error; err != nil {
				c.JSON(http.StatusBadRequest, models.Response{
					Success: "failed",
					Message: "Invalid product ID",
					Data: map[string]interface{}{
						"product_id": item.ProductID,
					},
				})
				return
			}
		}

		// Buat transaksi baru
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Failed to create transaction",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, models.Response{
			Success: "success",
			Message: "Transaction created successfully",
			Data:    input,
		})
	}
}
