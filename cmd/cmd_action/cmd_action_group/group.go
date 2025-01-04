/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_group

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action"

	"github.com/spf13/cobra"
)

var (
	groupsToAdd []string
	groupToAdd  string
)

func GetActionGroupCmd() *cobra.Command {
	return groupCmd
}

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("group called")
	},
}

func GetAllGroupsToAdd() []string {
	var allGroupsToAdd []string
	if groupToAdd != "" {
		allGroupsToAdd = append(allGroupsToAdd, groupToAdd)
	}
	if len(groupsToAdd) > 0 {
		allGroupsToAdd = append(allGroupsToAdd, groupsToAdd...)
	}
	return allGroupsToAdd
}

func init() {
	cmd_action.GetActionCmd().AddCommand(groupCmd)

	groupCmd.PersistentFlags().StringSliceVar(&groupsToAdd, "groups", []string{}, "add the list of groups to the server")
	groupCmd.PersistentFlags().StringVar(&groupToAdd, "group", "", "add the group to the server")

	groupCmd.MarkFlagsOneRequired("groups", "group")
}
