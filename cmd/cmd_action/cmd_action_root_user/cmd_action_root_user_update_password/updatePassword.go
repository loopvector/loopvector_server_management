/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_root_user_update_password

import (
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/cmd/cmd_action/cmd_action_root_user"
	"loopvector_server_management/controller"

	"github.com/spf13/cobra"
)

// updatePasswordCmd represents the updatePassword command
var updatePasswordCmd = &cobra.Command{
	Use:   "updatePassword",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		// println("updatePassword called: ", args[1])
		controller.UpdateRootUserPassword(
			cmd_action.GetServerName(),
			cmd_action.GetServerSshConnectionInfo(),
			controller.UpdateRootUserPasswordRequest{NewRootPassword: args[0]},
		)
	},
}

func init() {
	cmd_action_root_user.GetActionRootUserCmd().AddCommand(updatePasswordCmd)
}
