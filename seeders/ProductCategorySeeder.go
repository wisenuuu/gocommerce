package seeders

import (
	"ecommerce/models"

	"gorm.io/gorm"
)

func ProductCategorySeeder(db *gorm.DB) {

	categories := []models.ProductCategory{
		{Name: "Category 1"},
		{Name: "Category 2"},
		{Name: "Category 3"},
		{Name: "Category 4"},
		{Name: "Category 5"},
	}

	var productCategoryExisting []models.ProductCategory
	db.Find(&productCategoryExisting)
	if len(productCategoryExisting) != len(categories) {
		db.CreateInBatches(&categories, 100)
	}
}
