package model

import "log"

type ServerIDModel struct {
	ID uint64
}

type ServerNameModel struct {
	Name string
}

type Server struct {
	ID                    uint64            `gorm:"primaryKey;autoIncrement"`
	ServerName            string            `gorm:"unique;not null"`
	DisplayName           string            `gorm:"unique;not null"`
	ServerActiveStateName string            `gorm:"type:VARCHAR(255);not null"`
	ServerActiveState     ServerActiveState `gorm:"not null;references:Name;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

func (Server) Initialize() {
	GetDB().Debug().AutoMigrate(&Server{})
}

func (s *Server) CreateNew() (uint64, error) {
	if err := GetDB().Create(&s).Error; err != nil {
		log.Fatalf("failed to create server: %v", err)
		return 0, err
	}
	return s.ID, nil
}

func (s *Server) Delete() error {
	if err := GetDB().Where(&Server{ServerName: s.ServerName}).Delete(&Server{}).Error; err != nil {
		log.Fatalf("failed to delete server: %v", err)
		return err
	}
	return nil
}

func (s Server) GetAllActive() ([]Server, error) {
	var servers []Server
	if err := GetDB().Where(&Server{ServerActiveStateName: s.ServerActiveStateName}).Find(&servers).Error; err != nil {
		return nil, err
	}
	return servers, nil
}

func (s ServerNameModel) GetServerIdUsingServerName() (uint64, error) {
	var result Server
	if err := GetDB().Where(&Server{ServerName: s.Name}).Select("id").First(&result).Error; err != nil {
		return 0, err
	}
	return result.ID, nil
}

func (s *ServerNameModel) GetServerRootUserIpv4UsingServerName() (ServerRootUser, ServerIpv4, error) {
	serverId, err := s.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}
	serverRootUserRequest := ServerRootUser{
		ServerID: serverId,
	}
	serverRootUser, err := serverRootUserRequest.GetUsingServerId()
	if err != nil {
		panic(err)
	}
	serverIpv4Request := ServerIpv4{
		ServerID: serverId,
	}
	serverIpv4, err := serverIpv4Request.GetUsingServerId()
	if err != nil {
		panic(err)
	}
	return serverRootUser, serverIpv4, nil
}

func (s *ServerNameModel) GetRootUserUsingServerName() (ServerRootUser, error) {
	serverId, err := s.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}
	var serverRootUser = ServerRootUser{
		ServerID: serverId,
	}
	result, err := serverRootUser.GetUsingServerId()
	if err != nil {
		panic(err)
	}
	return result, nil
}

func (s ServerNameModel) GetServerUserUsingServerName(username string) (ServerUser, error) {
	serverId, err := s.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}
	var serverUser = ServerUser{
		ServerID: serverId,
		Username: username,
	}
	result, err := serverUser.GetUsingServerIdAndUsername()
	if err != nil {
		panic(err)
	}
	return result, nil
}

func (s *ServerNameModel) GetIpv4UsingServerName() (ServerIpv4, error) {
	serverId, err := s.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}
	var serverIpv4 = ServerIpv4{
		ServerID: serverId,
	}
	result, err := serverIpv4.GetUsingServerId()
	if err != nil {
		panic(err)
	}
	return result, nil
}

func (s *ServerNameModel) GetIpv6UsingServerName() (ServerIpv6, error) {
	serverId, err := s.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}
	var serverIpv6 = ServerIpv6{
		ServerID: serverId,
	}
	result, err := serverIpv6.GetUsingServerId()
	if err != nil {
		panic(err)
	}
	return result, nil
}

func (s ServerNameModel) AddNewUser(user ServerUser) error {
	serverId, err := s.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}
	user.ServerID = serverId
	return user.CreateNewIfItDoesNotExist()
}

// func (s *ServerNameModel) GetAppUsingServerName() (ServerApp, error) {
// 	serverId, err := s.GetServerIdUsingServerName()
// 	if err != nil {
// 		panic(err)
// 	}
// 	var serverApp = ServerApp{
// 		ServerID: serverId,
// 	}
// 	result, err := serverApp.Get()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return result, nil
// }
