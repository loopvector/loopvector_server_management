package helper

import (
	"io"
	"os"
)

func ReadFileToString(filePath string) (string, error) {
	// Check if the file exists
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file content
	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
