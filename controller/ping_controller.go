package controller

import (
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
)

// type AnsibleTaskPing struct {
// 	//BaseAnsibleTask
// 	// CustomInputName string   `yaml:"custom_input_name"`
// 	// SomeOtherVar1   string   `yaml:"some_other_var1"`
// 	// BooleanVar1     bool     `yaml:"boolean_var1"`
// 	// ArrayInput1     []string `yaml:"array_input1"`
// }

func PingServer(
	serverName model.ServerNameModel,
	serverSshConnectionInfo model.ServerSshConnectionInfo,
) {
	println("Pinging server: ", serverName.Name)

	RunAnsibleTasks(
		serverName,
		serverSshConnectionInfo,
		[]model.AnsibleTask{{FullPath: helper.KFullPathTaskPing}},
		nil,
	)

	// ansibleTaskPing := AnsibleTaskPing{
	// 	BaseAnsibleTask: BaseAnsibleTask{
	// 		TaskFullPath: helper.KFullPathTaskPing,
	// 	},
	// }

	// playbook := AnsiblePlaybook{
	// 	HostsTag: helper.KInventoryDefaultServerTag,
	// 	Become:   helper.KInventoryDefaultBecome,
	// }
	// playbook.AddTask(&ansibleTaskPing)

	// createPlaybook := GenerateAnsiblePlaybook{
	// 	PlaybookFilePath:    helper.KPlaybookFilePath,
	// 	PlaybookFileName:    helper.KPlaybookFileName,
	// 	AnsiblePlaybookList: []AnsiblePlaybook{playbook},
	// }

	// createPlaybook.GeneratePlaybook()

	// inventoryFileFullPath, err := helper.GetFullFilePath(inventoryFilePath, inventoryFileName)
	// if err != nil {
	// 	panic(err)
	// }
	// playbookFileFullPath, err := helper.GetFullFilePath(playbookFilePath, playbookFileName)
	// if err != nil {
	// 	panic(err)
	// }

	// runPlaybook := RunAnsiblePlaybook{
	// 	InventoryFileFullPath: inventoryFileFullPath,
	// 	PlaybookFileFullPath:  playbookFileFullPath,
	// }

	// runPlaybook.RunPlaybook()

}
