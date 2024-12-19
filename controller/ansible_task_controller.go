package controller

import "loopvector_server_management/model"

func RunAnsibleTasks(serverName model.ServerNameModel, ansibleTasks []model.AnsibleTask) error {
	serverRootUser, serverIpv4, err := serverName.GetServerRootUserIpv4UsingServerName()
	if err != nil {
		panic(err)
	}

	err = model.AnsibleInventoryFileRootUserIpv4{
		ServerName:     serverName.Name,
		ServerIpv4:     serverIpv4,
		ServerRootUser: serverRootUser,
	}.CreateNewUsingRootUserAndIpv4()

	if err != nil {
		panic(err)
	}

	ansiblePlaybookRunner, err := model.AnsiblePlaybookFile{
		AnsibleTasks: ansibleTasks,
	}.CreateNew()

	if err != nil {
		panic(err)
	}

	ansiblePlaybookRunner.Run()
	return nil
}
