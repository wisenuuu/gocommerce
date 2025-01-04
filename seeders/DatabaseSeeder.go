package seeders

import (
	"gorm.io/gorm"
)

func DatabaseSeeder(db *gorm.DB) {
	ProductCategorySeeder(db)
	UserSeeder(db)
	ProductSeeder(db)
}
