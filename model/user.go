package model

import (
	"log"
	"time"
)

type User struct {
	ID              uint64  `gorm:"primaryKey;autoIncrement"`
	Email           string  `gorm:"unique;not null"` // Email for login
	Username        *string `gorm:"unique;default:null"`
	Password        string  `gorm:"not null"`      // Hashed password
	IsAdmin         bool    `gorm:"default:false"` // Admin flag
	IsEmailVerified bool    `gorm:"default:false"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (User) Initialize() {
	GetDB().Debug().AutoMigrate(&User{})
}

func (u User) CreateNew() error {
	if err := GetDB().Where(&User{Email: u.Email}).FirstOrCreate(&u).Error; err != nil {
		log.Fatalf("failed to create user: %v", err)
		return err
	}
	return nil
}

func (u User) GetUsingEmailId() (User, error) {
	var user User
	if err := GetDB().Where(&User{Email: u.Email}).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r PasswordResetToken) GetUserUsingId() (User, error) {
	var user User
	if err := GetDB().Where(&User{ID: r.ID}).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (us UserSession) GetUserUsingId() (User, error) {
	var user User
	if err := GetDB().Where(&User{ID: us.UserID}).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (u User) UpdatePassword() error {
	return GetDB().Save(&u).Error
}
