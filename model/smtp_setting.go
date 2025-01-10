package model

import "time"

type SmtpSetting struct {
	ID           uint64 `gorm:"primaryKey"`
	SMTPHost     string `gorm:"not null"` // SMTP host
	SMTPPort     int    `gorm:"not null"` // SMTP port
	SMTPUser     string `gorm:"not null"` // SMTP username
	SMTPPassword string `gorm:"not null"` // SMTP password
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (SmtpSetting) Initialize() {
	GetDB().Debug().AutoMigrate(&SmtpSetting{})
}

func GetFirstSmtpSetting() (SmtpSetting, error) {
	var smtpSetting SmtpSetting
	if err := GetDB().First(&smtpSetting).Error; err != nil {
		return SmtpSetting{}, err
	}
	return smtpSetting, nil
}
