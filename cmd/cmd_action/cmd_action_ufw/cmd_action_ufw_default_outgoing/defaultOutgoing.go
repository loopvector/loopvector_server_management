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

// defaultOutgoingCmd represents the defaultOutgoing command
var defaultOutgoingCmd = &cobra.Command{
	Use:   "defaultOutgoing",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	ValidArgs: []string{controller.UfwTrafficPolicyAllow, controller.UfwTrafficPolicyDeny},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		controller.SetDefaultOutgoingUfwTrafficPolicy(
			cmd_action.GetServerName(),
			cmd_action.GetServerSshConnectionInfo(),
			controller.UfwTrafficPolicy{Policy: args[0]},
		)
	},
}

func init() {
	cmd_action_ufw.GetActionUfwCmd().AddCommand(defaultOutgoingCmd)

}
