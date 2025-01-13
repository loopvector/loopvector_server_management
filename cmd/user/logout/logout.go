/*
Copyright © 2025 Agilan Anandan <agilan@loopvector.com>
*/
package logout

import (
	"log"
	"loopvector_server_management/cmd/user"
	"loopvector_server_management/controller"
	"loopvector_server_management/model"

	"github.com/spf13/cobra"
)

var (
	email string
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if user.GetLoggedInUser() == (model.User{}) {
			log.Fatal("you are not logged in")
		} else if user.GetLoggedInUser().Email == email {
			controller.LogoutUser()
		} else {
			log.Fatal("you are not logged in as this user")
		}
	},
}

func init() {
	user.GetUserCmd().AddCommand(logoutCmd)

	logoutCmd.Flags().StringVar(&email, "email", "", "email of the user of this app")
}
