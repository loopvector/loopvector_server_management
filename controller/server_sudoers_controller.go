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

// func AddASudoer(
// 	serverName model.ServerNameModel,
// 	serverSshConnectionInfo model.ServerSshConnectionInfo,
// 	addLinesToFileRequest AddLinesToFileRequest,
// 	sudoerEntry SudoersAddRequest,
// ) error {
// 	var lines []LineToFileAddRequest
// 	lines = append(
// 		lines,
// 		_getLineFromRequest(sudoerEntry),
// 	)
// 	AddSudoerLines(
// 		serverName,
// 		serverSshConnectionInfo,
// 		addLinesToFileRequest,
// 		lines,
// 	)
// 	return nil
// }

func AddSudoerLines(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	addLinesToFileRequest AddLinesToFileRequest,
	lines []LineToFileAddRequest,
) error {
	serverId, err := serverName.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}
	callbacks := []RunAnsibleTaskCallback{{
		TaskNames: []string{"add line block to a file", "add line block to a file and set permissions"},
		OnChanged: func() {
			model.Sudoer{
				ServerID:   serverId,
				Identifier: addLinesToFileRequest.BlockTimestamp,
			}.Create()
		},
		OnUnchanged: func() {},
		OnFailed:    func() {},
	}}
	for _, line := range lines {
		AddLineBlockToFile(
			serverName,
			serverSshConnectionInfo,
			addLinesToFileRequest,
			[]LineToFileAddRequest{line},
			callbacks,
		)
	}
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
	AddSudoerLines(
		serverName,
		serverSshConnectionInfo,
		addLinesToFileRequest,
		lines,
	)
	// AddLineBlockToFile(
	// 	serverName,
	// 	serverSshConnectionInfo,
	// 	addLinesToFileRequest,
	// 	lines,
	// 	nil,
	// )
	return nil
}

func DeleteASudoerLine(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	deleteBlockFromFileRequest DeleteBlockFromFileRequest,
	identifier string,
) error {
	serverId, err := serverName.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}
	callbacks := []RunAnsibleTaskCallback{{
		TaskNames: []string{"delete line block from a file with block_timestamp: " + identifier},
		OnChanged: func() {
			model.Sudoer{
				ServerID:   serverId,
				Identifier: identifier,
			}.DeleteUsingIdentifierAndServerId()
		},
		OnUnchanged: func() {},
		OnFailed:    func() {},
	}}
	DeleteBlockFromFile(
		serverName,
		serverSshConnectionInfo,
		deleteBlockFromFileRequest,
		callbacks,
	)
	return nil
}

func ViewSudoerLines(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	readBlocksFromFileRequest ReadBlocksFromFileRequest,
) error {
	ReadBlocksFromFile(
		serverName,
		serverSshConnectionInfo,
		readBlocksFromFileRequest,
	)
	return nil
}

// func AddSudoerLineToSudoersFile(
// 	serverName model.ServerNameModel,
// 	serverSshConnectionInfo model.ServerSshConnectionInfo,
// 	addLinesToFileRequest AddLinesToFileRequest,
// 	lines []LineToFileAddRequest,
// ) {

// }

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
