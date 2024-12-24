package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

func AddGroupsToServer(serverName string, groupNames []string) {
	serverNameModel := model.ServerNameModel{Name: serverName}

	serverId, err := serverNameModel.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}

	vars := map[string]interface{}{
		"group_names": groupNames,
	}

	callbacks := []RunAnsibleTaskCallback{}

	for _, group := range groupNames {
		callbacks = append(callbacks, RunAnsibleTaskCallback{
			TaskName: "create group " + group,
			OnChanged: func() {
				println("Added group ", group, " to ", serverName)
				model.ServerIDModel{ID: serverId}.AddGroup(model.ServerGroup{ServerID: serverId, Name: group})
			},
			OnUnchanged: func() {
				println("Already added group ", group, " to ", serverName)
				model.ServerIDModel{ID: serverId}.AddGroup(model.ServerGroup{ServerID: serverId, Name: group})
			},
			OnFailed: func() {
				println("Failed to add group ", group, " to ", serverName)
			},
		})
	}

	_, err = RunAnsibleTasks(
		serverNameModel,
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
