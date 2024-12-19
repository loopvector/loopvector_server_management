package helper

const KInventoryFilePath = "ansible"
const KInventoryFileName = "inventory.ini"
const KInventoryDefaultServerTag = "all"
const KInventoryDefaultBecome = true

const KPlaybookFileName = "playbook.yaml"
const KPlaybookFilePath = "ansible"

const KFullPathTaskReboot = "task/maintain/power/reboot.yaml"
const KFullPathTaskInstallApps = "task/maintain/package/install_packages.yaml"
const KFullPathTaskPing = "task/maintain/access/ping.yaml"
const KFullPathTaskPackageUpdate = "task/maintain/package/packages_update.yaml"
const KFullPathTaskPackageUpgrade = "task/maintain/package/packages_upgrade.yaml"
const KFullPathTaskPackageAutoRemove = "task/maintain/package/packages_auto_remove.yaml"
const KSshCommonArgs = "-o StrictHostKeyChecking=no"
