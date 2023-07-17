package utils

import (
	"CareerCenter/pkg/config"
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(mailer *gomail.Message, cfg config.Config) error {
	port, _ := strconv.Atoi(cfg.CONFIG_SMTP_PORT)
	dialer := gomail.NewDialer(
		cfg.CONFIG_SMTP_HOST,
		port,
		cfg.CONFIG_AUTH_EMAIL,
		cfg.CONFIG_AUTH_PASSWORD,
	)

	errMail := dialer.DialAndSend(mailer)

	if errMail != nil {
		log.Fatal(errMail.Error())
		return errMail
	}

	return nil
}
