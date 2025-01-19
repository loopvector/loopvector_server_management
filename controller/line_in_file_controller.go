package controller

import (
	"encoding/json"
	"fmt"
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
	BlockTimestamp    string
	CommentDelimiter  string
	//FilePathOwner     string
}

type ReadBlocksFromFileRequest struct {
	FileFullPath     string
	AsSudo           bool
	CommentDelimiter string
}

type DeleteBlockFromFileRequest struct {
	FileFullPath     string
	AsSudo           bool
	CommentDelimiter string
	BlockTimestamp   string
}

func DeleteBlockFromFile(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	deleteBlockFromFileRequest DeleteBlockFromFileRequest,
	callbacks []RunAnsibleTaskCallback,
) error {
	vars := map[string]interface{}{
		"file_path":         deleteBlockFromFileRequest.FileFullPath,
		"should_become":     deleteBlockFromFileRequest.AsSudo,
		"comment_delimiter": deleteBlockFromFileRequest.CommentDelimiter,
		"block_timestamp":   deleteBlockFromFileRequest.BlockTimestamp,
	}
	_, err := RunAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		[]model.AnsibleTask{{
			FullPath: helper.KFullPathTaskDeleteBlock,
			Vars:     vars,
		}},
		callbacks,
	)
	if err != nil {
		panic(err)
	}

	return nil
}

func ReadBlocksFromFile(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	readBlocksFromFileRequest ReadBlocksFromFileRequest,
) error {
	vars := map[string]interface{}{
		"file_path":         readBlocksFromFileRequest.FileFullPath,
		"should_become":     readBlocksFromFileRequest.AsSudo,
		"comment_delimiter": readBlocksFromFileRequest.CommentDelimiter,
	}

	result, err := RunAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		[]model.AnsibleTask{{
			FullPath: helper.KFullPathTaskReadBlocks,
			Vars:     vars,
		}},
		nil,
	)

	for _, play := range result.Plays {
		for _, task := range play.Tasks {
			for host, hostData := range task.Hosts {
				if hostData.Msg != nil {
					// Parse the JSON message
					var msgContent map[string]string
					if err := json.Unmarshal(hostData.Msg, &msgContent); err != nil {
						// log.Printf("Error unmarshalling message for host %s: %v", host, err)
						continue
					}

					// Print the parsed message
					fmt.Printf("Host: %s\n", host)
					fmt.Println("Managed Blocks:")
					for blockID, blockContent := range msgContent {
						fmt.Printf("Identifier: %s\n", blockID)
						fmt.Printf("Block: %s\n", blockContent)
					}
				}
			}
		}
	}

	// log.Println("ansible task result: ", result)

	if err != nil {
		panic(err)
	}

	return nil
}

func AddLineBlockToFile(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
	addLinesToFileRequest AddLinesToFileRequest,
	lines []LineToFileAddRequest,
	callbacks []RunAnsibleTaskCallback,
) error {

	var linesToAddVar []string
	for _, sudoerEntry := range lines {
		linesToAddVar = append(linesToAddVar, sudoerEntry.Line)
	}

	vars := map[string]interface{}{
		"lines_to_add":      linesToAddVar,
		"mode":              addLinesToFileRequest.FilePermission,
		"file_path":         addLinesToFileRequest.FileFullPath,
		"should_become":     addLinesToFileRequest.AsSudo,
		"owner_user":        addLinesToFileRequest.FileOwnerUsername,
		"dir_mode":          addLinesToFileRequest.FileDirPermission,
		"dir_owner_user":    addLinesToFileRequest.FileDirOwner,
		"block_timestamp":   addLinesToFileRequest.BlockTimestamp,
		"comment_delimiter": addLinesToFileRequest.CommentDelimiter,
	}

	_, err := RunAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		[]model.AnsibleTask{{
			FullPath: helper.KFullPathTaskAddLineBlockToFile,
			Vars:     vars,
		}},
		callbacks,
	)

	if err != nil {
		panic(err)
	}

	return nil

}

// func AddLinesToFile(
// 	serverName model.ServerNameModel,
// 	serverSshConnectionInfo model.ServerSshConnectionInfo,
// 	addLinesToFileRequest AddLinesToFileRequest,
// 	sudoerEntries []LineToFileAddRequest,
// ) error {
// 	var linesToAddVar []string
// 	for _, sudoerEntry := range sudoerEntries {
// 		linesToAddVar = append(linesToAddVar, sudoerEntry.Line)
// 	}
// 	vars := map[string]interface{}{
// 		"lines_to_add":        linesToAddVar,
// 		"file_mode":           addLinesToFileRequest.FilePermission,
// 		"file_full_path":      addLinesToFileRequest.FileFullPath,
// 		"as_sudo":             addLinesToFileRequest.AsSudo,
// 		"file_owner_user":     addLinesToFileRequest.FileOwnerUsername,
// 		"file_dir_mode":       addLinesToFileRequest.FileDirPermission,
// 		"file_dir_owner_user": addLinesToFileRequest.FileDirOwner,
// 	}

// 	_, err := RunAnsibleTasks(
// 		serverName,
// 		serverSshConnectionInfo,
// 		[]model.AnsibleTask{{
// 			FullPath: helper.KFullPathTaskAddLinesToFile,
// 			Vars:     vars,
// 		}},
// 		nil,
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return nil
// }
