package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

type AddUsersToServerRequest struct {
	Username string
	Password string
}

func _buildAddUsersToServerRequestMap(requests []AddUsersToServerRequest) map[string]interface{} {
	// Initialize the result map
	var users []map[string]string

	// Iterate over the requests and build the user data
	for _, req := range requests {
		userData := map[string]string{
			"user_name":     req.Username,
			"user_password": req.Password,
		}
		users = append(users, userData)
	}

	return map[string]interface{}{
		"users": users,
	}
}

func AddUsersToServer(serverName model.ServerNameModel, users []AddUsersToServerRequest) error {

	vars := _buildAddUsersToServerRequestMap(users)

	_, err := RunAnsibleTasks(
		serverName,
		[]model.AnsibleTask{{
			FullPath: helper.KFullPathTaskAddUsers,
			Vars:     vars,
		}},
		nil,
	)

	if err != nil {
		panic(err)
	}

	return nil
}
