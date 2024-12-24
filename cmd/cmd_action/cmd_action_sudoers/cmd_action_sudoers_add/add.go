/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_sudoers_add

import (
	"errors"
	"fmt"
	"loopvector_server_management/cmd/cmd_action/cmd_action_sudoers"
	"loopvector_server_management/controller"
	"path/filepath"

	"github.com/spf13/cobra"
)

const sudoersFileCreatePath = "/etc/sudoers"

var (
	fileName           string
	filePermissionMode = "0440"
	groupName          string
	host               string
	runAsUser          string
	runAsGroup         string
	password           string
	command            string
	line               string
	lines              []string
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
		suggestions, err := controller.GetAllActiveServerNames()
		if err != nil {
			fmt.Println("Error querying database:", err)
			return nil, cobra.ShellCompDirectiveError
		}

		return suggestions, cobra.ShellCompDirectiveNoFileComp
	},
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		hasCombination1 := cmd.Flags().Changed("groupName") ||
			cmd.Flags().Changed("host") ||
			cmd.Flags().Changed("runAsUser") ||
			cmd.Flags().Changed("runAsGroup") ||
			cmd.Flags().Changed("password") ||
			cmd.Flags().Changed("command")

		hasCombination2 := cmd.Flags().Changed("line")

		hasCombination3 := cmd.Flags().Changed("lines")

		if (hasCombination1 && hasCombination2) ||
			(hasCombination1 && hasCombination3) ||
			(hasCombination2 && hasCombination3) {
			return errors.New("only one combination of flags can be used at a time")
		}

		serverName := args[0]

		fileFullPath := filepath.Join(sudoersFileCreatePath, fileName)

		if hasCombination1 {
			if fileName == "" || groupName == "" {
				return errors.New("flags --fileName and --groupName are required for this combination")
			}
			controller.AddASudoer(
				serverName,
				controller.AddLinesToFileRequest{
					FileFullPath:   fileFullPath,
					FilePermission: filePermissionMode,
					AsSudo:         true,
				},
				controller.SudoersAddRequest{
					GroupName:  groupName,
					Host:       host,
					RunAsUser:  runAsUser,
					RunAsGroup: runAsGroup,
					Password:   password,
					Command:    command,
				},
			)
			//fmt.Println("Combination 1 selected")
		} else if hasCombination2 {
			if fileName == "" || line == "" {
				return errors.New("flags --fileName and --line are required for this combination")
			}
			controller.AddLinesToFile(
				serverName,
				controller.AddLinesToFileRequest{
					FileFullPath:   fileFullPath,
					FilePermission: filePermissionMode,
					AsSudo:         true,
				},
				[]controller.LineToFileAddRequest{{
					Line: line,
				}},
			)
			//fmt.Println("Combination 2 selected")
		} else if hasCombination3 {
			if fileName == "" || len(lines) == 0 {
				return errors.New("flags --fileName and --lines are required for this combination")
			}
			linesRequest := []controller.LineToFileAddRequest{}
			for _, line := range lines {
				linesRequest = append(linesRequest, controller.LineToFileAddRequest{
					Line: line,
				})
			}
			controller.AddLinesToFile(
				serverName,
				controller.AddLinesToFileRequest{
					FileFullPath:   fileFullPath,
					FilePermission: filePermissionMode,
					AsSudo:         true,
				},
				linesRequest,
			)
			//fmt.Println("Combination 3 selected")
		} else {
			return errors.New("no valid combination of flags provided")
		}

		return nil
	},
}

func init() {
	cmd_action_sudoers.GetActionSudoersCmd().AddCommand(addCmd)

	addCmd.Flags().StringVar(&fileName, "fileName", "", "Custom sudoers file name that will be created at /etc/sudoers.d/ directory")
	//addCmd.Flags().StringVar(&filePermissionMode, "filePermissionMode", "", "File permission mode (e.g., 0644, 0755)")

	addCmd.MarkFlagRequired("fileName")
	//addCmd.MarkFlagRequired("filePermissionMode")

	addCmd.Flags().StringVar(&groupName, "groupName", "", "Group name (e.g., g_admin, sadmin)")
	addCmd.Flags().StringVar(&host, "host", "ALL", "Host (e.g., ALL, specific host)")
	addCmd.Flags().StringVar(&runAsUser, "runAsUser", "ALL", "RunAs user (e.g., ALL, root)")
	addCmd.Flags().StringVar(&runAsGroup, "runAsGroup", "", "RunAs group (e.g., ALL, empty if not needed)")
	addCmd.Flags().StringVar(&password, "password", "", "Password requirement (e.g., ALL, NOPASSWD, empty defaults to ALL)")
	addCmd.Flags().StringVar(&command, "command", "ALL", "Command(s) allowed to run (e.g., ALL, specific commands)")

	addCmd.Flags().StringVar(&line, "line", "", "Sudoers line (e.g. %g_admin ALL=(ALL) ALL)")

	addCmd.Flags().StringSliceVar(&lines, "lines", []string{}, "Sudoers lines (e.g. %g_admin ALL=(ALL) ALL,%sadmin ALL=(ALL:ALL) NOPASSWD:ALL)")

}
