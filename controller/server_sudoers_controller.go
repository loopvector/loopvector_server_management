package controller

import (
	"fmt"
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
	serverName string,
	addLinesToFileRequest AddLinesToFileRequest,
	sudoerEntry SudoersAddRequest,
) error {
	var lines []LineToFileAddRequest
	lines = append(
		lines,
		_getLineFromRequest(sudoerEntry),
	)
	AddLinesToFile(
		serverName,
		addLinesToFileRequest,
		lines,
	)
	return nil
}

func AddSudoers(
	serverName string,
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
	AddLinesToFile(
		serverName,
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
