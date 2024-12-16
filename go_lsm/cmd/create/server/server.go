/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
}

func init() {
	createCmd.AddCommand(serverCmd)

	var displayName string
	serverCmd.Flags().StringVarP(&displayName, "displayName", "dn", "", "display name for the server")
	serverCmd.MarkFlagRequired("displayName")

	var ipv4 string
	serverCmd.Flags().StringVarP(&ipv4, "ipv4", "v4", "", "ipv4 address of the server")

	var ipv4Subnet int
	serverCmd.Flags().IntVarP(&ipv4Subnet, "ipv4Subnet", "v4sn", 24, "subnet mask for the ipv4 address of the server")

	var ipv6 string
	serverCmd.Flags().StringVarP(&ipv6, "ipv6", "v6", "", "ipv6 address of the server")

	var ipv6Subnet int
	serverCmd.Flags().IntVarP(&ipv6Subnet, "ipv6Subnet", "v6sn", 112, "subnet mask for the ipv6 address of the server")

	var rootPassword string
	serverCmd.Flags().StringVarP(&rootPassword, "rootPassword", "rup", "", "root password for the server")

	var adminUsername string
	var adminPassword string
	serverCmd.Flags().StringVarP(&adminUsername, "adminUsername", "aun", "", "admin username for the server")
	serverCmd.Flags().StringVarP(&adminPassword, "adminPassword", "aup", "", "admin password for the server")
	serverCmd.MarkFlagsRequiredTogether("adminUsername", "adminPassword")

	var rootUserSSHKey string
	serverCmd.Flags().StringVarP(&rootUserSSHKey, "rootUserSSHKey", "ruk", "", "admin user ssh key for the server")

	var adminUserSSHKey string
	serverCmd.Flags().StringVarP(&adminUserSSHKey, "adminUserSSHKey", "auk", "", "admin user ssh key for the server")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
