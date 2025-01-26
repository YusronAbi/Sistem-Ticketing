package service

import (
	"SISTEM-TICKETING/config"
	"SISTEM-TICKETING/entity"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func RegisterUser(user *entity.User) (*entity.User, error) {
	db := config.DB

	// Cek apakah sudah ada user dengan email yang sama
	var existingUser entity.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("user with email %s already exists", user.Email)
	}

	// Simpan user baru ke database
	if err := db.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	return user, nil
}

// Fungsi untuk login dan autentikasi user
func AuthenticateUser(email string, password string) (*entity.User, error) {
	db := config.DB

	var user entity.User
	if err := db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found or invalid credentials")
		}
		log.Printf("Error authenticating user: %v", err)
		return nil, err
	}

	return &user, nil
}
