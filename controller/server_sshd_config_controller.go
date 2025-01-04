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
	AddLineBlockToFile(
		serverName,
		serverSshConnectionInfo,
		addLinesToFileRequest,
		lines,
	)
	return nil
}
