package controller

import (
	"fmt"
	"loopvector_server_management/controller/helper"

	"gopkg.in/yaml.v3"
)

type AnsiblePlaybook struct {
	HostsTag string        `yaml:"hosts"`
	Become   bool          `yaml:"become"`
	Tasks    []interface{} `yaml:"tasks"`
}

type GenerateAnsiblePlaybook struct {
	PlaybookFilePath    string
	PlaybookFileName    string
	AnsiblePlaybookList []AnsiblePlaybook
}

type RunAnsiblePlaybook struct {
	InventoryFileFullPath string
	PlaybookFileFullPath  string
}

type AnsibleTaskResult struct {
	Stats map[string]interface{} `json:"stats"`
	Plays []struct {
		Tasks []struct {
			TaskName string `json:"task"`
			Hosts    map[string]struct {
				Changed bool   `json:"changed"`
				Failed  bool   `json:"failed"`
				Msg     string `json:"msg,omitempty"`
			} `json:"hosts"`
		} `json:"tasks"`
	} `json:"plays"`
}

// func (p *AnsiblePlaybook) AddTask(task AnsibleTask) {
// 	p.Tasks = append(p.Tasks, task)
// }

// type CreateAnsiblePlaybookRequest struct {
// 	HostsTag             string
// 	Become               bool
// 	InventoryFilePath    string
// 	InventoryFileName    string
// 	PlaybookTaskFullPath []string
// }

func (r *GenerateAnsiblePlaybook) GeneratePlaybook() {
	data, err := yaml.Marshal(&r.AnsiblePlaybookList)
	if err != nil {
		panic(err)
	}
	finalYAML := "---\n" + string(data)
	helper.WriteToFile(r.PlaybookFilePath, r.PlaybookFileName, finalYAML)
}

func (r *RunAnsiblePlaybook) RunPlaybook() {
	result, err := helper.RunAnsiblePlaybook(r.InventoryFileFullPath, r.PlaybookFileFullPath)
	if err != nil {
		panic(err)
	}

	fmt.Print("task result: %v", result)
}
