package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

type UpdateRootUserPasswordRequest struct {
	NewRootPassword string
}

func UpdateRootUserPassword(serverName model.ServerNameModel, request UpdateRootUserPasswordRequest) error {

	vars := map[string]interface{}{
		"new_root_password": request.NewRootPassword,
		"unsafe":            true,
	}

	callback := RunAnsibleTaskCallback{
		TaskName: "update root password",
		OnChanged: func() {
			serverId := model.Server{ServerName: serverName.Name}.ServerActiveState.GetServerActiveStateIDUsingName()
			err := model.ServerRootUser{ServerID: serverId, Password: request.NewRootPassword}.UpdatePassword()
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
