package model

type ServerActiveState struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:VARCHAR(255);unique;not null"`
	DisplayName string `gorm:"not null"`
}

var serverActiveStateActiveData = ServerActiveState{
	Name:        "active",
	DisplayName: "Active",
}

var serverActiveStateNotActiveData = ServerActiveState{
	Name:        "not_active",
	DisplayName: "Not Active",
}

func (ServerActiveState) GetServerActiveStateData() ServerActiveState {
	return serverActiveStateActiveData
}

func (ServerActiveState) GetServerNotActiveStateData() ServerActiveState {
	return serverActiveStateNotActiveData
}

var serverActiveStates = []ServerActiveState{
	serverActiveStateActiveData,
	serverActiveStateNotActiveData,
}

func (ServerActiveState) Initialize() {
	GetDB().AutoMigrate(&ServerActiveState{})
	for _, state := range serverActiveStates {
		GetDB().Where(&ServerActiveState{Name: state.Name}).FirstOrCreate(&state)
	}
}

func (sas ServerActiveState) GetServerActiveStateIDUsingName() uint64 {
	var result ServerActiveState
	println("sas.Name: ", sas.Name)
	if err := GetDB().Where(&ServerActiveState{Name: sas.Name}).Select("id").First(&result).Error; err != nil {
		println("Error fetching ID: ", err.Error()) // Log the error for debugging
		return 0
	}

	println("Fetched ID: ", result.ID)
	return result.ID
}
