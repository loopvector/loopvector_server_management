/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_package

import (
	"loopvector_server_management/cmd/cmd_action"

	"github.com/spf13/cobra"
)

func GetActionPackageCmd() *cobra.Command {
	return packageCmd
}

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	cmd_action.GetActionCmd().AddCommand(packageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// packageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// packageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
