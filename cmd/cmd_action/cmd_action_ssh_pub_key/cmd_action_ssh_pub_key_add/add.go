/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_ssh_pub_key_add

import (
	"errors"
	"loopvector_server_management/cmd/cmd_action/cmd_action_ssh_pub_key"
	"loopvector_server_management/controller"
	"path/filepath"

	"github.com/spf13/cobra"
)

const sshPubKeyFileFullPath = ".ssh/authorized_keys"

var (
	username  string
	usernames []string
	key       string
	keys      []string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		suggestions := controller.GetAllActiveServerNames()
		// if err != nil {
		// 	fmt.Println("Error querying database:", err)
		// 	return nil, cobra.ShellCompDirectiveError
		// }

		return suggestions, cobra.ShellCompDirectiveNoFileComp
	},
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		serverName := args[0]

		if username == "" && len(usernames) == 0 {
			return errors.New("no username(s) provided")
		}

		if key == "" && len(keys) == 0 {
			return errors.New("no key(s) provided")
		}

		homeDirectory := "/home"
		authorizedKeysFilePathSuffix := ".ssh/authorized_keys"

		allUsernames := []string{}
		if len(usernames) > 0 {
			allUsernames = append(allUsernames, usernames...)
		} else {
			allUsernames = append(allUsernames, username)
		}

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

		for _, u := range allUsernames {
			userRootDirectory := filepath.Join(homeDirectory, u)
			userAuthorizedKeysFilePath := filepath.Join(userRootDirectory, authorizedKeysFilePathSuffix)
			//pathsToAddPubKeys = append(pathsToAddPubKeys, userAuthorizedKeysFilePath)
			lines := []controller.LineToFileAddRequest{}
			for _, key := range allKeys {
				lines = append(lines, controller.LineToFileAddRequest{Line: key})
			}
			controller.AddLinesToFile(
				serverName,
				controller.AddLinesToFileRequest{
					FileFullPath:      userAuthorizedKeysFilePath,
					FilePermission:    "0600",
					FileDirPermission: "0700",
					FileOwnerUsername: u,
					FileDirOwner:      u,
					AsSudo:            true,
				},
				lines,
			)
		}
		return nil
	},
}

func init() {
	cmd_action_ssh_pub_key.GetActionSshPubKeyCmd().AddCommand(addCmd)

	addCmd.Flags().StringVar(&key, "key", "", "pub key that can be used to ssh into the server")
	addCmd.Flags().StringSliceVar(&keys, "keys", []string{}, "pub keys that can be used to ssh into the server")

	addCmd.MarkFlagsOneRequired("key", "keys")
	addCmd.MarkFlagsMutuallyExclusive("key", "keys")

	addCmd.Flags().StringVar(&username, "username", "", "username that can be used to ssh into the server using the provided key(s)")
	addCmd.Flags().StringSliceVar(&usernames, "usernames", []string{}, "usernames that can be used to ssh into the server using the provided key(s)")

	addCmd.MarkFlagsOneRequired("username", "usernames")
	addCmd.MarkFlagsMutuallyExclusive("username", "usernames")

}
