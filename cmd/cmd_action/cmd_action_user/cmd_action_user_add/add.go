/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_user_add

import (
	"errors"
	"loopvector_server_management/cmd/cmd_action/cmd_action_user"
	"loopvector_server_management/controller"
	"loopvector_server_management/model"

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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		suggestions := controller.GetAllActiveServerNames()
		// if err != nil {
		// 	fmt.Println("Error querying database:", err)
		// 	return nil, cobra.ShellCompDirectiveError
		// }

		return suggestions, cobra.ShellCompDirectiveNoFileComp
	},
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		serverName := model.ServerNameModel{Name: args[0]}

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
				return errors.New("number of usernames and passwords must be the same")
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

		// for _, oneUser := range allUsers {
		// 	println("oneUser: ", oneUser.Groups)
		// }

		controller.AddUsersToServer(serverName, allUsers)
		return nil
	},
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
	cmd_action_user.GetActionUserCmd().AddCommand(addCmd)

	addCmd.Flags().StringVar(&username, "username", "", "username of the user to be added to the server")
	addCmd.Flags().StringVar(&password, "password", "", "password of the user to be added to the server")

	addCmd.MarkFlagsRequiredTogether("username", "password")

	addCmd.Flags().StringSliceVar(&usernames, "usernames", []string{}, "username of the users to be added to the server")
	addCmd.Flags().StringSliceVar(&passwords, "passwords", []string{}, "password of the users to be added to the server")

	addCmd.MarkFlagsRequiredTogether("usernames", "passwords")

	addCmd.Flags().StringVar(&group, "group", "", "group to which the user is to be added")
	addCmd.Flags().StringSliceVar(&groups, "groups", []string{}, "groups to which the user is to be added")

	addCmd.MarkFlagsMutuallyExclusive("group", "groups")

	addCmd.MarkFlagsMutuallyExclusive("username", "usernames")
}
