/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_service

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action"

	"github.com/spf13/cobra"
)

var (
	serviceNames []string
	serviceName  string
)

func GetActionServiceCmd() *cobra.Command {
	return serviceCmd
}

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service called")
	},
}

func GetAllServiceNames() []string {
	var allServiceNames []string
	if serviceName != "" {
		allServiceNames = append(allServiceNames, serviceName)
	}
	if serviceNames != nil && len(serviceNames) > 0 {
		allServiceNames = append(allServiceNames, serviceNames...)
	}
	return allServiceNames
}

func init() {
	cmd_action.GetActionCmd().AddCommand(serviceCmd)

	serviceCmd.PersistentFlags().StringSliceVar(&serviceNames, "serviceNames", []string{}, "Service names that has to be acted upon")
	serviceCmd.PersistentFlags().StringVar(&serviceName, "serviceName", "", "Service names that has to be acted upon")

	serviceCmd.MarkFlagsOneRequired("serviceNames", "serviceName")
	serviceCmd.MarkFlagsOneRequired("serviceNames", "serviceName")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
