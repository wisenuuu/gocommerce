package seeders

import (
	"ecommerce/models"

	"gorm.io/gorm"
)

func ProductSeeder(db *gorm.DB) {
	
	products := []models.Product{
		{Name: "Product 1", Price: 10000, CategoryID: 1},
		{Name: "Product 2", Price: 20000, CategoryID: 3},
		{Name: "Product 3", Price: 30000, CategoryID: 2},
	}

	var productExistring []models.Product
	db.Find(&productExistring)
	if len(productExistring) != len(products) {
		db.CreateInBatches(&products, 100)
	}
}
