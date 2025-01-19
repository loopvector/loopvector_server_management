/*
Copyright © 2025 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_sudoers_view

import (
	"loopvector_server_management/cmd/cmd_action/cmd_action_sudoers"

	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd_action_sudoers.ViewSudoers()
		return nil
	},
}

func init() {
	cmd_action_sudoers.GetActionSudoersCmd().AddCommand(viewCmd)
}
