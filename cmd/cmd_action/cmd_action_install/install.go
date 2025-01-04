/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_install

import (
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/controller"

	"github.com/spf13/cobra"
)

var (
	appsToInstallList []string
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		controller.InstallServerApps(
			cmd_action.GetServerName(),
			cmd_action.GetServerSshConnectionInfo(),
			appsToInstallList,
		)
	},
}

func init() {
	cmd_action.GetActionCmd().AddCommand(installCmd)

	installCmd.Flags().StringSliceVar(&appsToInstallList, "apps", []string{}, "install the list of apps to the server")

	installCmd.MarkFlagRequired("apps")
}
