/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package create

import (
	"fmt"
	"loopvector_server_management/cmd"

	"github.com/spf13/cobra"
)

func GetCreateCmd() *cobra.Command {
	return createCmd
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	// println("createCmd init called")
	cmd.GetRootCmd().AddCommand(createCmd)

	// println("passed createCmd args: ", createCmd.Args)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
