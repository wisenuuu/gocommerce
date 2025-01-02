package configs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password= dbname=gocommerce port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
