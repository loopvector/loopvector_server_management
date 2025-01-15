/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action

import (
	"errors"
	"fmt"
	"loopvector_server_management/cmd"
	"loopvector_server_management/controller"
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"

	"github.com/spf13/cobra"
)

var (
	serverName   string
	username     string
	userPassword string
	// usernames  []string
	ipv4 string
	ipv6 string
	port uint16
	// sshKey             string
	sshPvtKeyFullFilePath string
)

func GetActionCmd() *cobra.Command {
	return actionCmd
}

func GetServerSshConnectionInfo() model.ServerSshConnectionInfo {
	var err error = nil
	currentServerName := GetServerName()
	var currentServerIpv4 model.ServerIpv4
	if ipv4 != "" {
		currentServerIpv4, err = model.ServerIpv4{Ip: ipv4}.GetServerIpv4UsingIpAddress()
		if err != nil {
			println("Error getting ipv4: ", err.Error())
		}
	} else {
		currentServerIpv4, err = currentServerName.GetIpv4UsingServerName()
		if err != nil {
			println("Error getting ipv4: ", err.Error())
		}
	}
	var currentServerIpv6 model.ServerIpv6
	if currentServerIpv4 == (model.ServerIpv4{}) {
		if ipv6 != "" {
			currentServerIpv6, err = model.ServerIpv6{Ip: ipv6}.GetServerIpv6UsingIpAddress()
			if err != nil {
				println("Error getting ipv6: ", err.Error())
			}
		} else {
			currentServerIpv6, err = currentServerName.GetIpv6UsingServerName()
			if err != nil {
				println("Error getting ipv6: ", err.Error())
			}
		}
	}
	if currentServerIpv4 == (model.ServerIpv4{}) &&
		currentServerIpv6 == (model.ServerIpv6{}) {
		panic(errors.New("neither ipv4 or ipv6 is set"))
	}

	currentServerSshIp := ""
	if currentServerIpv4 != (model.ServerIpv4{}) {
		currentServerSshIp = currentServerIpv4.Ip
	} else {
		currentServerSshIp = currentServerIpv6.Ip
	}

	currentServerSshPort := port

	currentServerSshUsername := ""
	currentServerSshUserPassword := ""
	if username != "" && userPassword != "" {
		currentServerSshUsername = username
		currentServerSshUserPassword = userPassword
	} else {
		var currentServerUser model.ServerUser
		var currentServerRootUser model.ServerRootUser
		if username == helper.KRootUserUsername {
			currentServerRootUser, err = currentServerName.GetRootUserUsingServerName()
			if err != nil {
				println("Error getting username: ", err.Error())
			}
		} else {
			currentServerUser, err = currentServerName.GetServerUserUsingServerName(username)
			if err != nil {
				println("Error getting username: ", err.Error())
			}
		}
		if currentServerUser == (model.ServerUser{}) &&
			currentServerRootUser == (model.ServerRootUser{}) {
			panic(errors.New("neither username or root user is set"))
		}

		if currentServerUser != (model.ServerUser{}) {
			currentServerSshUsername = currentServerUser.Username
			currentServerSshUserPassword = currentServerUser.Password
		} else {
			currentServerSshUsername = helper.KRootUserUsername
			currentServerSshUserPassword = currentServerRootUser.Password
		}
	}

	currentServerSshPvtKeyFullFilePath := ""
	if sshPvtKeyFullFilePath != "" {
		currentServerSshPvtKeyFullFilePath = sshPvtKeyFullFilePath
	}

	return model.ServerSshConnectionInfo{
		ServerName:            currentServerName.Name,
		Username:              currentServerSshUsername,
		Password:              currentServerSshUserPassword,
		Port:                  currentServerSshPort,
		Ip:                    currentServerSshIp,
		SshPvtKeyFullFilePath: currentServerSshPvtKeyFullFilePath,
	}

}

// actionCmd represents the action command
var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// 	suggestions, err := controller.GetAllActiveServerNames()
	// 	if err != nil {
	// 		fmt.Println("Error querying database:", err)
	// 		return nil, cobra.ShellCompDirectiveError
	// 	}

	// 	return suggestions, cobra.ShellCompDirectiveNoFileComp
	// },
	//ValidArgs: controller.GetAllActiveServerNamesWithoutError(),
	PersistentPreRunE: func(c *cobra.Command, args []string) error {
		// println("root action persistentprerune called")
		err := cmd.ValidateLogin()
		if err != nil {
			// Validate the server name
			if serverName == "" {
				println("a server name is required")
				return fmt.Errorf("a server name is required")
			}
			if !_isValidServerName(serverName) {
				println("invalid server name: %s", serverName)
				return fmt.Errorf("invalid server name: %s", serverName)
			}
			return nil
		} else {
			return err
		}
	},
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	// Ensure the first argument is a valid server name
	// 	if len(args) < 1 {
	// 		return fmt.Errorf("a server name is required")
	// 	}
	// 	if !_isValidServerName(args[0]) {
	// 		return fmt.Errorf("invalid server name: %s", args[0])
	// 	}
	// 	return nil
	// },
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("parent action called")
	// },
}

func GetServerName() model.ServerNameModel {
	return model.ServerNameModel{Name: serverName}
}

func GetUsername() string {
	if username != "" {
		return username
	}
	return "root"
}

func _isValidServerName(serverName string) bool {
	validServers := controller.GetAllActiveServerNamesWithoutError()
	for _, name := range validServers {
		if name == serverName {
			return true
		}
	}
	return false
}

func init() {
	cmd.GetRootCmd().AddCommand(actionCmd)

	actionCmd.PersistentFlags().StringVar(&serverName, "serverName", "", "unique name of the server to be managed")

	actionCmd.MarkPersistentFlagRequired("serverName")

	actionCmd.PersistentFlags().StringVar(&username, "sshUsername", helper.KRootUserUsername, "username that can be used to ssh into the server")
	actionCmd.PersistentFlags().StringVar(&userPassword, "sshUserPassword", "", "ssh password for the provided username flag")
	// actionCmd.PersistentFlags().StringSliceVar(&usernames, "usernames", []string{}, "usernames that can be used to ssh into the server")

	// actionCmd.MarkFlagsMutuallyExclusive("username", "usernames")

	actionCmd.PersistentFlags().StringVar(&ipv4, "sshIpv4", "", "ipv4 address to be used to connect to the server")
	actionCmd.PersistentFlags().StringVar(&ipv6, "sshIpv6", "", "ipv6 address to be used to connect to the server")
	actionCmd.PersistentFlags().Uint16Var(&port, "sshPort", 22, "port to be used to connect to the server")
	// actionCmd.PersistentFlags().StringVar(&sshKey, "sshKey", "", "ssh key that can be used to ssh into the server")
	actionCmd.PersistentFlags().StringVar(&sshPvtKeyFullFilePath, "sshPvtKeyFullFilePath", "", "ssh private key pair full file path that can be used to ssh into the server")
}
