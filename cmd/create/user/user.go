/*
Copyright © 2025 Agilan Anandan <agilan@loopvector.com>
*/
package user

import (
	"log"
	"loopvector_server_management/cmd/create"
	"loopvector_server_management/controller"
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
	"strings"

	"github.com/spf13/cobra"
)

var (
	username string
	email    string
	password string
)

var (
	loggedInUser model.User
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
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Validating session. create user command")
		var err error
		loggedInUser, err = controller.ValidateSession()
		if err != nil {
			log.Println(err.Error())
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if loggedInUser == (model.User{}) {
			hashedPassword, err := helper.Encrypt(password)
			if err != nil {
				panic(err)
			}
			var un *string = nil
			if strings.TrimSpace(username) != "" {
				un = &username
			}
			controller.RegisterUser(un, email, hashedPassword, false)
		} else {
			log.Println("Another user already logged in")
		}
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
