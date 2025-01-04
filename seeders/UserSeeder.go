package seeders

import (
	"ecommerce/models"

	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

func UserSeeder(db *gorm.DB) {

	// Data yang akan diinput
	users := []models.User{
		{Username: "admin", Email: "admin@mail.com", Password: HashPassword("password")},
		{Username: "user", Email: "user@mail.com", Password: HashPassword("password")},
	}

	//validasi apakah data sudah ada atau belum
	var userExisting []models.User
	db.Find(&userExisting)
	if len(userExisting) != len(users) {
		db.CreateInBatches(&users, 100)
	}
}

// Hashing Password
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
