/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_sshd_add

import (
	"loopvector_server_management/cmd/cmd_action/cmd_action_sshd"

	"github.com/spf13/cobra"
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
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd_action_sshd.AddASshdConfig()
		return nil
	},
}

func init() {
	cmd_action_sshd.GetActionSshdCmd().AddCommand(addCmd)

}
