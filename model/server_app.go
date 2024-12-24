package model

type ServerApp struct {
	ID                        uint64 `gorm:"primaryKey;autoIncrement"`
	ServerID                  uint64 `gorm:"not null;index:uk_server_app_idx_server_id_name,unique;"`
	Server                    Server `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name                      string `gorm:"not_null;index:uk_server_app_idx_server_id_name,unique;"`
	Version                   string
	Vendor                    string
	ServerAppInstallStateName string                `gorm:"type:VARCHAR(255);not null"`
	ServerAppInstallState     ServerAppInstallState `gorm:"not null;references:Name;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

func (ServerApp) Initialize() {
	DB.AutoMigrate(&ServerApp{})
}

func (s ServerApp) RegisterInstall() error {
	if err := DB.Where(&ServerApp{
		ServerID: s.ID,
		Name:     s.Name,
	}).Attrs(&ServerApp{ServerAppInstallStateName: ServerAppInstallState{}.GetServerAppInstalledStateData().Name}).FirstOrCreate(&s).Error; err != nil {
		return err
	}
	return nil
}

func (s ServerApp) RegisterUninstall() error {
	if err := DB.Where(&ServerApp{
		ServerID: s.ID,
		Name:     s.Name,
	}).Attrs(&ServerApp{ServerAppInstallStateName: ServerAppInstallState{}.GetServerAppUninstalledStateData().Name}).
	FirstOrCreate(&s).
	Assign(&s).Error; err != nil {
		return err
	}
	return nil
}
