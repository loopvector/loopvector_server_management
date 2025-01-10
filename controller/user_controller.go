package controller

import (
	"fmt"
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
	"net/smtp"
	"time"
)

func RegisterUser(username, email, hashedPassword string) error {
	return model.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}.CreateNew()
}

func LoginUser(email, password string) error {
	user, err := model.User{
		Email: email,
	}.GetUsingEmailId()
	if err != nil {
		panic(err)
	}
	verified, err := helper.VerifyPassword(password, user.Password)
	if err != nil {
		panic(err)
	}
	if !verified { // Implement password verification
		panic("Invalid password")
	}
	return CreateNewUserSession(user)
}

func LogoutUser() error {
	token, err := LoadCurrentUserSessionToken()
	if err != nil {
		panic(err)
	}
	if token == "" {
		return nil
	}
	err = model.UserSession{
		Token: token,
	}.DeleteUserSession()
	if err != nil {
		panic(err)
	}
	clearSession()
	println("Logout success!")
	return nil
}

func ForgotPassword(email string) error {
	user, err := model.User{
		Email: email,
	}.GetUsingEmailId()
	if err != nil {
		panic(err)
	}
	token, err := user.CreateNewPasswordResetToken()
	if err != nil {
		panic(err)
	}
	sendPasswordResetEmail(user.Email, token)
	return nil
}

func ResetPassword(token, newPassword string) error {
	resetToken, err := model.PasswordResetToken{
		Token: token,
	}.GetUsingToken()
	if err != nil {
		panic(err)
	}
	if time.Now().After(resetToken.ExpiresAt) {
		panic("Token has expired")
	}
	user, err := model.PasswordResetToken{
		UserID: resetToken.UserID,
	}.GetUserUsingId()
	if err != nil {
		panic(err)
	}

	hashedPassword := helper.HashPassword(newPassword)
	user.Password = hashedPassword

	err = user.UpdatePassword()
	if err != nil {
		panic(err)
	}

	err = resetToken.DeleteUsingToken()
	if err != nil {
		panic(err)
	}

	return nil
}

func sendPasswordResetEmail(email, token string) {
	// db := getDB()
	var settings model.SmtpSetting
	settings, err := model.GetFirstSmtpSetting()
	if err != nil {
		panic(err)
	}

	smtpAddress := fmt.Sprintf("%s:%d", settings.SMTPHost, settings.SMTPPort)
	auth := smtp.PlainAuth("", settings.SMTPUser, settings.SMTPPassword, settings.SMTPHost)

	message := []byte(fmt.Sprintf(
		"Subject: Password Reset\n\nClick the link to reset your password: https://example.com/reset-password?token=%s\n",
		token,
	))

	err = smtp.SendMail(smtpAddress, auth, settings.SMTPUser, []string{email}, message)
	if err != nil {
		panic("Failed to send email: " + err.Error())
	}

	fmt.Println("Password reset email sent")
}
