package helper

import (
	"errors"
	"os"
	"path/filepath"
)

// GetFullFilePath validates the filename and generates the full file path.
// If the filepath is empty, it defaults to the current working directory.
// Returns the full file path or an error if validation fails.
func GetFullFilePath(filePath, fileName string) (string, error) {
	// Validate filename
	if fileName == "" {
		return "", errors.New("filename cannot be empty")
	}

	// Default to the current working directory if filePath is empty
	if filePath == "" {
		var err error
		filePath, err = os.Getwd()
		if err != nil {
			return "", errors.New("failed to get current working directory: " + err.Error())
		}
	}

	// Combine filepath and filename to get the full path
	fullPath := filepath.Join(filePath, fileName)
	return fullPath, nil
}

// WriteToFile writes content to a file at the specified filepath and filename.
// If the filepath is empty, it defaults to the current working directory.
// Returns an error if the filename is empty or any step fails.
func WriteToFile(filePath, fileName, content string) error {
	// Get the full file path
	fullPath, err := GetFullFilePath(filePath, fileName)
	if err != nil {
		return err
	}

	// Ensure the directory exists
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return errors.New("failed to create directories: " + err.Error())
	}

	// Create or overwrite the file
	file, err := os.Create(fullPath)
	if err != nil {
		return errors.New("failed to create file: " + err.Error())
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return errors.New("failed to write content to file: " + err.Error())
	}

	return nil
}
