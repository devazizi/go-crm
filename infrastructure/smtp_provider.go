package infrastructure

import (
	"net/smtp"
	"os"
)

type SmtpConnection struct {
	Auth smtp.Auth
}

func NewSMTP() SmtpConnection {
	host := os.Getenv("SMTP_HOST")
	from := os.Getenv("MAIL_FROM")
	password := os.Getenv("MAIL_PASSWORD")

	return SmtpConnection{
		Auth: smtp.PlainAuth("", from, password, host),
	}
}

func (connection SmtpConnection) SendEmail(clients []string, message []byte) error {

	return smtp.SendMail(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		connection.Auth,
		os.Getenv("MAIL_FROM"),
		clients,
		message,
	)
}
