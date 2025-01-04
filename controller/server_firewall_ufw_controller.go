package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
	"strconv"
)

type UfwTrafficPolicy struct {
	Policy string
}

type UfwPortsTrafficPolicyRequest struct {
	Port     uint32
	Protocol string
}

type UfwIpAddressesTrafficPolicyRequest struct {
	Ip       string
	Protocol string
}

type UfwIpAddressesWithPortTrafficPolicyRequest struct {
	Port     uint32
	Ip       string
	Protocol string
}

func _buildUfwPortsTrafficPolicyRequestMap(trafficPolicy string, requests []UfwPortsTrafficPolicyRequest) map[string]interface{} {
	// Initialize the result map
	var ports []map[string]string

	// Iterate over the requests and build the user data
	for _, req := range requests {
		port := map[string]string{
			"port":     strconv.FormatUint(uint64(req.Port), 10),
			"protocol": req.Protocol,
		}
		ports = append(ports, port)
	}

	return map[string]interface{}{
		"ufw_ports":      ports,
		"traffic_policy": trafficPolicy,
	}
}

func _buildUfwIpAddressesTrafficPolicyRequestMap(trafficPolicy string, requests []UfwIpAddressesTrafficPolicyRequest) map[string]interface{} {
	// Initialize the result map
	var ips []map[string]string

	// Iterate over the requests and build the user data
	for _, req := range requests {
		port := map[string]string{
			"ip":       req.Ip,
			"protocol": req.Protocol,
		}
		ips = append(ips, port)
	}

	return map[string]interface{}{
		"ufw_ips":        ips,
		"traffic_policy": trafficPolicy,
	}
}

func _buildUfwIpAddressesWithPortTrafficPolicyRequestMap(trafficPolicy string, requests []UfwIpAddressesWithPortTrafficPolicyRequest) map[string]interface{} {
	// Initialize the result map
	var ipPortPairs []map[string]string

	// Iterate over the requests and build the user data
	for _, req := range requests {
		port := map[string]string{
			"port":     strconv.FormatUint(uint64(req.Port), 10),
			"ip":       req.Ip,
			"protocol": req.Protocol,
		}
		ipPortPairs = append(ipPortPairs, port)
	}

	return map[string]interface{}{
		"ufw_ip_port_pairs": ipPortPairs,
		"traffic_policy":    trafficPolicy,
	}
}

const (
	UfwTrafficPolicyAllow = "allow"
	UfwTrafficPolicyDeny  = "deny"
)

func SetDefaultIncomingUfwTrafficPolicy(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	trafficPolicy UfwTrafficPolicy,
) error {
	vars := map[string]interface{}{
		"traffic_policy": trafficPolicy.Policy,
	}
	return _runAnsibleTask(
		serverName,
		serverSshConnectionInfo,
		vars,
		helper.KFullPathTaskUfwDefaultIncomingConfigure,
	)
}

// func DenyDefaultIncomingUfwTrafficPolicy(serverName model.ServerNameModel) error {
// 	vars := map[string]interface{}{
// 		"traffic_policy": UfwTrafficPolicyDeny,
// 	}
// 	return _runAnsibleTask(serverName, vars, helper.KFullPathTaskUfwDefaultIncomingConfigure)
// }

func SetDefaultOutgoingUfwTrafficPolicy(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	trafficPolicy UfwTrafficPolicy,
) error {
	vars := map[string]interface{}{
		"traffic_policy": trafficPolicy.Policy,
	}
	return _runAnsibleTask(
		serverName,
		serverSshConnectionInfo,
		vars,
		helper.KFullPathTaskUfwDefaultOutgoingConfigure,
	)
}

// func DenyDefaultOutgoingUfwTrafficPolicy(serverName model.ServerNameModel) error {
// 	vars := map[string]interface{}{
// 		"traffic_policy": kTrafficPolicyDeny,
// 	}
// 	return _runAnsibleTask(serverName, vars, helper.KFullPathTaskUfwDefaultOutgoingConfigure)
// }

func EnableUfw(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
) error {
	return _runAnsibleTask(
		serverName,
		serverSshConnectionInfo,
		nil,
		helper.KFullPathTaskUfwEnable,
	)
}

func DisableUfw(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
) error {
	return _runAnsibleTask(
		serverName,
		serverSshConnectionInfo,
		nil,
		helper.KFullPathTaskUfwDisable,
	)
}

func SetUfwPorts(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	ufwPorts []UfwPortsTrafficPolicyRequest,
	trafficPolicy UfwTrafficPolicy,
) error {
	vars := _buildUfwPortsTrafficPolicyRequestMap(trafficPolicy.Policy, ufwPorts)
	return _runAnsibleTask(
		serverName,
		serverSshConnectionInfo,
		vars,
		helper.KFullPathTaskUfwPortsConfigure,
	)
}

// func DenyUfwPorts(
// 	serverName model.ServerNameModel,
// 	ufwPorts []UfwPortsTrafficPolicyRequest,
// ) error {
// 	vars := _buildUfwPortsTrafficPolicyRequestMap(kTrafficPolicyDeny, ufwPorts)
// 	return _runAnsibleTask(serverName, vars, helper.KFullPathTaskUfwPortsConfigure)
// }

func SetUfwIpAddresses(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	ufwIpAddresses []UfwIpAddressesTrafficPolicyRequest,
	trafficPolicy UfwTrafficPolicy,
) error {
	vars := _buildUfwIpAddressesTrafficPolicyRequestMap(
		trafficPolicy.Policy,
		ufwIpAddresses,
	)
	return _runAnsibleTask(
		serverName,
		serverSshConnectionInfo,
		vars,
		helper.KFullPathTaskUfwIpAddressesConfigure,
	)
}

// func DenyUfwIpAddresses(
// 	serverName model.ServerNameModel,
// 	ufwIpAddresses []UfwIpAddressesTrafficPolicyRequest,
// ) error {
// 	vars := _buildUfwIpAddressesTrafficPolicyRequestMap(kTrafficPolicyDeny, ufwIpAddresses)
// 	return _runAnsibleTask(serverName, vars, helper.KFullPathTaskUfwIpAddressesConfigure)
// }

func SetUfwIpAddressesWithPort(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	ufwIpAddressesWithPort []UfwIpAddressesWithPortTrafficPolicyRequest,
	trafficPolicy UfwTrafficPolicy,
) error {
	vars := _buildUfwIpAddressesWithPortTrafficPolicyRequestMap(
		trafficPolicy.Policy,
		ufwIpAddressesWithPort,
	)
	return _runAnsibleTask(
		serverName,
		serverSshConnectionInfo,
		vars,
		helper.KFullPathTaskUfwIpAddressesWithPortConfigure,
	)
}

// func DenyUfwIpAddressesWithPort(
// 	serverName model.ServerNameModel,
// 	ufwIpAddressesWithPort []UfwIpAddressesWithPortTrafficPolicyRequest,
// ) error {
// 	vars := _buildUfwIpAddressesWithPortTrafficPolicyRequestMap(kTrafficPolicyDeny, ufwIpAddressesWithPort)
// 	return _runAnsibleTask(serverName, vars, helper.KFullPathTaskUfwIpAddressesWithPortConfigure)
// }

func _runAnsibleTask(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	vars map[string]interface{},
	taskFullPath string,
) error {
	_, err := RunAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		[]model.AnsibleTask{{
			FullPath: taskFullPath,
			Vars:     vars,
		}},
		nil,
	)

	if err != nil {
		return err
	}

	return nil

}
