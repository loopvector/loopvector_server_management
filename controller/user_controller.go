package controller

import (
	"fmt"
	"loopvector_server_management/controller/helper"
	"loopvector_server_management/model"
	"net/smtp"
	"time"
)

// func GenerateRandomBytes(n int) ([]byte, error) {
// 	b := make([]byte, n)
// 	_, err := rand.Read(b)
// 	// Note that err == nil only if we read len(b) bytes.
// 	if err != nil {
// 		return nil, stacktrace.Propagate(err, "")
// 	}

// 	return b, nil
// }

// func GetHashedPassword(password string) (string, error) {
// 	saltedBytes := []byte(password)
// 	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}

// 	hash := string(hashedBytes[:])
// 	return hash, nil
// }

func RegisterUser(
	username *string,
	email, hashedPassword string,
	isAdmin bool,
) error {
	err := model.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		IsAdmin:  isAdmin,
	}.CreateNew()

	if err != nil {
		panic(err)
	}
	return nil
}

func LoginUser(email, password string) error {
	user, err := model.User{
		Email: email,
	}.GetUsingEmailId()
	if err != nil {
		panic(err)
	}
	verified := helper.VerifyPassword(password, user.Password)
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

	hashedPassword, err := helper.Encrypt(newPassword)
	if err != nil {
		panic(err)
	}
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
	adminSetting, err := model.LoadAdminSetting()
	if err != nil {
		panic(err)
	}

	if adminSetting.SMTPHost != "" && adminSetting.SMTPUser != "" {
		smtpAddress := fmt.Sprintf("%s:%d", adminSetting.SMTPHost, adminSetting.SMTPPort)
		auth := smtp.PlainAuth("", adminSetting.SMTPUser, adminSetting.SMTPPassword, adminSetting.SMTPHost)

		message := []byte(fmt.Sprintf(
			"Subject: Password Reset\n\nClick the link to reset your password: https://example.com/reset-password?token=%s\n",
			token,
		))

		err = smtp.SendMail(smtpAddress, auth, adminSetting.SMTPUser, []string{email}, message)
		if err != nil {
			panic("Failed to send email: " + err.Error())
		}

		fmt.Println("Password reset email sent")
	} else {
		fmt.Println("Password reset email not sent. Incomplete SMTP admin settings")
	}
}
