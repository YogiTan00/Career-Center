package account

import (
	"CareerCenter/utils"
	"context"
	"crypto/rand"
	"fmt"

	"gopkg.in/gomail.v2"
)

const otpChars = "1234567890"

func (u UseCaseAccountInteractor) ForgetPassword(ctx context.Context, email string) error {
	otp, errOtp := GenerateOTP(6)
	if errOtp != nil {
		return errOtp
	}
	// get user by mail
	user, err := u.repoAccount.GetByEmail(ctx, email)
	if err != nil {
		return err
	}
	// update code verifikasi
	errOtpUpdate := u.repoAccount.UpdateOTP(ctx, email, otp)
	if errOtpUpdate != nil {
		return errOtpUpdate
	}

	// send email after change password
	bodyContent := "Ini adalah kode verifikasi Anda: " + otp + ", Mohon pastikan Anda tidak pernah memberitahukan kode ini pada siapa pun. Terima kasih telah menggunakan layanan kami."
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", u.cfg.CONFIG_AUTH_EMAIL)
	mailer.SetHeader("To", user.Email)
	mailer.SetHeader("Subject", "Kode Verifikasi")
	mailer.SetBody("text/html", bodyContent)

	errSend := utils.SendEmail(mailer, u.cfg)
	if errSend != nil {
		fmt.Println("Failed to send email:", errSend)
	} else {
		fmt.Println("Email sent successfully")
	}
	return nil
}

func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
