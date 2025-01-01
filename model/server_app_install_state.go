package model

type ServerAppInstallState struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:VARCHAR(255);unique;not null"`
	DisplayName string `gorm:"not null"`
}

var serverAppInstallStateInstalledData = ServerAppInstallState{
	Name:        "installed",
	DisplayName: "Installed",
}

var serverAppInstallStateUninstalledData = ServerAppInstallState{
	Name:        "not_installed",
	DisplayName: "Not Installed",
}

func (ServerAppInstallState) GetServerAppInstalledStateData() ServerAppInstallState {
	return serverAppInstallStateInstalledData
}

func (ServerAppInstallState) GetServerAppUninstalledStateData() ServerAppInstallState {
	return serverAppInstallStateUninstalledData
}

var serverAppInstallStates = []ServerAppInstallState{
	serverAppInstallStateInstalledData,
	serverAppInstallStateUninstalledData,
}

func (ServerAppInstallState) Initialize() {
	GetDB().AutoMigrate(&ServerAppInstallState{})
	for _, state := range serverAppInstallStates {
		GetDB().Where(&ServerAppInstallState{Name: state.Name}).FirstOrCreate(&state)
	}
}
