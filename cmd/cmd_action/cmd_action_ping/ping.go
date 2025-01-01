/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_ping

import (
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/controller"
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use: "ping",
	// TraverseChildren: true,
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
	// 	if parent := cmd.Parent(); parent != nil {
	// 		if parent.PersistentPreRunE != nil {
	// 			return parent.PersistentPreRunE(parent, args)
	// 		}
	// 	}
	// 	return nil
	// },
	RunE: func(cmd *cobra.Command, args []string) error {
		controller.RunAnsibleTasks(
			model.ServerNameModel{Name: cmd_action.ServerName},
			[]model.AnsibleTask{{FullPath: helper.KFullPathTaskPing}},
			nil,
		)
		return nil
	},
}

func init() {
	cmd_action.GetActionCmd().AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
