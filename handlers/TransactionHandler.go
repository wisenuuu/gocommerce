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

func CreateTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil userID dari context
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, models.Response{
				Success: "failed",
				Message: "User not authenticated",
			})
			return
		}

		// Bind input JSON
		var input struct {
			Items []struct {
				ProductID uint `json:"product_id" binding:"required"`
				Quantity  uint `json:"quantity" binding:"required"`
			} `json:"items" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Success: "failed",
				Message: "Invalid input",
				Data:    nil,
			})
			return
		}

		// Hitung total amount
		var totalAmount float64
		var transactionItems []models.TransactionItem //deklrasi untuk menyimpan mengelola yang akan dimasukan kedalam database

		for _, item := range input.Items {
			var product models.Product //deklarasi untuk proses pengambalian data dari product
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

			// Hitung harga berdasarkan jumlah dan harga produk
			itemTotal := float64(item.Quantity) * product.Price
			totalAmount += itemTotal

			// Tambahkan item ke daftar transactionItems
			transactionItems = append(transactionItems, models.TransactionItem{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     product.Price, // Harga produk dari database
			})
		}

		// Buat transaksi
		transaction := models.Transaction{
			UserID: userID.(uint),
			Amount: totalAmount,
		}
		if err := db.Create(&transaction).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Failed to create transaction",
				Data:    nil,
			})
			return
		}

		// Tambahkan items ke transaksi
		for i := range transactionItems {
			transactionItems[i].TransactionID = transaction.ID
		}
		if err := db.Create(&transactionItems).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Failed to create transaction items",
				Data:    nil,
			})
			return
		}

		// Sertakan items ke dalam respons transaction
		transaction.Items = transactionItems

		// Berikan respons dengan data transaksi lengkap
		c.JSON(http.StatusCreated, models.Response{
			Success: "success",
			Message: "Transaction created successfully",
			Data: map[string]interface{}{
				"transaction": transaction,
			},
		})
	}
}
