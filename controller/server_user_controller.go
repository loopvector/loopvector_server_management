package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

type AddUsersToServerRequest struct {
	Username string
	Password string
	Groups   []string
}

func _buildAddUsersToServerRequestMap(requests []AddUsersToServerRequest) map[string]interface{} {
	if len(requests) == 0 {
		return nil
	}

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

func _buildAddUsersToGroupsRequestMap(requests []AddUsersToServerRequest) map[string]interface{} {
	if len(requests) == 0 {
		return nil
	}
	var usersAndGroups []map[string]interface{}
	for _, req := range requests {
		if len(req.Groups) > 0 {
			userAndGroups := map[string]interface{}{
				"user_name":   req.Username,
				"user_groups": req.Groups,
			}
			usersAndGroups = append(usersAndGroups, userAndGroups)
		}
	}

	if len(usersAndGroups) == 0 {
		return nil
	}

	return map[string]interface{}{
		"users_and_groups": usersAndGroups,
	}
}

func AddUsersToServer(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	users []AddUsersToServerRequest,
) error {

	userVars := _buildAddUsersToServerRequestMap(users)

	if userVars != nil {

		callbacks := []RunAnsibleTaskCallback{}

		serverId, err := serverName.GetServerIdUsingServerName()
		if err != nil {
			panic(err)
		}

		for _, user := range users {
			callbacks = append(
				callbacks,
				RunAnsibleTaskCallback{
					TaskName: "create user " + user.Username + " if it does not exist",
					OnChanged: func() {
						model.ServerUser{
							ServerID: serverId,
							Username: user.Username,
							Password: user.Password,
						}.CreateNewIfItDoesNotExist()
					},
					OnUnchanged: func() {
						model.ServerUser{
							ServerID: serverId,
							Username: user.Username,
							Password: user.Password,
						}.CreateNewIfItDoesNotExist()
					},
					OnFailed: func() {
						// println("Failed to install ", app, " on ", args[0])
					},
				},
			)
		}

		_, err = RunAnsibleTasks(
			serverName,
			serverSshConnectionInfo,
			[]model.AnsibleTask{{
				FullPath: helper.KFullPathTaskAddUsers,
				Vars:     userVars,
			}},
			callbacks,
		)

		if err != nil {
			panic(err)
		}
	}

	userToGroupsVars := _buildAddUsersToGroupsRequestMap(users)

	if userToGroupsVars != nil {
		_, err := RunSimpleAnsibleTasks(
			serverName,
			serverSshConnectionInfo,
			helper.KFullPathTaskAddUsersToGroups,
			userToGroupsVars,
			nil,
		)

		if err != nil {
			panic(err)
		}
	}

	return nil
}
