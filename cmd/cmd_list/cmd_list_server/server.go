/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_list_server

import (
	"loopvector_server_management/cmd/cmd_list"
	"loopvector_server_management/controller"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	ValidArgs: []string{"all"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		controller.ListAllServers()
	},
}

func init() {
	cmd_list.GetListCmd().AddCommand(serverCmd)
}
