package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

type LineToFileAddRequest struct {
	Line string
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

func AddLinesToFile(
	serverName string,
	addLinesToFileRequest AddLinesToFileRequest,
	sudoerEntries []LineToFileAddRequest,
) error {
	serverNameModel := model.ServerNameModel{Name: serverName}

	// serverId, err := serverNameModel.GetServerIdUsingServerName()
	// if err != nil {
	// 	panic(err)
	// }

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
		serverNameModel,
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
