package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

func PackageUpgrade(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
) error {
	_, err := RunAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		[]model.AnsibleTask{{FullPath: helper.KFullPathTaskPackageUpgrade}},
		nil,
	)
	return err
}

func PackageUpdate(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
) error {
	_, err := RunAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		[]model.AnsibleTask{{FullPath: helper.KFullPathTaskPackageUpdate}},
		nil,
	)
	return err
}

func PackageAutoRemove(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
) error {
	_, err := RunAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		[]model.AnsibleTask{{FullPath: helper.KFullPathTaskPackageAutoRemove}},
		nil,
	)
	return err
}
