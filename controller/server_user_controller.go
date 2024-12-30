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

func AddUsersToServer(serverName model.ServerNameModel, users []AddUsersToServerRequest) error {

	userVars := _buildAddUsersToServerRequestMap(users)

	if userVars != nil {
		_, err := RunSimpleAnsibleTasks(
			serverName,
			helper.KFullPathTaskAddUsers,
			userVars,
			nil,
		)

		if err != nil {
			panic(err)
		}
	}

	userToGroupsVars := _buildAddUsersToGroupsRequestMap(users)

	if userToGroupsVars != nil {
		_, err := RunSimpleAnsibleTasks(
			serverName,
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
