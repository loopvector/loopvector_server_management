package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

type CreateServerRequest struct {
	DisplayName     string
	IPv4            *string
	IPv4Subnet      *uint64
	IPv6            *string
	IPv6Subnet      *uint64
	RootPassword    *string
	AdminUsername   *string
	AdminPassword   *string
	RootUserSSHKey  *string
	AdminUserSSHKey *string
}

func CreateNewServer(createServerRequest CreateServerRequest) {
	server := model.Server{
		ServerName:            helper.GenerateRandomUniqueServerName(),
		DisplayName:           createServerRequest.DisplayName,
		ServerActiveStateName: model.ServerActiveState{}.GetServerActiveStateData().Name,
	}
	server_id, err := server.CreateNew()
	if err != nil {
		panic(err)
	}

	if createServerRequest.IPv4 != nil && *createServerRequest.IPv4 != "" {
		serverIpv4 := model.ServerIpv4{
			ServerID:                server_id,
			Ip:                      *createServerRequest.IPv4,
			IpSubnet:                createServerRequest.IPv4Subnet,
			ServerIpActiveStateName: model.ServerIpActiveState{}.GetServerIpActiveStateData().Name,
		}
		err = serverIpv4.CreateNew()
		if err != nil {
			panic(err)
		}
	}

	if createServerRequest.IPv6 != nil && *createServerRequest.IPv6 != "" {
		serverIpv6 := model.ServerIpv6{
			ServerID:                server_id,
			Ip:                      *createServerRequest.IPv6,
			IpSubnet:                createServerRequest.IPv6Subnet,
			ServerIpActiveStateName: model.ServerIpActiveState{}.GetServerIpActiveStateData().Name,
		}
		err = serverIpv6.CreateNew()
		if err != nil {
			panic(err)
		}
	}

	if createServerRequest.RootPassword != nil && *createServerRequest.RootPassword != "" {
		rootUser := model.ServerRootUser{
			ServerID:                  server_id,
			Password:                  *createServerRequest.RootPassword,
			SshKey:                    createServerRequest.RootUserSSHKey,
			ServerUserActiveStateName: model.ServerUserActiveState{}.GetServerUserActiveStateData().Name,
		}
		err = rootUser.CreateNew()
		if err != nil {
			panic(err)
		}
	}

	if createServerRequest.AdminUsername != nil && *createServerRequest.AdminUsername != "" &&
		createServerRequest.AdminPassword != nil && *createServerRequest.AdminPassword != "" {
		adminUser := model.ServerUser{
			ServerID:                  server_id,
			Username:                  *createServerRequest.AdminUsername,
			Password:                  *createServerRequest.RootPassword,
			SshKey:                    createServerRequest.RootUserSSHKey,
			ServerUserActiveStateName: model.ServerUserActiveState{}.GetServerUserActiveStateData().Name,
		}
		err = adminUser.CreateNew()
		if err != nil {
			panic(err)
		}
	}
}

func DeleteServer(serverName string) {
	serverToDelete := model.Server{ServerName: serverName}
	err := serverToDelete.Delete()
	if err != nil {
		panic(err)
	}
}

func GetAllActiveServerNames() []string {
	return nil
}

func GetAllActiveServerNamesWithoutError() []string {
	servers, err := model.Server{
		ServerActiveStateName: model.ServerActiveState{}.GetServerActiveStateData().Name,
	}.GetAllActive()
	if err != nil {
		return []string{}
	}
	var serverNames []string
	for _, server := range servers {
		println("Server Name: ", server.ServerName)
		serverNames = append(serverNames, server.ServerName)
	}
	return serverNames
}

func ListAllServers() error {
	activeServer := model.Server{ServerActiveStateName: model.ServerActiveState{}.GetServerActiveStateData().Name}
	servers, err := activeServer.GetAllActive()
	if len(servers) == 0 {
		println("No active servers found")
	} else {
		for _, server := range servers {
			println("Server Name: ", server.ServerName)
			println("Server Display Name: ", server.DisplayName)
		}
	}
	return err
}
