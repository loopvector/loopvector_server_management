package model

import "time"

type Sudoer struct {
	ID         uint64 `gorm:"primary_key;auto_increment"`
	ServerID   uint64 `gorm:"not null"`
	Server     Server `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Identifier string `gorm:"not null;unique"`
	CreatedAt  time.Time
}

func (Sudoer) Initialize() {
	GetDB().AutoMigrate(&Sudoer{})
}

func (a Sudoer) Create() error {
	if err := GetDB().Where(&Sudoer{
		ServerID:   a.ServerID,
		Identifier: a.Identifier,
	}).
		Attrs(&Sudoer{
			Identifier: a.Identifier,
		}).
		FirstOrCreate(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a Sudoer) DeleteUsingIdentifierAndServerId() error {
	if err := GetDB().Where(&Sudoer{
		ServerID:   a.ServerID,
		Identifier: a.Identifier,
	}).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}
