package model

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _DB *gorm.DB

func GetDB() *gorm.DB {
	if _DB == nil {
		_connectToDatabase()
	}
	return _DB
}

func _connectToDatabase() {
	dbName := viper.GetString("DATABASE_NAME")
	dbUsername := viper.GetString("DATABASE_USERNAME")
	dbPassword := viper.GetString("DATABASE_PASSWORD")
	dbIpAddress := viper.GetString("DATABASE_IP_ADDRESS")
	dbPort := viper.GetString("DATABASE_PORT")
	// println("dbName: ", dbName)
	// println("dbUsername: ", dbUsername)
	// println("dbPassword: ", dbPassword)
	// println("dbIpAddress: ", dbIpAddress)
	// println("dbPort: ", dbPort)
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbIpAddress + ":" + dbPort + ")/" + dbName + "?parseTime=true"
	// log.Println("dsn: ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database err: " + err.Error())
	}
	_DB = db
}

func InitializeDB(shouldMigrate bool) {
	if shouldMigrate {
		ServerActiveState{}.Initialize()
		Server{}.Initialize()
		ServerIpActiveState{}.Initialize()
		ServerIpv4{}.Initialize()
		ServerIpv6{}.Initialize()
		ServerUserActiveState{}.Initialize()
		ServerRootUser{}.Initialize()
		ServerUser{}.Initialize()
		ServerAppInstallState{}.Initialize()
		ServerApp{}.Initialize()
		ServerGroup{}.Initialize()
		User{}.Initialize()
		UserSession{}.Initialize()
		PasswordResetToken{}.Initialize()
	}
}
