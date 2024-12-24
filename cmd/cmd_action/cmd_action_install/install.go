/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_install

import (
	"fmt"
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/controller"
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"

	"github.com/spf13/cobra"
)

var (
	appsToInstallList []string
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		suggestions, err := controller.GetAllActiveServerNames()
		if err != nil {
			fmt.Println("Error querying database:", err)
			return nil, cobra.ShellCompDirectiveError
		}

		return suggestions, cobra.ShellCompDirectiveNoFileComp
	},
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		// println("appsToInstallList: ", appsToInstallList)
		// Task to install the list of apps at once
		vars := map[string]interface{}{
			"package_names": appsToInstallList,
		}

		callbacks := []controller.RunAnsibleTaskCallback{}

		serverId, err := model.ServerNameModel{Name: args[0]}.GetServerIdUsingServerName()
		if err != nil {
			panic(err)
		}

		for _, app := range appsToInstallList {
			callbacks = append(
				callbacks,
				controller.RunAnsibleTaskCallback{
					TaskName: "install package " + app,
					OnChanged: func() {
						// println("Installed ", app, " on ", args[0])
						model.ServerApp{
							ServerID: serverId,
							Name:     app,
						}.RegisterInstall()
					},
					OnUnchanged: func() {
						// println("Already installed ", app, " on ", args[0])
						model.ServerApp{
							ServerID: serverId,
							Name:     app,
						}.RegisterInstall()
					},
					OnFailed: func() {
						// println("Failed to install ", app, " on ", args[0])
					},
				},
			)
		}

		_, err = controller.RunAnsibleTasks(
			model.ServerNameModel{Name: args[0]},
			[]model.AnsibleTask{{FullPath: helper.KFullPathTaskInstallApps, Vars: vars}},
			callbacks,
		)

		if err != nil {
			panic(err)
		}

		// for _, play := range result.Plays {
		// 	for _, taskResult := range play.Tasks {
		// 		for _, appToInstall := range appsToInstallList {
		// 			println("About to check install status of ", appToInstall, " on ", args[0], " in task name ", taskResult.Task.Name)
		// 			if taskResult.Task.Name == "install package "+appToInstall {
		// 				if taskResult.Hosts["name="+args[0]].Failed == nil {
		// 					if taskResult.Hosts["name="+args[0]].Changed {
		// 						println("Installed ", appToInstall, " on ", args[0])
		// 					} else {
		// 						println("Already installed ", appToInstall, " on ", args[0])
		// 					}
		// 				} else {
		// 					println("Failed to install ", appToInstall, " on ", args[0])
		// 				}
		// 			}
		// 		}
		// 	}
		// }

		//fmt.Print("task result: %v", result)

		// Task to install the list of apps on by one
		// for _, app := range appsToInstallList {
		// 	vars := map[string]interface{}{
		// 		"package_name": app,
		// 	}
		// 	controller.RunAnsibleTasks(
		// 		model.ServerNameModel{Name: args[0]},
		// 		[]model.AnsibleTask{{FullPath: helper.KFullPathTaskInstallApp, Vars: vars}},
		// 	)
		// }

	},
}

func init() {
	cmd_action.GetActionCmd().AddCommand(installCmd)

	installCmd.Flags().StringSliceVar(&appsToInstallList, "apps", []string{}, "install the list of apps to the server")

	installCmd.MarkFlagRequired("apps")
}
