package controller

import (
	"fmt"
	"loopvector_server_management/model"
	"os"
	"path/filepath"
)

func CreateNewUserSession(user model.User) error {
	token, err := user.CreateNewUserSession()
	if err != nil {
		panic(err)
	}
	saveSession(token)
	println("Login success!")
	return nil
}

func LoadCurrentUserSessionToken() (string, error) {
	token := loadSession()
	if token == "" {
		fmt.Println("No active session found")
		return "", nil
	}
	return token, nil
}

func saveSession(token string) {
	sessionFile := filepath.Join(os.TempDir(), "app_session")
	err := os.WriteFile(sessionFile, []byte(token), 0600)
	if err != nil {
		panic("Failed to save session")
	}
}

func loadSession() string {
	sessionFile := filepath.Join(os.TempDir(), "app_session")
	data, err := os.ReadFile(sessionFile)
	if err != nil {
		return ""
	}
	return string(data)
}

func clearSession() {
	sessionFile := filepath.Join(os.TempDir(), "app_session")
	_ = os.Remove(sessionFile) // Ignore errors if the file doesn't exist
}
