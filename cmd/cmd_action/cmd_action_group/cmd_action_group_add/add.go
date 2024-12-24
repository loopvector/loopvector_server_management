/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_group_add

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action/cmd_action_group"
	"loopvector_server_management/controller"

	"github.com/spf13/cobra"
)

var (
	groupsToAdd []string
	groupToAdd  string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
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
		if groupToAdd != "" {
			controller.AddGroupsToServer(args[0], []string{groupToAdd})
		}
		if len(groupsToAdd) != 0 {
			controller.AddGroupsToServer(args[0], groupsToAdd)
		}
	},
}

func init() {
	cmd_action_group.GetActionGroupCmd().AddCommand(addCmd)

	addCmd.Flags().StringSliceVar(&groupsToAdd, "groups", []string{}, "add the list of groups to the server")
	addCmd.Flags().StringVar(&groupToAdd, "group", "", "add the group to the server")

	addCmd.MarkFlagsOneRequired("groups", "group")
}
