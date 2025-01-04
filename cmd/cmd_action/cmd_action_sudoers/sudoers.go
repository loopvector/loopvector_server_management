/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_sudoers

import (
	"errors"
	"fmt"
	"loopvector_server_management/cmd/cmd_action"
	"loopvector_server_management/controller"
	"path/filepath"

	"github.com/spf13/cobra"
)

const sudoersFileCreatePath = "/etc/sudoers.d"

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

func GetActionSudoersCmd() *cobra.Command {
	return sudoersCmd
}

// sudoersCmd represents the sudoers command
var sudoersCmd = &cobra.Command{
	Use:   "sudoers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sudoers called")
	},
}

func AddSudoers(cmd *cobra.Command) error {
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

	serverName := cmd_action.GetServerName()

	fileFullPath := filepath.Join(sudoersFileCreatePath, fileName)

	if hasCombination1 {
		if fileName == "" || groupName == "" {
			return errors.New("flags --fileName and --groupName are required for this combination")
		}
		controller.AddASudoer(
			serverName,
			cmd_action.GetServerSshConnectionInfo(),
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
		controller.AddLineBlockToFile(
			serverName,
			cmd_action.GetServerSshConnectionInfo(),
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
		// linesRequest := []controller.LineToFileAddRequest{}
		for _, line := range lines {
			oneLine := controller.LineToFileAddRequest{
				Line: line,
			}
			controller.AddLineBlockToFile(
				serverName,
				cmd_action.GetServerSshConnectionInfo(),
				controller.AddLinesToFileRequest{
					FileFullPath:   fileFullPath,
					FilePermission: filePermissionMode,
					AsSudo:         true,
				},
				[]controller.LineToFileAddRequest{oneLine},
			)
		}

		//fmt.Println("Combination 3 selected")
	} else {
		return errors.New("no valid combination of flags provided")
	}

	return nil
}

func init() {
	cmd_action.GetActionCmd().AddCommand(sudoersCmd)

	sudoersCmd.PersistentFlags().StringVar(&fileName, "fileName", "", "Custom sudoers file name that will be created at /etc/sudoers.d/ directory")
	//addCmd.Flags().StringVar(&filePermissionMode, "filePermissionMode", "", "File permission mode (e.g., 0644, 0755)")

	sudoersCmd.MarkFlagRequired("fileName")
	//addCmd.MarkFlagRequired("filePermissionMode")

	sudoersCmd.PersistentFlags().StringVar(&groupName, "groupName", "", "Group name (e.g., g_admin, sadmin)")
	sudoersCmd.PersistentFlags().StringVar(&host, "host", "ALL", "Host (e.g., ALL, specific host)")
	sudoersCmd.PersistentFlags().StringVar(&runAsUser, "runAsUser", "ALL", "RunAs user (e.g., ALL, root)")
	sudoersCmd.PersistentFlags().StringVar(&runAsGroup, "runAsGroup", "", "RunAs group (e.g., ALL, empty if not needed)")
	sudoersCmd.PersistentFlags().StringVar(&password, "password", "", "Password requirement (e.g., ALL, NOPASSWD, empty defaults to ALL)")
	sudoersCmd.PersistentFlags().StringVar(&command, "command", "ALL", "Command(s) allowed to run (e.g., ALL, specific commands)")

	sudoersCmd.PersistentFlags().StringVar(&line, "line", "", "Sudoers line (e.g. %g_admin ALL=(ALL) ALL)")

	sudoersCmd.PersistentFlags().StringSliceVar(&lines, "lines", []string{}, "Sudoers lines (e.g. %g_admin ALL=(ALL) ALL,%sadmin ALL=(ALL:ALL) NOPASSWD:ALL)")
}
