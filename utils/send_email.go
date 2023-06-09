package utils

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	from := "your-email@example.com"  // Replace with your email address
	password := "your-email-password" // Replace with your email password

	smtpHost := "smtp.example.com" // Replace with your SMTP server address
	smtpPort := "587"              // Replace with your SMTP server port

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, body))

	err := smtp.SendMail(fmt.Sprintf("%s:%s", smtpHost, smtpPort), auth, from, []string{to}, msg)
	if err != nil {
		return err
	}

	return nil
}
