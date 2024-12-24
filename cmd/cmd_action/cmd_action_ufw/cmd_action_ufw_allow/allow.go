/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_ufw_allow

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action/cmd_action_ufw"

	"github.com/spf13/cobra"
)

// allowCmd represents the allow command
var allowCmd = &cobra.Command{
	Use:   "allow",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("allow called")
	},
}

func init() {
	cmd_action_ufw.GetActionUfwCmd().AddCommand(allowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
