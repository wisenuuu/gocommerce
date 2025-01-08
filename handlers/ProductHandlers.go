package handlers

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product

		// Mengambil data produk dari database
		if err := db.Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Failed to fetch products",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Success: "success",
			Message: "List of products",
			Data:    products,
		})
	}
}

func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product

		// Memeriksa apakah produk dengan ID tertentu ada
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, models.Response{
				Success: "failed",
				Message: "Product not found",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Success: "success",
			Message: "Product found",
			Data:    product,
		})
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Product

		// Validasi request body
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Success: "failed",
				Message: "Invalid request body",
				Data:    nil,
			})
			return
		}

		// Validasi apakah CategoryID ada
		var categoryExists bool
		if err := db.Model(&models.ProductCategory{}).
			Select("count(*) > 0").
			Where("id = ?", input.CategoryID).
			Find(&categoryExists).Error; err != nil || !categoryExists {
			c.JSON(http.StatusNotFound, models.Response{
				Success: "failed",
				Message: "Category not found",
				Data:    nil,
			})
			return
		}

		// Buat produk baru
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Failed to create product",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, models.Response{
			Success: "success",
			Message: "Product created successfully",
			Data:    input,
		})
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		id := c.Param("id")

		// Memeriksa apakah produk dengan ID tertentu ada
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, models.Response{
				Success: "failed",
				Message: "Product not found",
				Data:    nil,
			})
			return
		}

		var input models.Product
		// Validasi request body
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Success: "failed",
				Message: "Invalid input",
				Data:    nil,
			})
			return
		}

		// Validasi apakah CategoryID ada
		if input.CategoryID != 0 {
			if err := db.First(&models.ProductCategory{}, input.CategoryID).Error; err != nil {
				c.JSON(http.StatusNotFound, models.Response{
					Success: "failed",
					Message: "Category not found",
					Data:    nil,
				})
				return
			}
		}

		// Update produk
		if err := db.Model(&product).Updates(input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Failed to update product",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Success: "success",
			Message: "Product updated successfully",
			Data:    product,
		})
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		id := c.Param("id")

		// Memeriksa apakah produk dengan ID tertentu ada
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, models.Response{
				Success: "failed",
				Message: "Product not found",
				Data:    nil,
			})
			return
		}

		// Hapus produk
		if err := db.Delete(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Failed to delete product",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Success: "success",
			Message: "Product deleted successfully",
			Data:    product,
		})
	}
}
