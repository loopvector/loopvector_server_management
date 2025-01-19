package controller

import "loopvector_server_management/model"

type RunAnsibleTaskCallback struct {
	TaskNames   []string
	OnChanged   func()
	OnUnchanged func()
	OnFailed    func()
}

func RunSimpleAnsibleTasks(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	taskFullPath string,
	vars map[string]interface{},
	callback *RunAnsibleTaskCallback,
) (model.AnsiblePlaybookRunResult, error) {
	callbacks := []RunAnsibleTaskCallback{}
	if callback != nil {
		callbacks = append(callbacks, *callback)
	} else {
		callbacks = nil
	}
	return RunAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		[]model.AnsibleTask{{
			FullPath: taskFullPath,
			Vars:     vars,
		}},
		callbacks,
	)
}

func RunAnsibleTasks(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	ansibleTasks []model.AnsibleTask,
	callbacks []RunAnsibleTaskCallback,
) (model.AnsiblePlaybookRunResult, error) {
	// serverRootUser, serverIpv4, err := serverName.GetServerRootUserIpv4UsingServerName()
	// if err != nil {
	// 	panic(err)
	// }

	// err = model.AnsibleInventoryFileRootUserIpv4{
	// 	ServerName:     serverName.Name,
	// 	ServerIpv4:     serverIpv4,
	// 	ServerRootUser: serverRootUser,
	// }.CreateNewUsingRootUserAndIpv4()

	err := serverSshConnectionInfo.CreateNew()

	if err != nil {
		panic(err)
	}

	ansiblePlaybookRunner, err := model.AnsiblePlaybookFile{
		AnsibleTasks: ansibleTasks,
	}.CreateNew()

	if err != nil {
		panic(err)
	}

	result, err := ansiblePlaybookRunner.Run()

	if err != nil {
		panic(err)
	}

	for _, play := range result.Plays {
		for _, taskResult := range play.Tasks {
			for _, callback := range callbacks {
				// println("About to check install status of ", callback.TaskName, " on ", serverName.Name, " in task name ", taskResult.Task.Name)
				for _, taskName := range callback.TaskNames {
					if taskResult.Task.Name == taskName {
						if taskResult.Hosts["name="+serverName.Name].Failed == nil {
							if taskResult.Hosts["name="+serverName.Name].Changed {
								// println("changed ", callback.TaskName, " on ", serverName.Name)
								callback.OnChanged()
							} else {
								// println("unchanged ", callback.TaskName, " on ", serverName.Name)
								callback.OnUnchanged()
							}
						} else {
							// println("failed ", callback.TaskName, " on ", serverName.Name)
							callback.OnFailed()
						}
					}
				}
			}
		}
	}

	return result, nil
}
