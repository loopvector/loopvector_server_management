/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_ssh_pub_key

import (
	"errors"
	"fmt"
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/controller"
	"path/filepath"

	"github.com/spf13/cobra"
)

// const sshPubKeyFileFullPath = ".ssh/authorized_keys"

var (
	key                   string
	keys                  []string
	filePermissionMode    = "0600"
	fileDirPermissionMode = "0700"
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

func AddSshKeys(cmd *cobra.Command) error {
	username := cmd_action.GetUsername()
	// if len(allUsernames) == 0 {
	// 	return errors.New("no username(s) provided")
	// }

	if key == "" && len(keys) == 0 {
		return errors.New("no key(s) provided")
	}

	homeDirectory := "/home"
	authorizedKeysFilePathSuffix := ".ssh/authorized_keys"

	// pathsToAddPubKeys := []string{}
	// for _, u := range allUsernames {
	// 	userRootDirectory := filepath.Join(homeDirectory, u)
	// 	userAuthorizedKeysFilePath := filepath.Join(userRootDirectory, authorizedKeysFilePathSuffix)
	// 	pathsToAddPubKeys = append(pathsToAddPubKeys, userAuthorizedKeysFilePath)
	// }

	allKeys := []string{}
	if len(keys) > 0 {
		allKeys = append(allKeys, keys...)
	} else {
		allKeys = append(allKeys, key)
	}

	// for _, u := range allUsernames {
	userRootDirectory := filepath.Join(homeDirectory, username)
	userAuthorizedKeysFilePath := filepath.Join(userRootDirectory, authorizedKeysFilePathSuffix)
	//pathsToAddPubKeys = append(pathsToAddPubKeys, userAuthorizedKeysFilePath)
	// lines := []controller.LineToFileAddRequest{}
	for _, key := range allKeys {
		oneLine := controller.LineToFileAddRequest{Line: key}
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
			},
			[]controller.LineToFileAddRequest{oneLine},
		)
	}
	return nil
}

func init() {
	cmd_action.GetActionCmd().AddCommand(sshPubKeyCmd)

	sshPubKeyCmd.PersistentFlags().StringVar(&key, "key", "", "pub key that can be used to ssh into the server")
	sshPubKeyCmd.PersistentFlags().StringSliceVar(&keys, "keys", []string{}, "pub keys that can be used to ssh into the server")

	sshPubKeyCmd.MarkFlagsOneRequired("key", "keys")
	sshPubKeyCmd.MarkFlagsMutuallyExclusive("key", "keys")

	// sshPubKeyCmd.MarkFlagsOneRequired("username", "usernames")
	// sshPubKeyCmd.MarkFlagsMutuallyExclusive("username", "usernames")
}
