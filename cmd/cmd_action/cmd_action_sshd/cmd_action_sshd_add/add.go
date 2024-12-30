/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_sshd_add

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action/cmd_action_sshd"
	"loopvector_server_management/controller"
	"path/filepath"

	"github.com/spf13/cobra"
)

const sshdConfigFileCreatePath = "/etc/ssh/sshd_config.d"

var (
	//port                        uint16
	sshdConfigKey         string
	sshdConfigValue       string
	matchDirective        string
	fileName              string
	filePermissionMode    = "0755"
	fileDirPermissionMode = "0755"
	// usersPermitRootLogin        []string
	// permitRootLogin             bool
	// usersPubkeyAuthentication   []string
	// pubkeyAuthentication        bool
	// usersPasswordAuthentication []string
	// passwordAuthentication      bool
	// usersPermitEmptyPassword    []string
	// permitEmptyPassword         bool
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
		suggestions, err := controller.GetAllActiveServerNames()
		if err != nil {
			fmt.Println("Error querying database:", err)
			return nil, cobra.ShellCompDirectiveError
		}

		return suggestions, cobra.ShellCompDirectiveNoFileComp
	},
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		serverName := args[0]
		fileFullPath := filepath.Join(sshdConfigFileCreatePath, fileName)
		controller.AddASshdConfig(
			serverName,
			controller.AddLinesToFileRequest{
				FileFullPath:      fileFullPath,
				FilePermission:    filePermissionMode,
				AsSudo:            true,
				FileDirPermission: fileDirPermissionMode,
			},
			controller.SSHDConfigAddRequest{
				Key:            sshdConfigKey,
				Value:          sshdConfigValue,
				MatchDirective: matchDirective,
			},
		)
	},
}

func init() {
	cmd_action_sshd.GetActionSshdCmd().AddCommand(addCmd)

	addCmd.Flags().StringVar(&fileName, "fileName", "", "Custom sshd_config file name that will be created at /etc/ssh/sshd_config.d/ directory")
	//sshdCmd.Flags().StringVar(&filePermissionMode, "filePermissionMode", "", "File permission mode (e.g., 0644, 0755)")

	addCmd.MarkFlagRequired("fileName")
	//sshdCmd.MarkFlagRequired("filePermissionMode")

	addCmd.Flags().StringVar(&sshdConfigKey, "sshdConfigKey", "", "sshd config key ex Port, PermitRootLogin etc")
	addCmd.Flags().StringVar(&sshdConfigValue, "sshdConfigValue", "", "sshd config value ex 5623, no, yes etc")

	addCmd.MarkFlagRequired("sshdConfigKey")
	addCmd.MarkFlagsRequiredTogether("sshdConfigKey", "sshdConfigValue")

	addCmd.Flags().StringVar(&matchDirective, "matchDirectiveValue", "", "match directive value ex all, User alice, bob etc")
}
