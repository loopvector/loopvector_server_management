package helper

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

func SplitYAMLTag(tag string) []string {
	parts := strings.Split(tag, ",")
	return parts
}

// Helper function to check if a value is "empty" (zero value for its type)
func IsEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

// GenerateConfig writes a struct to a YAML file
func GenerateConfig[T any](filePath string, data T) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	if err := encoder.Encode(&data); err != nil {
		return fmt.Errorf("failed to encode data to YAML: %w", err)
	}
	return nil
}

// LoadConfig reads a YAML file and unmarshals it into the provided struct
func LoadConfig[T any](filePath string) (T, error) {
	var data T

	file, err := os.Open(filePath)
	if err != nil {
		return data, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return data, fmt.Errorf("failed to decode YAML file: %w", err)
	}

	return data, nil
}

// UpdateConfig updates specific fields in a YAML file based on the provided struct
func UpdateConfig[T any](filePath string, updates T) error {
	// Load existing data
	existingData, err := LoadConfig[T](filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to load existing data: %w", err)
	}

	// Merge updates into the existing data
	mergedData := mergeStruct(existingData, updates)

	// Write the merged data back to the file
	return GenerateConfig(filePath, mergedData)
}

// mergeStruct merges non-zero fields from source into destination
func mergeStruct[T any](dest, src T) T {
	destBytes, _ := yaml.Marshal(dest)
	srcBytes, _ := yaml.Marshal(src)

	var destMap, srcMap map[string]interface{}
	_ = yaml.Unmarshal(destBytes, &destMap)
	_ = yaml.Unmarshal(srcBytes, &srcMap)

	for key, value := range srcMap {
		if value != nil {
			destMap[key] = value
		}
	}

	finalBytes, _ := yaml.Marshal(destMap)
	_ = yaml.Unmarshal(finalBytes, &dest)
	return dest
}
