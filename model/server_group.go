package model

type ServerGroup struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	ServerID uint64 `gorm:"not null;index:uk_server_group_idx_server_id_name,unique;"`
	Name     string `gorm:"not null;index:uk_server_group_idx_server_id_name,unique;"`
	Exists   bool   `gorm:"not null;"`
}

func (ServerGroup) Initialize() {
	GetDB().AutoMigrate(&ServerGroup{})
}

func (id ServerIDModel) AddGroup(serverGroup ServerGroup) error {
	if err := GetDB().Where(&ServerGroup{
		ServerID: id.ID,
		Name:     serverGroup.Name,
	}).Attrs(&ServerGroup{Exists: true}).
		FirstOrCreate(&serverGroup).
		Assign(&serverGroup).Error; err != nil {
		return err
	}
	return nil
}

func (u ServerGroup) GetUsingServerIdAndName() (ServerGroup, error) {
	var serverGroup ServerGroup
	if err := GetDB().Where(&ServerGroup{ServerID: u.ServerID, Name: u.Name}).First(&serverGroup).Error; err != nil {
		return ServerGroup{}, err
	}
	return serverGroup, nil
}

// func (sg ServerGroup) Remove() error {
// 	if err := GetDB().Where(&ServerGroup{
// 		ServerID: sg.ID,
// 		Name:     sg.Name,
// 	}).Attrs(&ServerGroup{Exists: false}).FirstOrCreate(&sg).Assign(&sg).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
