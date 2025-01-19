package model

import "time"

type SshdConfig struct {
	ID         uint64 `gorm:"primary_key;auto_increment"`
	ServerID   uint64 `gorm:"not null"`
	Server     Server `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Identifier string `gorm:"not null;unique"`
	CreatedAt  time.Time
}

func (SshdConfig) Initialize() {
	GetDB().AutoMigrate(&SshdConfig{})
}

func (a SshdConfig) Create() error {
	if err := GetDB().Where(&SshdConfig{
		ServerID: a.ServerID,
	}).
		Attrs(&SshdConfig{
			Identifier: a.Identifier,
		}).
		FirstOrCreate(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a SshdConfig) DeleteUsingIdentifierAndServerId() error {
	if err := GetDB().Where(&SshdConfig{
		ServerID:   a.ServerID,
		Identifier: a.Identifier,
	}).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}
