package services

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

const FROM = ""
const HOST = "smtp.gmail.com"
const PORT = "587"
const PASSWORD = ""

// Надіслати повідомлення
func SendEmail(to []string, subject string, body string) error {
	request := Mail{
		Sender:  FROM,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	msg := buildMessage(request)

	connectionString := fmt.Sprintf("%s:%s", HOST, PORT)
	auth := smtp.PlainAuth("", FROM, PASSWORD, HOST)
	err := smtp.SendMail(connectionString, auth, FROM, to, []byte(msg))

	if err != nil {
		return err
	}

	return nil
}

// Create message for SMTP
func buildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
