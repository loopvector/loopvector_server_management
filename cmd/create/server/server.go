/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package server

import (
	"loopvector_server_management/cmd/create"
	"loopvector_server_management/controller"

	"github.com/spf13/cobra"
)

var (
	displayName     string
	ipv4            string
	ipv4Subnet      uint64
	ipv6            string
	ipv6Subnet      uint64
	rootPassword    string
	adminUsername   string
	adminPassword   string
	rootUserSSHKey  string
	adminUserSSHKey string
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var csr = controller.CreateServerRequest{
			DisplayName:     displayName,
			IPv4:            &ipv4,
			IPv4Subnet:      &ipv4Subnet,
			IPv6:            &ipv6,
			IPv6Subnet:      &ipv6Subnet,
			RootPassword:    &rootPassword,
			AdminUsername:   &adminUsername,
			AdminPassword:   &adminPassword,
			RootUserSSHKey:  &rootUserSSHKey,
			AdminUserSSHKey: &adminUserSSHKey,
		}
		// fmt.Println(fmt.Sprintf("%+v", csr))
		controller.CreateNewServer(csr)
		return nil
	},
}

func init() {
	//println("serverCmd init called")
	create.GetCreateCmd().AddCommand(serverCmd)

	//println("passed serverCmd args: ", serverCmd.Args)

	serverCmd.Flags().StringVar(&displayName, "displayName", "", "display name for the server")

	serverCmd.Flags().StringVar(&ipv4, "ipv4", "", "ipv4 address of the server")

	serverCmd.Flags().Uint64Var(&ipv4Subnet, "ipv4Subnet", 32, "subnet mask for the ipv4 address of the server")

	serverCmd.Flags().StringVar(&ipv6, "ipv6", "", "ipv6 address of the server")

	serverCmd.Flags().Uint64Var(&ipv6Subnet, "ipv6Subnet", 128, "subnet mask for the ipv6 address of the server")

	serverCmd.Flags().StringVar(&rootPassword, "rootPassword", "", "root password for the server")

	serverCmd.Flags().StringVar(&adminUsername, "adminUsername", "", "admin username for the server")
	serverCmd.Flags().StringVar(&adminPassword, "adminPassword", "", "admin password for the server")

	serverCmd.Flags().StringVar(&rootUserSSHKey, "rootUserSSHKey", "", "admin user ssh key for the server")

	serverCmd.Flags().StringVar(&adminUserSSHKey, "adminUserSSHKey", "", "admin user ssh key for the server")

	serverCmd.MarkFlagRequired("displayName")
	serverCmd.MarkFlagsRequiredTogether("adminUsername", "adminPassword")
	serverCmd.MarkFlagsRequiredTogether("adminUsername", "adminUserSSHKey")
}
