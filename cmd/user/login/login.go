/*
Copyright © 2025 Agilan Anandan <agilan@loopvector.com>
*/
package login

import (
	"log"
	"loopvector_server_management/cmd/user"
	"loopvector_server_management/controller"
	"loopvector_server_management/model"

	"github.com/spf13/cobra"
)

var (
	email    string
	password string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if user.GetLoggedInUser() == (model.User{}) {
			controller.LoginUser(email, password)
		} else if user.GetLoggedInUser().Email == email {
			log.Println("User already logged in")
		} else if user.GetLoggedInUser().Email != email {
			log.Println("Another user already logged in")
		}
	},
}

func init() {
	user.GetUserCmd().AddCommand(loginCmd)

	loginCmd.Flags().StringVar(&email, "email", "", "email of the user of this app")
	loginCmd.Flags().StringVar(&password, "password", "", "password of the user of this app")
}
