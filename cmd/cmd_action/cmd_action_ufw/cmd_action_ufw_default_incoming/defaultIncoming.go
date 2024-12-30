/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action/cmd_action_ufw"
	"loopvector_server_management/controller"
	"loopvector_server_management/model"

	"github.com/spf13/cobra"
)

// defaultIncomingCmd represents the defaultIncoming command
var defaultIncomingCmd = &cobra.Command{
	Use:   "defaultIncoming",
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

		trafficPolicies := []string{controller.UfwTrafficPolicyAllow, controller.UfwTrafficPolicyDeny}
		suggestions = append(suggestions, trafficPolicies...)
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	},
	Args: cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		controller.SetDefaultIncomingUfwTrafficPolicy(
			model.ServerNameModel{Name: args[0]},
			controller.UfwTrafficPolicy{Policy: args[1]},
		)
	},
}

func init() {
	cmd_action_ufw.GetActionUfwCmd().AddCommand(defaultIncomingCmd)

}
