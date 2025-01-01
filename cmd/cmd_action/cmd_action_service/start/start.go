/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_service

import (
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/cmd/cmd_action/cmd_action_service"
	"loopvector_server_management/controller"
	"loopvector_server_management/model"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		controller.StartServices(
			model.ServerNameModel{Name: cmd_action.ServerName},
			controller.ServiceActionRequest{
				ServiceNames: cmd_action_service.ServiceNames,
			},
		)
	},
}

func init() {
	cmd_action_service.GetActionServiceCmd().AddCommand(startCmd)
}
