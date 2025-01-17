package emailsender

import (
	"fmt"
	"net/smtp"
)

type IEmailSender interface {
	Send(subject, body string, recipients ...string) error
}

type EmailSender struct {
	smtpHost     string
	smtpPort     string
	smtpUsername string
	smtpPassword string
}

func New(smtpHost, smtpPort, smtpUsername, smtpPassword string) *EmailSender {
	return &EmailSender{
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		smtpUsername: smtpUsername,
		smtpPassword: smtpPassword,
	}
}

func (s *EmailSender) Send(subject, body string, recipients ...string) error {
	auth := smtp.PlainAuth("", s.smtpUsername, s.smtpPassword, s.smtpHost)
	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", s.smtpUsername, "", subject, body))
	err := smtp.SendMail(s.smtpHost+":"+s.smtpPort, auth, s.smtpUsername, recipients, msg)
	return err
}

type EmailSenderMock struct {
}

func NewMock() *EmailSenderMock {
	return &EmailSenderMock{}
}

func (sender *EmailSenderMock) Send(subject, body string, recipients ...string) error {
	msg := fmt.Sprintf("Subject: %s\nBody: %s\n", subject, body)
	fmt.Println("_____ Start EmailSenderMock message _____")
	fmt.Print(msg)
	fmt.Println("_____  End EmailSenderMock message  _____")
	return nil
}
