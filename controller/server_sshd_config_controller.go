package controller

import (
	"fmt"
	"loopvector_server_management/model"
)

type SSHDConfigAddRequest struct {
	Key            string
	Value          string
	MatchDirective string
}

func _generateSshdConfigLinesFromRequest(config SSHDConfigAddRequest) []LineToFileAddRequest {
	var lines []LineToFileAddRequest

	if config.MatchDirective != "" {
		lines = append(
			lines,
			LineToFileAddRequest{
				Line: fmt.Sprintf("Match %s", config.MatchDirective),
			},
		)
		lines = append(
			lines,
			LineToFileAddRequest{
				Line: fmt.Sprintf("    %s %s", config.Key, config.Value),
			},
		)
	} else {
		lines = append(
			lines,
			LineToFileAddRequest{
				Line: fmt.Sprintf("%s %s", config.Key, config.Value),
			},
		)
	}

	return lines
}

func DeleteASshdConfig(
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
			model.SshdConfig{
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

func ViewSshdConfigs(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	readBlocksFromFileRequest ReadBlocksFromFileRequest,
) {
	ReadBlocksFromFile(
		serverName,
		serverSshConnectionInfo,
		readBlocksFromFileRequest)
}

func AddASshdConfig(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	addLinesToFileRequest AddLinesToFileRequest,
	addSshdConfigRequest SSHDConfigAddRequest,
) error {
	lines := _generateSshdConfigLinesFromRequest(addSshdConfigRequest)
	// lines = append(lines, LineToFileAddRequest{Line: "\n"})
	// var block string
	// for _, line := range lines {
	// 	block = block + line.Line
	// }
	serverId, err := serverName.GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}
	// _, blockTimestamp := helper.GetCurrentTimestampMillis()
	callbacks := []RunAnsibleTaskCallback{{
		TaskNames: []string{"add line block to a file", "add line block to a file and set permissions"},
		OnChanged: func() {
			model.SshdConfig{
				ServerID:   serverId,
				Identifier: addLinesToFileRequest.BlockTimestamp,
			}.Create()
		},
		OnUnchanged: func() {},
		OnFailed:    func() {},
	}}
	AddLineBlockToFile(
		serverName,
		serverSshConnectionInfo,
		addLinesToFileRequest,
		lines,
		callbacks,
	)
	return nil
}
