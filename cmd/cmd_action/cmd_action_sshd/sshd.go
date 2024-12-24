/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_sshd

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action"

	"github.com/spf13/cobra"
)

func GetActionSshdCmd() *cobra.Command {
	return sshdCmd
}

// sshdCmd represents the sshd command
var sshdCmd = &cobra.Command{
	Use:   "sshd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sshd called")
	},
}

func init() {
	cmd_action.GetActionCmd().AddCommand(sshdCmd)

	

}
