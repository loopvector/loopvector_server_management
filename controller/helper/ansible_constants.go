package helper

const KInventoryFilePath = "ansible"
const KInventoryFileName = "inventory.ini"
const KInventoryDefaultServerTag = "all"
const KInventoryDefaultBecome = true

const KPlaybookFileName = "playbook.yaml"
const KPlaybookFilePath = "ansible"

const KFullPathTaskAddUsers = "task/configure/user/add_users.yaml"
const KFullPathTaskAddUser = "task/configure/user/add_user.yaml"

const KFullPathTaskAddLinesToFile = "task/configure/file/add_lines.yaml"

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
