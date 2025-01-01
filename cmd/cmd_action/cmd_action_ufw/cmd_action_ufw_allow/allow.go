/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_ufw_allow

import (
	"loopvector_server_management/cmd/cmd_action/cmd_action_ufw"
	"loopvector_server_management/cmd/cmd_action/cmd_action_ufw/helper"
	"loopvector_server_management/controller"

	"github.com/spf13/cobra"
)

// allowCmd represents the allow command
var allowCmd = &cobra.Command{
	Use:   "allow",
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
		return helper.RunUfwTrafficPolicyCommandE(
			cmd,
			args,
			controller.UfwTrafficPolicy{Policy: controller.UfwTrafficPolicyAllow},
		)
	},
}

func init() {
	cmd_action_ufw.GetActionUfwCmd().AddCommand(allowCmd)

	helper.InitUfwTrafficPolicyCommandFlags(allowCmd)
}
