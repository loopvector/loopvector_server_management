package controller

import (
	"errors"
	"fmt"
	"loopvector_server_management/model"
	"os"
	"path/filepath"
	"time"
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

func ValidateSession() (model.User, error) {
	token := loadSession()
	if token == "" {
		return model.User{}, errors.New("no active session found")
	}

	userSession, err := model.UserSession{
		Token: token,
	}.GetUsingToken()
	if err != nil {
		return model.User{}, err
	}
	if time.Now().After(userSession.ExpiresAt) {
		return model.User{}, errors.New("session expired or invalid")
	}

	user, err := userSession.GetUserUsingId()
	if err != nil {
		return model.User{}, err
	}

	return user, nil
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
