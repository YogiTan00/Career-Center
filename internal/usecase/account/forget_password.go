package account

import (
	"CareerCenter/utils"
	"context"
	"fmt"
	"net/smtp"
	"time"
)

func (u UseCaseAccountInteractor) ForgetPassword(ctx context.Context, email string) error {
	pw := utils.RandomStringNumber(10)
	token := fmt.Sprintf("%s-%d", email, time.Now().Unix())

	from := "CareerCenter@gmail.com"
	password := "careercenter"
	to := []string{email}

	url := fmt.Sprintf("http://localhost:8080/ganti-password?token=%s", token)
	message := fmt.Sprintf("Klik tautan ini untuk mengganti password Anda: %s", url)

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, []byte(message))

	if err != nil {
		return err
	}

	savePw, err := utils.HashPassword(pw)
	if err != nil {
		return err
	}
	err = u.repoAccount.UpdatePassword(ctx, email, savePw)
	if err != nil {
		return err
	}
	return nil
}
