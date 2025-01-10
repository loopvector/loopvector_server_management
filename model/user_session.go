package model

import (
	"loopvector_server_management/controller/helper"
	"time"
)

type UserSession struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint64    `gorm:"not null"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Reference to User
	Token     string    `gorm:"unique;not null"`                               // Session token
	ExpiresAt time.Time `gorm:"not null"`                                      // Expiration time
	CreatedAt time.Time
}

func (UserSession) Initialize() {
	GetDB().Debug().AutoMigrate(&UserSession{})
}

func (user User) CreateNewUserSession() (string, error) {
	token := helper.GenerateToken()
	expiresAt := time.Now().Add(1 * time.Hour) // 1-hour session expiry
	session := UserSession{UserID: user.ID, Token: token, ExpiresAt: expiresAt}
	err := GetDB().Create(&session).Error
	if err != nil {
		return "", err
	}
	return token, nil
}

func (us UserSession) DeleteUserSession() error {
	err := GetDB().Where(&UserSession{Token: us.Token}).Delete(&UserSession{}).Error
	if err != nil {
		return err
	}
	return nil
}
