package model

import "log"

type ServerIpv6 struct {
	ID                      uint64              `gorm:"primaryKey;autoIncrement"`
	ServerID                uint64              `gorm:"not null;index:uk_server_ipv6_idx_server_id_ip_subnet,unique;"`
	Server                  Server              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ip                      string              `gorm:"not null;index:uk_server_ipv6_idx_server_id_ip_subnet,unique"`
	IpSubnet                *uint64             `gorm:"index:uk_server_ipv6_idx_server_id_ip_subnet,unique"`
	ServerIpActiveStateName string              `gorm:"type:VARCHAR(255);not null"`
	ServerIpActiveState     ServerIpActiveState `gorm:"not null;references:Name;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

func (ServerIpv6) Initialize() {
	GetDB().AutoMigrate(&ServerIpv6{})
}

func (ipv6 *ServerIpv6) CreateNew() error {
	if err := GetDB().Create(&ipv6).Error; err != nil {
		log.Fatalf("failed to create server: %v", err)
		return err
	}
	return nil
}

func (s *ServerIpv6) GetUsingServerId() (ServerIpv6, error) {
	var serverIpv6Details ServerIpv6
	if err := GetDB().Where(&ServerIpv6{ServerID: s.ServerID}).Find(&serverIpv6Details).Error; err != nil {
		return ServerIpv6{}, err
	}
	return serverIpv6Details, nil
}

func (s ServerIpv6) GetServerIpv6UsingIpAddress() (ServerIpv6, error) {
	var serverIpv6Details ServerIpv6
	if err := GetDB().Where(&ServerIpv6{Ip: s.Ip}).First(&serverIpv6Details).Error; err != nil {
		return ServerIpv6{}, err
	}
	return serverIpv6Details, nil
}
