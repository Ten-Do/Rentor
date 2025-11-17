package service

import (
	"fmt"
	"net/smtp"
)

type EmailService interface {
	SendEmail(to, subject, body string) error
}

type emailService struct {
	from     string
	password string
	host     string
	port     string
}

func NewEmailService(from, password, host, port string) EmailService {
	return &emailService{
		from:     from,
		password: password,
		host:     host,
		port:     port,
	}
}

func (s *emailService) SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.from, s.password, s.host)

	msg := []byte(
		"From: " + s.from + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=utf-8\r\n\r\n" +
			body,
	)

	addr := s.host + ":" + s.port

	return smtp.SendMail(addr, auth, s.from, []string{to}, msg)
}
