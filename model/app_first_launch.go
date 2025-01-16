package model

import (
	"log"
)

type AppFirstLaunch struct {
	Key   string `gorm:"primaryKey"`
	Value string
}

func (AppFirstLaunch) Initialize() {
	GetDB().AutoMigrate(&AppFirstLaunch{})
}

func (AppFirstLaunch) CheckFirstLaunch() bool {
	var count int64
	// Check if the 'AppFirstLaunch' table has a record for initialization
	if err := GetDB().Model(&AppFirstLaunch{}).Where(&AppFirstLaunch{Key: "initialized"}).Count(&count).Error; err != nil {
		log.Printf("Error checking initialization status: %v", err)
		// Assuming first launch in case of error
		return true
	}
	return count == 0
}

func (AppFirstLaunch) UpdateAppFirstLaunch() error {
	return GetDB().Exec("INSERT INTO app_first_launches (`key`, `value`) VALUES ('initialized', 'true') ON DUPLICATE KEY UPDATE `value`='true'").Error
}
