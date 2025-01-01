package model

type ServerIpActiveState struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:VARCHAR(255);unique;not null"`
	DisplayName string `gorm:"unique;not null"`
}

var serverIpActiveStateActiveData = ServerIpActiveState{
	Name:        "active",
	DisplayName: "Active",
}

var serverIpActiveStateNotActiveData = ServerIpActiveState{
	Name:        "not_active",
	DisplayName: "Not Active",
}

func (ServerIpActiveState) GetServerIpActiveStateData() ServerIpActiveState {
	return serverIpActiveStateActiveData
}

func (ServerIpActiveState) GetServerIpNotActiveStateData() ServerIpActiveState {
	return serverIpActiveStateNotActiveData
}

var serverIpActiveStates = []ServerIpActiveState{
	serverIpActiveStateActiveData,
	serverIpActiveStateNotActiveData,
}

func (ServerIpActiveState) Initialize() {
	GetDB().AutoMigrate(&ServerIpActiveState{})
	for _, state := range serverIpActiveStates {
		GetDB().Where(&ServerIpActiveState{Name: state.Name}).FirstOrCreate(&state)
	}
}
