package handlers

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProductCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var productCategories []models.ProductCategory

		// Ambil semua kategori produk dari database
		if err := db.Find(&productCategories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Failed to retrieve product categories",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Success: "success",
			Message: "List of product categories",
			Data:    productCategories,
		})
	}
}

func GetProductCategoriesById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var productCategory models.ProductCategory

		// Cari kategori produk berdasarkan ID
		if err := db.First(&productCategory, id).Error; err != nil {
			c.JSON(http.StatusNotFound, models.Response{
				Success: "failed",
				Message: "Product category not found",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Success: "success",
			Message: "Product category found",
			Data:    productCategory,
		})
	}
}

func CreateProductCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.ProductCategory

		// Validasi request body
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Success: "failed",
				Message: "Invalid request body",
				Data:    nil,
			})
			return
		}

		// Simpan kategori produk baru ke database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Failed to create product category",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, models.Response{
			Success: "success",
			Message: "Product category created successfully",
			Data:    input,
		})
	}
}

func UpdateProductCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var productCategory models.ProductCategory

		// Cari kategori produk berdasarkan ID
		if err := db.First(&productCategory, id).Error; err != nil {
			c.JSON(http.StatusNotFound, models.Response{
				Success: "failed",
				Message: "Product category not found",
				Data:    nil,
			})
			return
		}

		var input models.ProductCategory
		// Validasi request body
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				Success: "failed",
				Message: "Invalid input",
				Data:    nil,
			})
			return
		}

		// Update kategori produk
		if err := db.Model(&productCategory).Updates(input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Failed to update product category",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Success: "success",
			Message: "Product category updated successfully",
			Data:    productCategory,
		})
	}
}

func DeleteProductCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var productCategory models.ProductCategory

		// Cari kategori produk berdasarkan ID
		if err := db.First(&productCategory, id).Error; err != nil {
			c.JSON(http.StatusNotFound, models.Response{
				Success: "failed",
				Message: "Product category not found",
				Data:    nil,
			})
			return
		}

		// Hapus kategori produk
		if err := db.Delete(&productCategory).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: "failed",
				Message: "Failed to delete product category",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.Response{
			Success: "success",
			Message: "Product category deleted successfully",
			Data:    productCategory,
		})
	}
}
