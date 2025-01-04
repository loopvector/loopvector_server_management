package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

type UpdateRootUserPasswordRequest struct {
	NewRootPassword string
}

func UpdateRootUserPassword(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	request UpdateRootUserPasswordRequest,
) error {

	vars := map[string]interface{}{
		"new_root_password": request.NewRootPassword,
		"unsafe":            true,
	}

	//println("update root password for serverName: ", serverName.Name)

	callback := RunAnsibleTaskCallback{
		TaskName: "update root password",
		OnChanged: func() {
			serverId, err := serverName.GetServerIdUsingServerName()
			if err != nil {
				panic(err)
			}
			//println("update root password for serverId: ", serverId)
			err = model.ServerRootUser{ServerID: serverId, Password: request.NewRootPassword}.UpdatePassword()
			if err != nil {
				panic(err)
			}
		},
		OnUnchanged: func() {

		},
		OnFailed: func() {

		},
	}

	_, err := RunSimpleAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		helper.KFullPathTaskChangeRootPassword,
		vars,
		// vars,
		&callback,
	)

	if err != nil {
		panic(err)
	}
	return nil
}
