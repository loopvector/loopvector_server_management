/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action

import (
	"fmt"
	"loopvector_server_management/cmd"

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("action called")
	},
}

func init() {
	cmd.GetRootCmd().AddCommand(actionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// actionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// actionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
