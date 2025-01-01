package model

import "log"

type ServerRootUser struct {
	ID                        uint64 `gorm:"primaryKey;autoIncrement"`
	ServerID                  uint64 `gorm:"not null;index:uk_server_root_user_idx_server_id_password,unique;"`
	Server                    Server `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Password                  string `gorm:"not null;index:uk_server_root_user_idx_server_id_password,unique;"`
	SshKey                    *string
	Port                      uint16                `gorm:"default:22"`
	ServerUserActiveStateName string                `gorm:"type:VARCHAR(255);not null"`
	ServerUserActiveState     ServerUserActiveState `gorm:"not null;references:Name;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

func (ServerRootUser) Initialize() {
	GetDB().AutoMigrate(&ServerRootUser{})
}

func (ru *ServerRootUser) CreateNew() error {
	if err := GetDB().Create(&ru).Error; err != nil {
		log.Fatalf("failed to create server: %v", err)
		return err
	}
	return nil
}

func (ru ServerRootUser) UpdatePassword() error {
	if err := GetDB().Where(&ServerRootUser{ServerID: ru.ServerID}).Updates(ServerRootUser{Password: ru.Password}).Error; err != nil {
		log.Fatalf("failed to update password: %v", err)
		return err
	}
	return nil
}

func (s *ServerRootUser) GetUsingServerId() (ServerRootUser, error) {
	var serverRootUser ServerRootUser
	if err := GetDB().Where(&ServerRootUser{ServerID: s.ServerID}).Find(&serverRootUser).Error; err != nil {
		return ServerRootUser{}, err
	}
	return serverRootUser, nil
}
