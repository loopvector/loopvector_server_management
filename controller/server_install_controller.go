package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

func InstallServerApps(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	appsToInstall []string,
) error {
	vars := map[string]interface{}{
		"package_names": appsToInstall,
	}

	callbacks := []RunAnsibleTaskCallback{}

	serverId, err := serverName.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}

	for _, app := range appsToInstall {
		callbacks = append(
			callbacks,
			RunAnsibleTaskCallback{
				TaskNames: []string{"install package " + app},
				OnChanged: func() {
					// println("Installed ", app, " on ", args[0])
					model.ServerApp{
						ServerID: serverId,
						Name:     app,
					}.RegisterInstall()
				},
				OnUnchanged: func() {
					// println("Already installed ", app, " on ", args[0])
					model.ServerApp{
						ServerID: serverId,
						Name:     app,
					}.RegisterInstall()
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
		[]model.AnsibleTask{{FullPath: helper.KFullPathTaskInstallApps, Vars: vars}},
		callbacks,
	)

	if err != nil {
		panic(err)
	}

	return nil
}
