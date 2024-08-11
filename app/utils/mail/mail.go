package mail

import (
	"gopkg.in/gomail.v2"
)

func Prepare(from string, to string, subject string, body string) (*gomail.Dialer, *gomail.Message) {
	msg := gomail.NewMessage()

	msg.SetHeader("From", from)

	msg.SetHeader("To", to)

	msg.SetHeader("Subject", subject)

	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer("localhost", 1025, from, "")

	return dialer, msg
}

func Send(dialer *gomail.Dialer, message *gomail.Message) error {
	err := dialer.DialAndSend(message)
	return err
}
