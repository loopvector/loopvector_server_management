/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_ufw_deny

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action/cmd_action_ufw"

	"github.com/spf13/cobra"
)

// denyCmd represents the deny command
var denyCmd = &cobra.Command{
	Use:   "deny",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deny called")
	},
}

func init() {
	cmd_action_ufw.GetActionUfwCmd().AddCommand(denyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// denyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// denyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
