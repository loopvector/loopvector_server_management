package controller

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"os/exec"
// )

// Define BaseAnsibleTask struct


// // type RunAnsibleTaskRequest struct {
// // 	HostsTag              string
// // 	Become                bool
// // 	InventoryFullFilePath string
// // 	PlaybookFullFilePath  string
// // }

// type AnsibleTaskResult struct {
// 	Stats map[string]interface{} `json:"stats"`
// 	Plays []struct {
// 		Tasks []struct {
// 			TaskName string `json:"task"`
// 			Hosts    map[string]struct {
// 				Changed bool   `json:"changed"`
// 				Failed  bool   `json:"failed"`
// 				Msg     string `json:"msg,omitempty"`
// 			} `json:"hosts"`
// 		} `json:"tasks"`
// 	} `json:"plays"`
// }

// func (r *RunAnsibleTaskRequest) RunAnsibleTask() (*AnsibleTaskResult, error) {
// 	cmd := exec.Command("docker", "run", "--rm",
// 		"-v", "C:\\ansible:/ansible", // Adjust based on your environment
// 		"-w", "/ansible",
// 		"ansible-image", // Replace with your actual Ansible Docker image
// 		"ansible-playbook", "-i", r.InventoryFullFilePath, r.PlaybookFullFilePath, "--json")

// 	// Capture JSON output
// 	var out bytes.Buffer
// 	var stderr bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = &stderr

// 	// Run the command
// 	err := cmd.Run()
// 	if err != nil {
// 		return nil, fmt.Errorf("ansible-playbook failed: %v, stderr: %s", err, stderr.String())
// 	}

// 	// Parse JSON output
// 	var result AnsibleTaskResult
// 	err = json.Unmarshal(out.Bytes(), &result)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse Ansible output: %v", err)
// 	}

// 	return &result, nil
// }
