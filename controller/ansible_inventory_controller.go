package controller

// type AnsibleInventoryFileCreateRequest struct {
// 	ServerName           string  `inventory:"name"`
// 	Ip                   string  `inventory:"ansible_host"`
// 	Port                 uint16  `inventory:"ansible_port,omitEmpty"`
// 	Username             string  `inventory:"ansible_user,omitEmpty"`
// 	UserPassword         string  `inventory:"ansible_password,omitEmpty"`
// 	BecomePassword       string  `inventory:"ansible_become_password,omitEmpty"`
// 	UserSSHKey           *string `inventory:"ansible_private_key_file,omitEmpty"`
// 	AnsibleSshCommonArgs string  `inventory:"ansible_ssh_common_args,omitEmpty"`
// 	FilePath             string
// 	FileName             string
// }

// func (r *AnsibleInventoryFileCreateRequest) CreateNewFile() error {
// 	inventoryContent, err := helper.GenerateInventoryFileContent([]interface{}{r})
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = helper.WriteToFile(r.FilePath, r.FileName, inventoryContent)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return nil
// }
