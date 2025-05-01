package smtp

import (
	"fmt"
	"net/smtp"
)

type EmailSender interface {
	SendEmail(to []string, subject, body string) error
}

type SMTPEmailSender struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

func NewSMTPEmailSender(host string, port string, username, password, from string) *SMTPEmailSender {

	return &SMTPEmailSender{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (s *SMTPEmailSender) SendEmail(to []string, subject, body string) error {
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)

	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		s.From, to, subject, body,
	))

	if err := smtp.SendMail(addr, auth, s.From, to, msg); err != nil {
		return err
	}
	return nil
}
