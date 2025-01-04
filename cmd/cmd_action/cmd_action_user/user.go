/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_user

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/controller"

	"github.com/spf13/cobra"
)

var (
	username  string
	password  string
	usernames []string
	passwords []string
	group     string
	groups    []string
)

func GetActionUserCmd() *cobra.Command {
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("user called")
	},
}

func GetAllUsersToAdd() []controller.AddUsersToServerRequest {
	allUsers := []controller.AddUsersToServerRequest{}

	if username != "" && password != "" {
		allUsers = append(
			allUsers,
			controller.AddUsersToServerRequest{
				Username: username,
				Password: password,
				Groups:   _getGroups(),
			},
		)
	} else {
		if len(usernames) != len(passwords) {
			panic("number of usernames and passwords must be the same")
		}
		for i := 0; i < len(usernames); i++ {
			allUsers = append(
				allUsers,
				controller.AddUsersToServerRequest{
					Username: usernames[i],
					Password: passwords[i],
					Groups:   _getGroups(),
				},
			)
		}
	}

	return allUsers
}

func _getGroups() []string {
	allGroups := []string{}
	if group != "" {
		allGroups = append(allGroups, group)
	} else if len(groups) > 0 {
		allGroups = append(allGroups, groups...)
	}
	return allGroups
}

func init() {
	cmd_action.GetActionCmd().AddCommand(userCmd)

	userCmd.PersistentFlags().StringVar(&username, "username", "", "username of the user to be added to the server")
	userCmd.PersistentFlags().StringVar(&password, "password", "", "password of the user to be added to the server")

	userCmd.MarkFlagsRequiredTogether("username", "password")

	userCmd.PersistentFlags().StringSliceVar(&usernames, "usernames", []string{}, "username of the users to be added to the server")
	userCmd.PersistentFlags().StringSliceVar(&passwords, "passwords", []string{}, "password of the users to be added to the server")

	userCmd.MarkFlagsRequiredTogether("usernames", "passwords")

	userCmd.PersistentFlags().StringVar(&group, "group", "", "group to which the user is to be added")
	userCmd.PersistentFlags().StringSliceVar(&groups, "groups", []string{}, "groups to which the user is to be added")

	userCmd.MarkFlagsMutuallyExclusive("group", "groups")

	userCmd.MarkFlagsMutuallyExclusive("username", "usernames")
}
