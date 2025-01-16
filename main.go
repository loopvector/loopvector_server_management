/*
Copyright © 2024 Agilan Anandan (agilan@loopvector.com)
*/
package main

import (
	"log"
	"loopvector_server_management/cmd"
	"loopvector_server_management/helper"
	"loopvector_server_management/model"
	"strings"

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
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_root_user"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_root_user/cmd_action_root_user_update_password"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_service"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_service/cmd_action_service_enable"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_service/cmd_action_service_restart"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_service/cmd_action_service_start"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_ssh_pub_key"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_ssh_pub_key/cmd_action_ssh_pub_key_add"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_sshd"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_sshd/cmd_action_sshd_add"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_sudoers"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_sudoers/cmd_action_sudoers_add"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_ufw"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_ufw/cmd_action_ufw_allow"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_ufw/cmd_action_ufw_default_incoming"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_ufw/cmd_action_ufw_default_outgoing"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_ufw/cmd_action_ufw_deny"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_ufw/cmd_action_ufw_disable"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_ufw/cmd_action_ufw_enable"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_user"
	_ "loopvector_server_management/cmd/cmd_action/cmd_action_user/cmd_action_user_add"
	_ "loopvector_server_management/cmd/cmd_delete"
	_ "loopvector_server_management/cmd/cmd_delete/cmd_delete_server"
	_ "loopvector_server_management/cmd/cmd_list"
	_ "loopvector_server_management/cmd/cmd_list/cmd_list_server"
	_ "loopvector_server_management/cmd/create"
	_ "loopvector_server_management/cmd/create/server"
	_ "loopvector_server_management/cmd/create/user"
	_ "loopvector_server_management/cmd/database"
	_ "loopvector_server_management/cmd/database/migrate"
	_ "loopvector_server_management/cmd/user"
	_ "loopvector_server_management/cmd/user/login"
	_ "loopvector_server_management/cmd/user/logout"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Proceeding with environment variables.")
	}

	viper.AutomaticEnv()

	shouldInitializeDb := model.AppFirstLaunch{}.CheckFirstLaunch()
	model.InitializeDB(shouldInitializeDb)
	cmd.Execute()

	model.GenerateAdminSetting(model.AdminConfig{
		SMTPHost:                      viper.GetString(helper.KSmtpHost),
		SMTPPort:                      viper.GetUint16(helper.KSmtpPort),
		SMTPUser:                      viper.GetString(helper.KSmtpUser),
		SMTPPassword:                  viper.GetString(helper.KSmtpPassword),
		SignupDomainWhitelist:         strings.Split(viper.GetString(helper.KSignUpDomainWhitelist), ","),
		UserEmailVerificationRequired: viper.GetBool(helper.KUserEmailVerificationRequired),
	})
}
