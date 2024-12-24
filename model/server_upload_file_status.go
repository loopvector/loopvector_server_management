package model

type ServerUploadFileStatus struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:VARCHAR(255);unique;not null"`
	DisplayName string `gorm:"unique;not null"`
}
