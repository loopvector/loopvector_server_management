/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action

import (
	"fmt"
	"loopvector_server_management/cmd"
	"loopvector_server_management/controller"

	"github.com/spf13/cobra"
)

func GetActionCmd() *cobra.Command {
	return actionCmd
}

// actionCmd represents the action command
var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// 	suggestions, err := controller.GetAllActiveServerNames()
	// 	if err != nil {
	// 		fmt.Println("Error querying database:", err)
	// 		return nil, cobra.ShellCompDirectiveError
	// 	}

	// 	return suggestions, cobra.ShellCompDirectiveNoFileComp
	// },
	//ValidArgs: controller.GetAllActiveServerNamesWithoutError(),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		println("root action persistentprerune called")
		// Validate the server name
		if ServerName == "" {
			println("a server name is required")
			return fmt.Errorf("a server name is required")
		}
		if !_isValidServerName(ServerName) {
			println("invalid server name: %s", ServerName)
			return fmt.Errorf("invalid server name: %s", ServerName)
		}
		return nil
	},
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	// Ensure the first argument is a valid server name
	// 	if len(args) < 1 {
	// 		return fmt.Errorf("a server name is required")
	// 	}
	// 	if !_isValidServerName(args[0]) {
	// 		return fmt.Errorf("invalid server name: %s", args[0])
	// 	}
	// 	return nil
	// },
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("parent action called")
	// },
}

func _isValidServerName(serverName string) bool {
	validServers := controller.GetAllActiveServerNamesWithoutError()
	for _, name := range validServers {
		if name == serverName {
			return true
		}
	}
	return false
}

var (
	ServerName string
)

func init() {
	cmd.GetRootCmd().AddCommand(actionCmd)

	actionCmd.PersistentFlags().StringVarP(&ServerName, "serverName", "n", "", "unique name of the server to be managed")

	actionCmd.MarkPersistentFlagRequired("serverName")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// actionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// actionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
