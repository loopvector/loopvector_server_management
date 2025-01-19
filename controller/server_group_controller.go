package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

func AddGroupsToServer(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	groupNames []string,
) {
	serverId, err := serverName.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}

	vars := map[string]interface{}{
		"group_names": groupNames,
	}

	callbacks := []RunAnsibleTaskCallback{}

	for _, group := range groupNames {
		callbacks = append(callbacks, RunAnsibleTaskCallback{
			TaskNames: []string{"create group " + group},
			OnChanged: func() {
				// println("Added group ", group, " to ", serverName.Name)
				model.ServerIDModel{ID: serverId}.AddGroup(model.ServerGroup{ServerID: serverId, Name: group})
			},
			OnUnchanged: func() {
				// println("Already added group ", group, " to ", serverName.Name)
				model.ServerIDModel{ID: serverId}.AddGroup(model.ServerGroup{ServerID: serverId, Name: group})
			},
			OnFailed: func() {
				// println("Failed to add group ", group, " to ", serverName.Name)
			},
		})
	}

	_, err = RunAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		[]model.AnsibleTask{{
			FullPath: helper.KFullPathTaskAddGroups,
			Vars:     vars,
		}},
		callbacks,
	)

	if err != nil {
		panic(err)
	}
}
