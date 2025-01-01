/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_package_upgrade

import (
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/cmd/cmd_action/cmd_action_package"
	"loopvector_server_management/controller"
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"

	"github.com/spf13/cobra"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		controller.RunAnsibleTasks(
			model.ServerNameModel{Name: cmd_action.ServerName},
			[]model.AnsibleTask{{FullPath: helper.KFullPathTaskPackageUpgrade}},
			nil,
		)
	},
}

func init() {
	cmd_action_package.GetActionPackageCmd().AddCommand(upgradeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upgradeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upgradeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
