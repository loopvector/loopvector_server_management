/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_delete_server

import (
	"loopvector_server_management/cmd/cmd_delete"
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
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		controller.DeleteServer(args[0])
		return nil
	},
}

func init() {
	cmd_delete.GetDeleteCmd().AddCommand(serverCmd)
}
