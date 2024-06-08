package seeder

import (
	"bwastartup/config"
	"bwastartup/models"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

// go run db/seeder.go SeedUser
func SeedUser() {
	tx := config.DB.Begin()
	user := []models.User{
		{
			Name:           "nur",
			Occupation:     "Programmer",
			Email:          "nurarifin@gmail.com",
			Password:       "password",
			Role:           "user",
			AvatarFileName: "nur.jpg",
			Token:          "awd",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			Name:           "arifin",
			Occupation:     "Programmer",
			Email:          "arifingdr@gmail.com",
			Password:       "password",
			Role:           "user",
			AvatarFileName: "arifin.jpg",
			Token:          "awd",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}

	if err := SaveRolesToDatabase(&user, tx); err != nil {
		log.Fatal("Failed to save user to database:", err)
	}
	tx.Commit()
}

// SaveUserToDatabase function untuk menyimpan data pengguna ke dalam database (contoh)
func SaveRolesToDatabase(user *[]models.User, tx *gorm.DB) error {
	// Simulasikan penyimpanan data ke dalam database
	// Di sini Anda bisa menggunakan ORM atau database driver yang digunakan dalam proyek Anda
	// Contoh sederhana: hanya mencetak informasi pengguna
	if err := config.DB.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Saved user to database: %v\n", user)
	return nil
}
