/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd

import (
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/cmd/cmd_action/cmd_action_ufw"
	"loopvector_server_management/controller"

	"github.com/spf13/cobra"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		controller.EnableUfw(
			cmd_action.GetServerName(),
			cmd_action.GetServerSshConnectionInfo(),
		)
	},
}

func init() {
	cmd_action_ufw.GetActionUfwCmd().AddCommand(enableCmd)
}
