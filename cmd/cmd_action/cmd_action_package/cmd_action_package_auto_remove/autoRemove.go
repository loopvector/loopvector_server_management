/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_package_auto_remove

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action/cmd_action_package"
	"loopvector_server_management/controller"
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"

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
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		suggestions, err := controller.GetAllActiveServerNames()
		if err != nil {
			fmt.Println("Error querying database:", err)
			return nil, cobra.ShellCompDirectiveError
		}

		return suggestions, cobra.ShellCompDirectiveNoFileComp
	},
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		controller.RunAnsibleTasks(
			model.ServerNameModel{Name: args[0]},
			[]model.AnsibleTask{{FullPath: helper.KFullPathTaskPackageAutoRemove}},
			nil,
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
