package model

import "time"

type SignUpDomainWhitelistSetting struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Domain    string `gorm:"type:text;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SignUpDomainWhitelistSetting) Initialize() {
	GetDB().Debug().AutoMigrate(&SignUpDomainWhitelistSetting{})
}
