package util

import (
	"gopkg.in/gomail.v2"
	"go_server/internal/config"
)

func SendEmail(to, subject, body string) error {
	cfg := config.Get()
	
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.SMTPUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPassword)

	return d.DialAndSend(m)
} 