package model

type ServerGroup struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	ServerID uint64 `gorm:"not null;index:uk_server_group_idx_server_id_name,unique;"`
	Name     string `gorm:"not null;index:uk_server_group_idx_server_id_name,unique;"`
	Exists   bool   `gorm:"not null;"`
}

func (ServerGroup) Initialize() {
	DB.AutoMigrate(&ServerGroup{})
}

func (id ServerIDModel) AddGroup(serverGroup ServerGroup) error {
	if err := DB.Where(&ServerGroup{
		ServerID: id.ID,
		Name:     serverGroup.Name,
	}).Attrs(&ServerGroup{Exists: true}).
		FirstOrCreate(&serverGroup).
		Assign(&serverGroup).Error; err != nil {
		return err
	}
	return nil
}

// func (sg ServerGroup) Remove() error {
// 	if err := DB.Where(&ServerGroup{
// 		ServerID: sg.ID,
// 		Name:     sg.Name,
// 	}).Attrs(&ServerGroup{Exists: false}).FirstOrCreate(&sg).Assign(&sg).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
