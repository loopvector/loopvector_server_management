/*
Copyright © 2025 Agilan Anandan <agilan@loopvector.com>
*/
package cmd

import (
	"loopvector_server_management/cmd/cmd_action/cmd_action_sshd"

	"github.com/spf13/cobra"
)

var (
	identifier string
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd_action_sshd.DeleteASshdConfig(identifier)
	},
}

func init() {
	cmd_action_sshd.GetActionSshdCmd().AddCommand(deleteCmd)

	deleteCmd.Flags().StringVar(&identifier, "identifier", "", "identifier of the block to be deleted")
	deleteCmd.MarkFlagRequired("identifier")
}
