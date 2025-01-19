package helper

const KInventoryFilePath = "ansible"
const KInventoryFileName = "inventory.ini"
const KInventoryDefaultServerTag = "all"
const KInventoryDefaultBecome = true

const KPlaybookFileName = "playbook.yaml"
const KPlaybookFilePath = "ansible"

const KFullPathTaskServiceEnable = "task/configure/service/enable_services.yaml"
const KFullPathTaskServiceStart = "task/configure/service/start_services.yaml"
const KFullPathTaskServiceRestart = "task/configure/service/restart_services.yaml"

const KFullPathTaskUfwEnable = "task/configure/security/enable_ufw.yaml"
const KFullPathTaskUfwDisable = "task/configure/security/disable_ufw.yaml"

const KFullPathTaskUfwDefaultIncomingConfigure = "task/configure/security/ufw_default_incoming_traffic_policy.yaml"
const KFullPathTaskUfwDefaultOutgoingConfigure = "task/configure/security/ufw_default_outgoing_traffic_policy.yaml"

const KFullPathTaskUfwPortsConfigure = "task/configure/security/ports_ufw_traffic_policy.yaml"
const KFullPathTaskUfwIpAddressesConfigure = "task/configure/security/ip_addresses_ufw_traffic_policy.yaml"
const KFullPathTaskUfwIpAddressesWithPortConfigure = "task/configure/security/ip_addresses_with_port_ufw_traffic_policy.yaml"

const KFullPathTaskAddUsersToGroups = "task/configure/user/add_users_to_groups.yaml"

const KFullPathTaskChangeRootPassword = "task/configure/security/change_root_password.yaml"

//const KFullPathTaskAddUserToGroup = "task/configure/user/add_user_to_groups.yaml"

const KFullPathTaskAddUsers = "task/configure/user/add_users.yaml"
const KFullPathTaskAddUser = "task/configure/user/add_user.yaml"

const KFullPathTaskAddLinesToFile = "task/configure/file/add_lines.yaml"
const KFullPathTaskAddLineBlockToFile = "task/configure/file/add_line_block.yaml"
const KFullPathTaskReadBlocks = "task/configure/file/read_blocks.yaml"
const KFullPathTaskDeleteBlock = "task/configure/file/delete_block.yaml"

const KFullPathTaskAddGroups = "task/configure/group/add_groups.yaml"
const KFullPathTaskAddGroup = "task/configure/group/add_group.yaml"

const KFullPathTaskReboot = "task/maintain/power/reboot.yaml"

const KFullPathTaskInstallApps = "task/maintain/package/install_packages.yaml"
const KFullPathTaskInstallApp = "task/maintain/package/install_package.yaml"

const KFullPathTaskPing = "task/maintain/access/ping.yaml"

const KFullPathTaskPackageUpdate = "task/maintain/package/packages_update.yaml"

const KFullPathTaskPackageUpgrade = "task/maintain/package/packages_upgrade.yaml"

const KFullPathTaskPackageAutoRemove = "task/maintain/package/packages_auto_remove.yaml"

const KSshCommonArgs = "-o StrictHostKeyChecking=no"
