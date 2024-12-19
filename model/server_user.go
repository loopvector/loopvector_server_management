package model

import "log"

type ServerUser struct {
	ID                        uint64 `gorm:"primaryKey;autoIncrement"`
	ServerID                  uint64 `gorm:"not null;index:uk_server_user_idx_server_id_username,unique"`
	Server                    Server `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Username                  string `gorm:"not null;index:uk_server_user_idx_server_id_username,unique"`
	Password                  string `gorm:"not null"`
	SshKey                    *string
	Port                      *uint16 `gorm:"default:22"`
	FirstName                 *string
	MiddleName                *string
	LastName                  *string
	Nickname                  *string
	Email                     *string
	PhoneNumber               *string
	ServerUserActiveStateName string                `gorm:"type:VARCHAR(255);not null"`
	ServerUserActiveState     ServerUserActiveState `gorm:"not null;references:Name;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

func (ServerUser) Initialize() {
	DB.AutoMigrate(&ServerUser{})
}

func (u *ServerUser) CreateNew() error {
	if err := DB.Create(&u).Error; err != nil {
		log.Fatalf("failed to create server: %v", err)
		return err
	}
	return nil
}
