/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_install

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/controller"
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"

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
			[]model.AnsibleTask{{FullPath: helper.KFullPathTaskInstallApps, Vars: appsToInstallList}},
		)
	},
}

func init() {
	cmd_action.GetActionCmd().AddCommand(installCmd)

	installCmd.Flags().StringArrayVar(&appsToInstallList, "apps", []string{}, "install the list of apps to the server")

	installCmd.MarkFlagRequired("apps")
}
