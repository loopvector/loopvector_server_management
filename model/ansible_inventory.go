package model

import "loopvector_server_management/controller/helper"

// type AnsibleInventoryFileRootUserIpv4 struct {
// 	ServerName     string
// 	ServerIpv4     ServerIpv4
// 	ServerRootUser ServerRootUser
// }

// type AnsibleInventoryFileGenericDetails struct {
// 	ServerName            ServerNameModel
// 	ServerIpv4            *ServerIpv4
// 	ServerIpv6            *ServerIpv6
// 	ServerRootUser        *ServerRootUser
// 	ServerUser            *ServerUser
// 	SshPvtKeyFullFilePath *string
// 	Port                  *uint16
// }

type ServerSshConnectionInfo struct {
	ServerName            string
	Username              string
	Password              string
	Port                  uint16
	Ip                    string
	SshPvtKeyFullFilePath string
}

type _AnsibleInventoryFile struct {
	ServerName           string  `inventory:"name"`
	Ip                   string  `inventory:"ansible_host"`
	Port                 uint16  `inventory:"ansible_port,omitEmpty"`
	Username             string  `inventory:"ansible_user,omitEmpty"`
	UserPassword         string  `inventory:"ansible_password,omitEmpty"`
	BecomePassword       string  `inventory:"ansible_become_password,omitEmpty"`
	UserSSHKey           string `inventory:"ansible_private_key_file,omitEmpty"`
	AnsibleSshCommonArgs string  `inventory:"ansible_ssh_common_args,omitEmpty"`
	FilePath             string
	FileName             string
}

func (f ServerSshConnectionInfo) CreateNew() error {
	println("Creating inventory file for serverName", f.ServerName)
	println("Ip: ", f.Ip)
	println("Port: ", f.Port)
	println("Username: ", f.Username)
	println("UserPassword: ", f.Password)
	println("BecomePassword: ", f.Password)
	println("UserSSHKey: ", f.SshPvtKeyFullFilePath)
	ansibleInventoryFile := _AnsibleInventoryFile{
		ServerName:           f.ServerName,
		Ip:                   f.Ip,
		Port:                 f.Port,
		Username:             f.Username,
		UserPassword:         f.Password,
		BecomePassword:       f.Password,
		UserSSHKey:           f.SshPvtKeyFullFilePath,
		AnsibleSshCommonArgs: helper.KSshCommonArgs,
		FilePath:             helper.KInventoryFilePath,
		FileName:             helper.KInventoryFileName,
	}

	return _generateInventoryFile(ansibleInventoryFile)
}

// func (f AnsibleInventoryFileRootUserIpv4) CreateNewUsingRootUserAndIpv4() error {
// 	ansibleInventoryFile := _AnsibleInventoryFile{
// 		ServerName:           f.ServerName,
// 		Ip:                   f.ServerIpv4.Ip,
// 		Port:                 f.ServerRootUser.Port,
// 		Username:             helper.KRootUserUsername,
// 		UserPassword:         f.ServerRootUser.Password,
// 		BecomePassword:       f.ServerRootUser.Password,
// 		UserSSHKey:           f.ServerRootUser.SshKey,
// 		AnsibleSshCommonArgs: helper.KSshCommonArgs,
// 		FilePath:             helper.KInventoryFilePath,
// 		FileName:             helper.KInventoryFileName,
// 	}

// 	return _generateInventoryFile(ansibleInventoryFile)
// }

func _generateInventoryFile(ansibleInventoryFile _AnsibleInventoryFile) error {
	inventoryContent, err := helper.GenerateInventoryFileContent([]interface{}{ansibleInventoryFile})
	if err != nil {
		return err
	}

	err = helper.WriteToFile(ansibleInventoryFile.FilePath, ansibleInventoryFile.FileName, inventoryContent)
	if err != nil {
		return err
	}
	return nil
}
