package controller

import (
	"log"
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
	// log.Println("AddUsersToServer called")
	userVars := _buildAddUsersToServerRequestMap(users)
	serverId, err := serverName.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}

	// log.Println("AddUsersToServer userVars: ", userVars)

	if userVars != nil {

		callbacks := []RunAnsibleTaskCallback{}

		for _, user := range users {
			callbacks = append(
				callbacks,
				RunAnsibleTaskCallback{
					TaskNames: []string{"create user " + user.Username + " if it does not exist"},
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

	// log.Println("AddUsersToServer userToGroupsVars: ", userToGroupsVars)
	if userToGroupsVars != nil {
		callbacks := []RunAnsibleTaskCallback{}
		for _, user := range users {
			serverUser, err := model.ServerUser{
				ServerID: serverId,
				Username: user.Username,
			}.GetUsingServerIdAndUsername()
			if err != nil {
				panic(err)
			}
			for _, group := range user.Groups {
				serverGroup, err := model.ServerGroup{
					ServerID: serverId,
					Name:     group,
				}.GetUsingServerIdAndName()
				if err != nil {
					panic(err)
				}
				callbacks = append(
					callbacks, RunAnsibleTaskCallback{
						TaskNames: []string{"add user " + user.Username + " to group " + group},
						OnChanged: func() {
							// log.Println("ServerUserGroup callback OnChanged")
							model.ServerUserGroup{
								ServerID:      serverId,
								ServerGroupID: serverGroup.ID,
								ServerUserID:  serverUser.ID,
							}.Create()
						},
						OnUnchanged: func() {
							// log.Println("ServerUserGroup callback OnUnchanged")
							model.ServerUserGroup{
								ServerID:      serverId,
								ServerGroupID: serverGroup.ID,
								ServerUserID:  serverUser.ID,
							}.Create()
						},
						OnFailed: func() {
							log.Println("ServerUserGroup callback OnFailed")
						},
					})
			}
		}

		// callbacks := []RunAnsibleTaskCallback{
		// 	{
		// 		TaskNames: []string{"add users to groups"},
		// 		OnChanged: func() {
		// 			model.ServerUserGroup{
		// 				ServerID:      serverId,
		// 				ServerGroupID: serverGroup.ID,
		// 				ServerUserID:  serverUser.ID,
		// 			}.Create()
		// 		},
		// 		OnUnchanged: func() {
		// 			model.ServerUserGroup{
		// 				ServerID:      serverId,
		// 				ServerGroupID: serverGroup.ID,
		// 				ServerUserID:  serverUser.ID,
		// 			}.Create()
		// 		},
		// 		OnFailed: func() {},
		// 	},
		// }
		// log.Println("ServerUserGroup running ansible task, callback length: ", len(callbacks))
		_, err := RunAnsibleTasks(
			serverName,
			serverSshConnectionInfo,
			[]model.AnsibleTask{{
				FullPath: helper.KFullPathTaskAddUsersToGroups,
				Vars:     userToGroupsVars,
			}},
			callbacks,
		)

		if err != nil {
			panic(err)
		}
	}

	return nil
}
