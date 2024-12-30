package model

import (
	"encoding/json"
	"fmt"
	"loopvector_server_management/controller/helper"

	"gopkg.in/yaml.v3"
)

type AnsiblePlaybookFile struct {
	AnsibleTasks []AnsibleTask
}

// type AnsiblePlaybookFileWithVarsArray struct {
// 	AnsibleTasks []AnsibleTaskWithVarsArray
// }

type AnsibleTask struct {
	FullPath string
	Vars     map[string]interface{}
	//VarsArray []map[string]interface{}
}

// type AnsibleTaskWithVarsArray struct {
// 	FullPath string
// 	Vars     []map[string]interface{}
// }

type _AnsiblePlaybookFile struct {
	HostsTag string         `yaml:"hosts"`
	Become   bool           `yaml:"become"`
	Tasks    []_AnsibleTask `yaml:"tasks"`
	FileName string         `yaml:"-"`
	FilePath string         `yaml:"-"`
}

type AnsiblePlaybookRunner struct {
	InventoryFileFullPath string
	PlaybookFileFullPath  string
}

type BaseAnsibleTask struct {
	TaskFullPath string `yaml:"include_tasks"`
}

// type _AnsibleTask interface {
// 	GetTaskFullPath() string
// }

type _AnsibleTask struct {
	TaskFullPath string                 `yaml:"include_tasks"`
	Vars         map[string]interface{} `yaml:"vars,omitempty"`
	// VarsArray    []map[string]interface{} `yaml:"vars,omitempty"`
}

// type _AnsibleTaskWithVarsArray struct {
// 	TaskFullPath string                   `yaml:"include_tasks"`
// 	Vars         []map[string]interface{} `yaml:"vars,omitempty"`
// }

type AnsiblePlaybookRunResult struct {
	Stats map[string]interface{} `json:"stats"`
	Plays []struct {
		Tasks []struct {
			Task struct {
				Duration struct {
					End   string `json:"end"`
					Start string `json:"start"`
				} `json:"duration"`
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"task"`
			Hosts map[string]struct {
				Changed bool    `json:"changed"`
				Failed  *bool   `json:"failed,omitempty"`
				Msg     *string `json:"msg,omitempty"`
			} `json:"hosts"`
		} `json:"tasks"`
	} `json:"plays"`
}

// type _AnsibleTaskVar struct {
// 	Vars map[string]interface{} `yaml:"vars,omitempty"`
// }

func (t *BaseAnsibleTask) GetTaskFullPath() string {
	return t.TaskFullPath
}

func (p *_AnsiblePlaybookFile) _AddTask(ansibleTask AnsibleTask) {
	// println("Adding task with vars: ", ansibleTask.Vars)
	p.Tasks = append(p.Tasks, _AnsibleTask{
		TaskFullPath: ansibleTask.FullPath,
		Vars:         ansibleTask.Vars,
		// VarsArray:    ansibleTask.VarsArray,
	})
}

func (f AnsiblePlaybookFile) CreateNew() (AnsiblePlaybookRunner, error) {
	ansiblePlaybookFile := _AnsiblePlaybookFile{
		HostsTag: helper.KInventoryDefaultServerTag,
		Become:   helper.KInventoryDefaultBecome,
		FileName: helper.KPlaybookFileName,
		FilePath: helper.KPlaybookFilePath,
	}
	for _, task := range f.AnsibleTasks {
		ansiblePlaybookFile._AddTask(task)
	}

	ansiblePlaybookList := []_AnsiblePlaybookFile{ansiblePlaybookFile}

	data, err := yaml.Marshal(&ansiblePlaybookList)
	if err != nil {
		panic(err)
	}
	finalYAML := "---\n" + string(data)
	helper.WriteToFile(ansiblePlaybookFile.FilePath, ansiblePlaybookFile.FileName, finalYAML)

	inventoryFileFullPath, err := helper.GetFullFilePath(helper.KInventoryFilePath, helper.KInventoryFileName)
	if err != nil {
		panic(err)
	}
	playbookFileFullPath, err := helper.GetFullFilePath(helper.KPlaybookFilePath, helper.KPlaybookFileName)
	if err != nil {
		panic(err)
	}

	return AnsiblePlaybookRunner{
		InventoryFileFullPath: inventoryFileFullPath,
		PlaybookFileFullPath:  playbookFileFullPath,
	}, nil
}

func (f AnsiblePlaybookRunner) Run() (AnsiblePlaybookRunResult, error) {
	resultMap, err := helper.RunAnsiblePlaybook(f.InventoryFileFullPath, f.PlaybookFileFullPath)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.Marshal(resultMap)
	if err != nil {
		return AnsiblePlaybookRunResult{}, fmt.Errorf("error marshaling map to JSON: %w", err)
	}

	// Unmarshal into the struct
	var result AnsiblePlaybookRunResult
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return AnsiblePlaybookRunResult{}, fmt.Errorf("error unmarshaling JSON to struct: %w", err)
	}

	//fmt.Print("task result: %v", result)
	return result, nil
}

// func (t _AnsibleTask) MarshalYAML() (interface{}, error) {
// 	// Start with the base structure
// 	data := map[string]interface{}{
// 		"include_tasks": t.TaskFullPath,
// 	}

// 	// Use reflection to dynamically extract additional fields
// 	vars := map[string]interface{}{}
// 	tValue := reflect.ValueOf(t)
// 	tType := reflect.TypeOf(t)

// 	for i := 0; i < tValue.NumField(); i++ {
// 		field := tType.Field(i)
// 		yamlTag := field.Tag.Get("yaml")

// 		// Check for `omitempty` in the tag
// 		if yamlTag == "" || yamlTag == "include_tasks" {
// 			continue
// 		}

// 		// Split yamlTag into components (e.g., "field_name,omitempty")
// 		tagParts := helper.SplitYAMLTag(yamlTag)

// 		// Get field value
// 		fieldValue := tValue.Field(i)

// 		// Apply omitempty check
// 		if len(tagParts) > 1 && tagParts[1] == "omitempty" && helper.IsEmptyValue(fieldValue) {
// 			continue
// 		}

// 		// Add to vars map
// 		vars[tagParts[0]] = fieldValue.Interface()
// 	}

// 	if len(vars) > 0 {
// 		data["vars"] = vars
// 	}

// 	return data, nil
// }
