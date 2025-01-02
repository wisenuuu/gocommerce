package migrations

import (
	"ecommerce/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Product{},
		&models.ProductCategory{},
		&models.Transaction{},
		&models.TransactionItem{},
		&models.User{},
	)
}
