/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_package_auto_remove

import (
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/cmd/cmd_action/cmd_action_package"
	"loopvector_server_management/controller"

	"github.com/spf13/cobra"
)

// autoRemoveCmd represents the autoRemove command
var autoRemoveCmd = &cobra.Command{
	Use:   "autoRemove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		controller.PackageAutoRemove(
			cmd_action.GetServerName(),
			cmd_action.GetServerSshConnectionInfo(),
		)
	},
}

func init() {
	cmd_action_package.GetActionPackageCmd().AddCommand(autoRemoveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// autoRemoveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// autoRemoveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
