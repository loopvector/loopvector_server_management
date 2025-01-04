/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_sshd

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action"
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

func GetActionSshdCmd() *cobra.Command {
	return sshdCmd
}

// sshdCmd represents the sshd command
var sshdCmd = &cobra.Command{
	Use:   "sshd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sshd called")
	},
}

func AddASshdConfig() {
	fileFullPath := filepath.Join(sshdConfigFileCreatePath, fileName)
	controller.AddASshdConfig(
		cmd_action.GetServerName(),
		cmd_action.GetServerSshConnectionInfo(),
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
}

func init() {
	cmd_action.GetActionCmd().AddCommand(sshdCmd)

	sshdCmd.PersistentFlags().StringVar(&fileName, "fileName", "", "Custom sshd_config file name that will be created at /etc/ssh/sshd_config.d/ directory")
	//sshdCmd.Flags().StringVar(&filePermissionMode, "filePermissionMode", "", "File permission mode (e.g., 0644, 0755)")

	sshdCmd.MarkPersistentFlagRequired("fileName")
	//sshdCmd.MarkFlagRequired("filePermissionMode")

	sshdCmd.PersistentFlags().StringVar(&sshdConfigKey, "sshdConfigKey", "", "sshd config key ex Port, PermitRootLogin etc")
	sshdCmd.PersistentFlags().StringVar(&sshdConfigValue, "sshdConfigValue", "", "sshd config value ex 5623, no, yes etc")

	sshdCmd.MarkPersistentFlagRequired("sshdConfigKey")
	sshdCmd.MarkFlagsRequiredTogether("sshdConfigKey", "sshdConfigValue")

	sshdCmd.PersistentFlags().StringVar(&matchDirective, "matchDirectiveValue", "", "match directive value ex all, User alice, bob etc")

}
