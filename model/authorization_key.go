package model

import (
	"time"
)

type AuthorizationKey struct {
	ID           uint64     `gorm:"primary_key;auto_increment"`
	ServerID     uint64     `gorm:"not null;index:uk_authorization_key_idx_server_id_server_user_id_public_key,unique"`
	Server       Server     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ServerUserID uint64     `gorm:"not null;index:uk_authorization_key_idx_server_id_server_user_id_public_key,unique"`
	ServerUser   ServerUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PublicKey    string     `gorm:"type:TEXT;not null;index:uk_authorization_key_idx_server_id_server_user_id_public_key,unique,length:255" json:"public_key"`
	Identifier   string     `gorm:"not null;unique"`
	CreatedAt    time.Time
}

func (AuthorizationKey) Initialize() {
	GetDB().AutoMigrate(&AuthorizationKey{})
}

func (a AuthorizationKey) Create() error {
	if err := GetDB().Where(&AuthorizationKey{
		ServerID:     a.ServerID,
		ServerUserID: a.ServerUserID,
		PublicKey:    a.PublicKey,
	}).
		Attrs(&AuthorizationKey{
			PublicKey:  a.PublicKey,
			Identifier: a.Identifier,
		}).
		FirstOrCreate(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a AuthorizationKey) GetUsingServerIDAndServerUserID() ([]AuthorizationKey, error) {
	var authorizationKeys []AuthorizationKey
	if err := GetDB().Where(&AuthorizationKey{
		ServerID:     a.ServerID,
		ServerUserID: a.ServerUserID,
	}).Find(&authorizationKeys).Error; err != nil {
		return nil, err
	}
	return authorizationKeys, nil
}

func (a AuthorizationKey) DeleteUsingIdentifierUserIdAndServerId() error {
	if err := GetDB().Where(&AuthorizationKey{
		ServerID:     a.ServerID,
		ServerUserID: a.ServerUserID,
		Identifier:   a.Identifier,
	}).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}
