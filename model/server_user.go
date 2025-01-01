package model

type ServerUser struct {
	ID                        uint64 `gorm:"primaryKey;autoIncrement"`
	ServerID                  uint64 `gorm:"not null;index:uk_server_user_idx_server_id_username,unique"`
	Server                    Server `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Username                  string `gorm:"not null;index:uk_server_user_idx_server_id_username,unique"`
	Password                  string `gorm:"not null"`
	SshKey                    *string
	Port                      *uint16 `gorm:"default:22"`
	FirstName                 *string
	MiddleName                *string
	LastName                  *string
	Nickname                  *string
	Email                     *string
	PhoneNumber               *string
	ServerUserActiveStateName string                `gorm:"type:VARCHAR(255);not null"`
	ServerUserActiveState     ServerUserActiveState `gorm:"not null;references:Name;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

func (ServerUser) Initialize() {
	GetDB().AutoMigrate(&ServerUser{})
}

func (u *ServerUser) CreateNew() error {
	if err := GetDB().Where(&ServerUser{ServerID: u.ServerID, Username: u.Username}).
		Attrs(&ServerUser{ServerUserActiveStateName: ServerUserActiveState{}.GetServerUserActiveStateData().Name}).
		FirstOrCreate(&u).Error; err != nil {
		return err
	}
	return nil
}

// func (u *ServerUser) Remove() error {
// 	if err := GetDB().Where(&ServerUser{ServerID: u.ServerID, Username: u.Username}).
// 		Attrs(&ServerUser{ServerUserActiveStateName: ServerUserActiveState{}.GetServerUserNotActiveStateData().Name}).
// 		FirstOrCreate(&u).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
