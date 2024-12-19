package model

type ServerUserActiveState struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:VARCHAR(255);unique;not null"`
	DisplayName string `gorm:"not null"`
}

var serverUserActiveStateActiveData = ServerUserActiveState{
	Name:        "active",
	DisplayName: "Active",
}

var serverUserActiveStateNotActiveData = ServerUserActiveState{
	Name:        "not_active",
	DisplayName: "Not Active",
}

func (ServerUserActiveState) GetServerUserActiveStateData() ServerUserActiveState {
	return serverUserActiveStateActiveData
}

func (ServerUserActiveState) GetServerUserNotActiveStateData() ServerUserActiveState {
	return serverUserActiveStateNotActiveData
}

var serverUserActiveStates = []ServerUserActiveState{
	serverUserActiveStateActiveData,
	serverUserActiveStateNotActiveData,
}

func (ServerUserActiveState) Initialize() {
	DB.AutoMigrate(&ServerUserActiveState{})
	for _, state := range serverUserActiveStates {
		DB.Where(&ServerUserActiveState{Name: state.Name}).FirstOrCreate(&state)
	}
}
