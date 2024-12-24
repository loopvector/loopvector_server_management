package controller

import (
	"fmt"
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
	serverName string,
	addLinesToFileRequest AddLinesToFileRequest,
	addSshdConfigRequest SSHDConfigAddRequest,
) error {
	lines := _generateSshdConfigLinesFromRequest(addSshdConfigRequest)
	lines = append(lines, LineToFileAddRequest{Line: "\n"})
	AddLinesToFile(
		serverName,
		addLinesToFileRequest,
		lines,
	)
	return nil
}
