package controller

import (
	"fmt"
	"loopvector_server_management/model"
)

type SudoersAddRequest struct {
	GroupName  string
	Host       string
	RunAsUser  string
	RunAsGroup string
	Password   string
	Command    string
}

func _getLineFromRequest(sudoerEntry SudoersAddRequest) LineToFileAddRequest {
	passwordCommandPair := sudoerEntry.Command
	if sudoerEntry.Password != "" {
		passwordCommandPair = fmt.Sprintf("%s:%s", sudoerEntry.Password, sudoerEntry.Command)
	}

	// Form the runas part
	runAs := sudoerEntry.RunAsUser
	if sudoerEntry.RunAsGroup != "" {
		runAs = fmt.Sprintf("%s:%s", sudoerEntry.RunAsUser, sudoerEntry.RunAsGroup)
	}

	// Form the sudoers line
	line := fmt.Sprintf(
		"%%%s %s=(%s) %s",
		sudoerEntry.GroupName,
		sudoerEntry.Host,
		runAs,
		passwordCommandPair,
	)
	return LineToFileAddRequest{
		Line: line,
	}
}

func AddASudoer(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	addLinesToFileRequest AddLinesToFileRequest,
	sudoerEntry SudoersAddRequest,
) error {
	var lines []LineToFileAddRequest
	lines = append(
		lines,
		_getLineFromRequest(sudoerEntry),
	)
	AddLineBlockToFile(
		serverName,
		serverSshConnectionInfo,
		addLinesToFileRequest,
		lines,
	)
	return nil
}

func AddSudoers(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	addLinesToFileRequest AddLinesToFileRequest,
	sudoerEntries []SudoersAddRequest,
) error {
	var lines []LineToFileAddRequest

	for _, sudoerEntry := range sudoerEntries {
		lines = append(
			lines,
			_getLineFromRequest(sudoerEntry),
		)
	}
	AddLineBlockToFile(
		serverName,
		serverSshConnectionInfo,
		addLinesToFileRequest,
		lines,
	)
	return nil
}

// func AddASudoerLine(
// 	serverName string,
// 	addLinesToFileRequest AddLinesToFileRequest,
// 	sudoerLineEntry LineToFileAddRequest,
// ) error {
// 	AddLinesToFile(
// 		serverName,
// 		addLinesToFileRequest,
// 		[]LineToFileAddRequest{sudoerLineEntry},
// 	)
// 	return nil
// }
