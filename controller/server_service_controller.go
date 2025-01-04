package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

type ServiceActionRequest struct {
	ServiceNames []string
}

func EnableServices(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	serviceNames ServiceActionRequest,
) error {
	vars := map[string]interface{}{
		"service_names": serviceNames.ServiceNames,
	}

	_, err := RunSimpleAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		helper.KFullPathTaskServiceEnable,
		vars,
		nil,
	)

	if err != nil {
		panic(err)
	}

	return nil
}

func StartServices(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	serviceNames ServiceActionRequest,
) error {
	vars := map[string]interface{}{
		"service_names": serviceNames.ServiceNames,
	}

	_, err := RunSimpleAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		helper.KFullPathTaskServiceStart,
		vars,
		nil,
	)

	if err != nil {
		panic(err)
	}

	return nil
}

func RestartServices(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	serviceNames ServiceActionRequest,
) error {
	vars := map[string]interface{}{
		"service_names": serviceNames.ServiceNames,
	}

	_, err := RunSimpleAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		helper.KFullPathTaskServiceRestart,
		vars,
		nil,
	)

	if err != nil {
		panic(err)
	}

	return nil
}
