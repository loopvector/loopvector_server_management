/*
Copyright © 2024 Agilan Anandan (agilan@loopvector.com)
*/
package main

import (
	"loopvector_server_management/cmd"
	"loopvector_server_management/model"

	_ "loopvector_server_management/cmd/cmd_action"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_group"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_group/cmd_action_group_add"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_install"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_package"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_package/cmd_action_package_auto_remove"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_package/cmd_action_package_update"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_package/cmd_action_package_upgrade"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_ping"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_reboot"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_sshd"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_sshd/cmd_action_sshd_add"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_sudoers"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_sudoers/cmd_action_sudoers_add"
	_ "loopvector_server_management/cmd/cmd_delete"
	_ "loopvector_server_management/cmd/cmd_delete/cmd_delete_server"
	_ "loopvector_server_management/cmd/cmd_list"
	_ "loopvector_server_management/cmd/cmd_list/cmd_list_server"
	_ "loopvector_server_management/cmd/create"
	_ "loopvector_server_management/cmd/create/server"
	_ "loopvector_server_management/cmd/migrate"
	_ "loopvector_server_management/cmd/migrate/database"
)

func main() {
	model.InitializeDB(false)
	cmd.Execute()
}
