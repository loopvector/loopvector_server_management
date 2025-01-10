package model

import (
	"loopvector_server_management/controller/helper"
	"time"
)

type PasswordResetToken struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	UserID    uint64    `gorm:"not null"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Foreign key to User
	Token     string    `gorm:"unique;not null"`                               // Reset token
	ExpiresAt time.Time `gorm:"not null"`                                      // Token expiration time
	CreatedAt time.Time
}

func (PasswordResetToken) Initialize() {
	GetDB().Debug().AutoMigrate(&PasswordResetToken{})
}

func (r PasswordResetToken) GetUsingToken() (PasswordResetToken, error) {
	var tokenData PasswordResetToken
	err := GetDB().Where(&PasswordResetToken{Token: r.Token}).First(&tokenData).Error
	if err != nil {
		return PasswordResetToken{}, err
	}
	return tokenData, nil
}

func (u User) CreateNewPasswordResetToken() (string, error) {
	token := helper.GenerateToken()
	expiresAt := time.Now().Add(1 * time.Hour) // 1-hour token expiry
	tokenData := PasswordResetToken{UserID: u.ID, Token: token, ExpiresAt: expiresAt}
	err := GetDB().Create(&tokenData).Error
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r PasswordResetToken) DeleteUsingToken() error {
	return GetDB().Where(&PasswordResetToken{Token: r.Token}).Delete(&PasswordResetToken{}).Error
}
