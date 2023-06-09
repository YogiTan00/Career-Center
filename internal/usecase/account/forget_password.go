package account

import (
	"CareerCenter/utils"
	"context"
	"fmt"
)

func (u UseCaseAccountInteractor) ForgetPassword(ctx context.Context, email string) error {
	to := email // Replace with the recipient's email address
	subject := "Password Reset"
	body := "Click the following link to reset your password: http://example.com/reset?token=abcd1234"

	err := utils.SendEmail(to, subject, body)
	if err != nil {
		fmt.Println("Failed to send email:", err)
	} else {
		fmt.Println("Email sent successfully")
	}
	return nil
}

//func (u UseCaseAccountInteractor) ForgetPassword(ctx context.Context, email string) error {
//	pw := utils.RandomStringNumber(10)
//	token := fmt.Sprintf("%s-%d", email, time.Now().Unix())
//
//	from := "CareerCenter@gmail.com"
//	password := "careercenter"
//	to := []string{email}
//
//	url := fmt.Sprintf("http://localhost:8080/ganti-password?token=%s", token)
//	message := fmt.Sprintf("Klik tautan ini untuk mengganti password Anda: %s", url)
//
//	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
//	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, []byte(message))
//
//	if err != nil {
//		return err
//	}
//
//	savePw, err := utils.HashPassword(pw)
//	if err != nil {
//		return err
//	}
//	err = u.repoAccount.UpdatePassword(ctx, email, savePw)
//	if err != nil {
//		return err
//	}
//	return nil
//}
