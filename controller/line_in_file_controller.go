package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

type LineToFileAddRequest struct {
	Line string
}

type LineBlockToFileAddRequest struct {
	Block string
}

type AddLinesToFileRequest struct {
	FileFullPath      string
	FilePermission    string
	AsSudo            bool
	FileOwnerUsername string
	FileDirPermission string
	FileDirOwner      string
	//FilePathOwner     string
}

func AddLineBlockToFile(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	addLinesToFileRequest AddLinesToFileRequest,
	lines []LineToFileAddRequest,
) error {

	var linesToAddVar []string
	for _, sudoerEntry := range lines {
		linesToAddVar = append(linesToAddVar, sudoerEntry.Line)
	}

	vars := map[string]interface{}{
		"lines_to_add":   linesToAddVar,
		"mode":           addLinesToFileRequest.FilePermission,
		"file_path":      addLinesToFileRequest.FileFullPath,
		"should_become":  addLinesToFileRequest.AsSudo,
		"owner_user":     addLinesToFileRequest.FileOwnerUsername,
		"dir_mode":       addLinesToFileRequest.FileDirPermission,
		"dir_owner_user": addLinesToFileRequest.FileDirOwner,
	}

	_, err := RunAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		[]model.AnsibleTask{{
			FullPath: helper.KFullPathTaskAddLineBlockToFile,
			Vars:     vars,
		}},
		nil,
	)

	if err != nil {
		panic(err)
	}

	return nil

}

func AddLinesToFile(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	addLinesToFileRequest AddLinesToFileRequest,
	sudoerEntries []LineToFileAddRequest,
) error {
	var linesToAddVar []string
	for _, sudoerEntry := range sudoerEntries {
		linesToAddVar = append(linesToAddVar, sudoerEntry.Line)
	}
	vars := map[string]interface{}{
		"lines_to_add":        linesToAddVar,
		"file_mode":           addLinesToFileRequest.FilePermission,
		"file_full_path":      addLinesToFileRequest.FileFullPath,
		"as_sudo":             addLinesToFileRequest.AsSudo,
		"file_owner_user":     addLinesToFileRequest.FileOwnerUsername,
		"file_dir_mode":       addLinesToFileRequest.FileDirPermission,
		"file_dir_owner_user": addLinesToFileRequest.FileDirOwner,
	}

	_, err := RunAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		[]model.AnsibleTask{{
			FullPath: helper.KFullPathTaskAddLinesToFile,
			Vars:     vars,
		}},
		nil,
	)

	if err != nil {
		panic(err)
	}

	return nil
}
