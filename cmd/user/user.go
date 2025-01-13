/*
Copyright © 2025 Agilan Anandan <agilan@loopvector.com>
*/
package user

import (
	"log"
	"loopvector_server_management/cmd"
	"loopvector_server_management/controller"
	"loopvector_server_management/model"

	"github.com/spf13/cobra"
)

var (
	loggedInUser model.User
)

func GetLoggedInUser() model.User {
	return loggedInUser
}

func GetUserCmd() *cobra.Command {
	return userCmd
}

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
		log.Println("Validating session. user command")
		var err error
		loggedInUser, err = controller.ValidateSession()
		if err != nil {
			log.Println(err.Error())
		}
		return nil
	},
}

func init() {
	cmd.GetRootCmd().AddCommand(userCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
