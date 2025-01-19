/*
Copyright © 2024 Agilan Anandan <agilan@loopvector.com>
*/
package cmd_action_ssh_pub_key_add

import (
	"errors"
	"loopvector_server_management/cmd/cmd_action/cmd_action_ssh_pub_key"

	"github.com/spf13/cobra"
)

var (
	key  string
	keys []string
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
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if key == "" && len(keys) == 0 {
			return errors.New("no key(s) provided")
		}
		allKeys := []string{}
		if len(keys) > 0 {
			allKeys = append(allKeys, keys...)
		} else {
			allKeys = append(allKeys, key)
		}
		return cmd_action_ssh_pub_key.AddSshKeys(allKeys)
	},
}

func init() {
	cmd_action_ssh_pub_key.GetActionSshPubKeyCmd().AddCommand(addCmd)

	addCmd.PersistentFlags().StringVar(&key, "key", "", "pub key that can be used to ssh into the server")
	addCmd.PersistentFlags().StringSliceVar(&keys, "keys", []string{}, "pub keys that can be used to ssh into the server")

	addCmd.MarkFlagsOneRequired("key", "keys")
	addCmd.MarkFlagsMutuallyExclusive("key", "keys")
}
