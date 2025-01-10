/*
Copyright © 2025 Agilan Anandan <agilan@loopvector.com>
*/
package cmd

import (
	"loopvector_server_management/cmd/create"
	"loopvector_server_management/controller"
	"loopvector_server_management/controller/helper"

	"github.com/spf13/cobra"
)

var (
	username string
	email    string
	password string
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		hashedPassword := helper.HashPassword(password)
		controller.RegisterUser(username, email, hashedPassword)
	},
}

func init() {
	create.GetCreateCmd().AddCommand(userCmd)

	userCmd.Flags().StringVar(&username, "username", "", "username of the user of this app")
	userCmd.Flags().StringVar(&email, "email", "", "email of the user of this app")
	userCmd.Flags().StringVar(&password, "password", "", "password of the user of this app")

	userCmd.MarkFlagRequired("email")
	userCmd.MarkFlagRequired("password")
}
