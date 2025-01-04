package model

import "log"

type ServerIpv4 struct {
	ID                      uint64              `gorm:"primaryKey;autoIncrement"`
	ServerID                uint64              `gorm:"not null;index:uk_server_ipv4_idx_server_id_ip_subnet,unique;"`
	Server                  Server              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ip                      string              `gorm:"not null;index:uk_server_ipv4_idx_server_id_ip_subnet,unique"`
	IpSubnet                *uint64             `gorm:"index:uk_server_ipv4_idx_server_id_ip_subnet,unique"`
	ServerIpActiveStateName string              `gorm:"type:VARCHAR(255);not null"`
	ServerIpActiveState     ServerIpActiveState `gorm:"not null;references:Name;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

func (ServerIpv4) Initialize() {
	GetDB().AutoMigrate(&ServerIpv4{})
}

func (ipv4 *ServerIpv4) CreateNew() error {
	if err := GetDB().Create(&ipv4).Error; err != nil {
		log.Fatalf("failed to create server: %v", err)
		return err
	}
	return nil
}

func (s *ServerIpv4) GetUsingServerId() (ServerIpv4, error) {
	var serverIpv4Details ServerIpv4
	if err := GetDB().Where(&ServerIpv4{ServerID: s.ServerID}).Find(&serverIpv4Details).Error; err != nil {
		return ServerIpv4{}, err
	}
	return serverIpv4Details, nil
}

func (s ServerIpv4) GetServerIpv4UsingIpAddress() (ServerIpv4, error) {
	var serverIpv4Details ServerIpv4
	if err := GetDB().Where(&ServerIpv4{Ip: s.Ip}).First(&serverIpv4Details).Error; err != nil {
		return ServerIpv4{}, err
	}
	return serverIpv4Details, nil
}
