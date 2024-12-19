package helper

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func RunAnsiblePlaybook(inventoryPath, playbookPath string) (map[string]interface{}, error) {
	// Ansible command with JSON output
	cmd := exec.Command(
		"ansible-playbook",
		"-i", inventoryPath,
		playbookPath,
	)

	cmd.Env = append(os.Environ(), "ANSIBLE_CONFIG=ansible/ansible.cfg")

	println("command: " + cmd.String())

	// Capture the output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error running ansible-playbook: %s", err)
	}

	// Parse JSON output
	// var result map[string]interface{}
	// if err := json.Unmarshal(stdout.Bytes(), &result); err != nil {
	// 	return nil, fmt.Errorf("error parsing JSON output: %s", err)
	// }

	return nil, nil
}

func GenerateInventoryFileContent(input []interface{}) (string, error) {
	var inventoryLines []string

	inventoryLines = append(inventoryLines, "[all]")
	for _, item := range input {
		// Ensure the item is a struct or a pointer to a struct
		val := reflect.ValueOf(item)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		if val.Kind() != reflect.Struct {
			return "", fmt.Errorf("each item in the input must be a struct")
		}

		typ := val.Type()
		var serverLines []string

		// Iterate through the struct fields
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			tag := field.Tag.Get("inventory")
			if tag == "" {
				continue // Skip if no inventory tag
			}

			tags := strings.Split(tag, ",")
			key := tags[0]
			omitEmpty := len(tags) > 1 && tags[1] == "omitEmpty"

			// Get the field value
			fieldValue := val.Field(i)

			// Skip if omitEmpty is true and the field is zero-value
			if omitEmpty && _isZeroValue(fieldValue) {
				continue
			}

			// Format the inventory line
			line := fmt.Sprintf("%s=\"%v\"", key, fieldValue.Interface())
			serverLines = append(serverLines, line)
		}

		// Combine lines for the server
		inventoryLines = append(inventoryLines, strings.Join(serverLines, " "))
	}

	// Join all server entries
	return strings.Join(inventoryLines, "\n"), nil
}

func _isZeroValue(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.String:
		return val.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Bool:
		return !val.Bool()
	case reflect.Interface, reflect.Ptr, reflect.Slice, reflect.Map, reflect.Array:
		return val.IsNil()
	}
	return false
}
