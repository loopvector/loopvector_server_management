package model

type ServerUserGroup struct {
	ID            uint64      `gorm:"primary_key;auto_increment"`
	ServerID      uint64      `gorm:"not null;index:uk_server_user_group_idx_server_id_user_id_group_id,unique"`
	Server        Server      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ServerUserID  uint64      `gorm:"not null;index:uk_server_user_group_idx_server_id_user_id_group_id,unique"`
	ServerUser    ServerUser  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ServerGroupID uint64      `gorm:"not null;index:uk_server_user_group_idx_server_id_user_id_group_id,unique"`
	ServerGroup   ServerGroup `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (ServerUserGroup) Initialize() {
	GetDB().AutoMigrate(&ServerUserGroup{})
}

func (a ServerUserGroup) Create() error {
	// log.Println(
	// 	"Creating ServerUserGroup, ServerID: ",
	// 	a.ServerID,
	// 	" ServerUserID: ",
	// 	a.ServerUserID,
	// 	" ServerGroupID: ",
	// 	a.ServerGroupID,
	// )
	if err := GetDB().Where(&ServerUserGroup{
		ServerID:      a.ServerID,
		ServerUserID:  a.ServerUserID,
		ServerGroupID: a.ServerGroupID,
	}).
		Attrs(&ServerUserGroup{
			ServerID:      a.ServerID,
			ServerUserID:  a.ServerUserID,
			ServerGroupID: a.ServerGroupID,
		}).
		FirstOrCreate(&a).Error; err != nil {
		return err
	}
	return nil
}
