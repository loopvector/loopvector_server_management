/*
Copyright © 2025 Agilan Anandan <agilan@loopvector.com>
*/
package cmd

import (
	"loopvector_server_management/cmd/create"
	"loopvector_server_management/controller"
	"strings"

	"github.com/spf13/cobra"
)

var (
	username       string
	email          string
	hashedPassword string
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var un *string = nil
		if strings.TrimSpace(username) != "" {
			un = &username
		}
		controller.RegisterUser(un, email, hashedPassword, true)
	},
}

func init() {
	create.GetCreateCmd().AddCommand(adminCmd)

	adminCmd.Flags().StringVar(&username, "username", "", "username of the user of this app")
	adminCmd.Flags().StringVar(&email, "email", "", "email of the user of this app")
	adminCmd.Flags().StringVar(&hashedPassword, "hashedPassword", "", "password of the user of this app")

	adminCmd.MarkFlagRequired("email")
	adminCmd.MarkFlagRequired("hashedPassword")
}
