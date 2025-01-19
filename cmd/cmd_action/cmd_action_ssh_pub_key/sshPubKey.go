/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_ssh_pub_key

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/controller"
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
	"path/filepath"

	"github.com/spf13/cobra"
)

// const sshPubKeyFileFullPath = ".ssh/authorized_keys"

var (
	filePermissionMode           = "0600"
	fileDirPermissionMode        = "0700"
	homeDirectory                = "/home"
	authorizedKeysFilePathSuffix = ".ssh/authorized_keys"
)

func GetActionSshPubKeyCmd() *cobra.Command {
	return sshPubKeyCmd
}

// sshPubKeyCmd represents the sshPubKey command
var sshPubKeyCmd = &cobra.Command{
	Use:   "sshPubKey",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sshPubKey called")
	},
}

func ViewKeys() error {
	username := cmd_action.GetUsername()
	userRootDirectory := filepath.Join(homeDirectory, username)
	userAuthorizedKeysFilePath := filepath.Join(userRootDirectory, authorizedKeysFilePathSuffix)
	controller.ReadBlocksFromFile(
		cmd_action.GetServerName(),
		cmd_action.GetServerSshConnectionInfo(),
		controller.ReadBlocksFromFileRequest{
			FileFullPath:     userAuthorizedKeysFilePath,
			AsSudo:           true,
			CommentDelimiter: helper.KCommentDelimiterHash,
		})
	return nil
}

func DeleteKey(identifier string) error {
	username := cmd_action.GetUsername()
	userRootDirectory := filepath.Join(homeDirectory, username)
	userAuthorizedKeysFilePath := filepath.Join(userRootDirectory, authorizedKeysFilePathSuffix)
	serverId, err := cmd_action.GetServerName().GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}
	serverUser, err := model.ServerUser{
		ServerID: serverId,
		Username: username,
	}.GetUsingServerIdAndUsername()
	if err != nil {
		panic(err)
	}
	callbacks := []controller.RunAnsibleTaskCallback{{
		TaskNames: []string{"delete line block from a file with block_timestamp: " + identifier},
		OnChanged: func() {
			model.AuthorizationKey{
				ServerUserID: serverUser.ID,
				ServerID:     serverId,
				Identifier:   identifier,
			}.DeleteUsingIdentifierUserIdAndServerId()
		},
		OnUnchanged: func() {},
		OnFailed:    func() {},
	}}
	controller.DeleteBlockFromFile(
		cmd_action.GetServerName(),
		cmd_action.GetServerSshConnectionInfo(),
		controller.DeleteBlockFromFileRequest{
			FileFullPath:     userAuthorizedKeysFilePath,
			AsSudo:           true,
			CommentDelimiter: helper.KCommentDelimiterHash,
			BlockTimestamp:   identifier,
		}, callbacks)
	return nil
}

func AddSshKeys(keys []string) error {
	username := cmd_action.GetUsername()
	// if len(allUsernames) == 0 {
	// 	return errors.New("no username(s) provided")
	// }

	// pathsToAddPubKeys := []string{}
	// for _, u := range allUsernames {
	// 	userRootDirectory := filepath.Join(homeDirectory, u)
	// 	userAuthorizedKeysFilePath := filepath.Join(userRootDirectory, authorizedKeysFilePathSuffix)
	// 	pathsToAddPubKeys = append(pathsToAddPubKeys, userAuthorizedKeysFilePath)
	// }

	serverId, err := cmd_action.GetServerName().GetServerIdUsingServerName()
	if err != nil {
		panic(err)
	}
	serverUser, err := model.ServerUser{
		ServerID: serverId,
		Username: username,
	}.GetUsingServerIdAndUsername()
	if err != nil {
		panic(err)
	}
	// for _, u := range allUsernames {
	userRootDirectory := filepath.Join(homeDirectory, username)
	userAuthorizedKeysFilePath := filepath.Join(userRootDirectory, authorizedKeysFilePathSuffix)
	//pathsToAddPubKeys = append(pathsToAddPubKeys, userAuthorizedKeysFilePath)
	// lines := []controller.LineToFileAddRequest{}
	for _, key := range keys {
		oneLine := controller.LineToFileAddRequest{Line: key}
		_, blockTimestamp := helper.GetCurrentTimestampMillis()
		callbacks := []controller.RunAnsibleTaskCallback{{
			TaskNames: []string{"add line block to a file", "add line block to a file and set permissions"},
			OnChanged: func() {
				model.AuthorizationKey{
					ServerUserID: serverUser.ID,
					ServerID:     serverId,
					PublicKey:    key,
					Identifier:   blockTimestamp,
				}.Create()
			},
			OnUnchanged: func() {},
			OnFailed:    func() {},
		}}
		controller.AddLineBlockToFile(
			cmd_action.GetServerName(),
			cmd_action.GetServerSshConnectionInfo(),
			controller.AddLinesToFileRequest{
				FileFullPath:      userAuthorizedKeysFilePath,
				FilePermission:    filePermissionMode,
				FileDirPermission: fileDirPermissionMode,
				FileOwnerUsername: username,
				FileDirOwner:      username,
				AsSudo:            true,
				BlockTimestamp:    blockTimestamp,
				CommentDelimiter:  helper.KCommentDelimiterHash,
			},
			[]controller.LineToFileAddRequest{oneLine},
			callbacks,
		)
	}
	return nil
}

func init() {
	cmd_action.GetActionCmd().AddCommand(sshPubKeyCmd)
}
