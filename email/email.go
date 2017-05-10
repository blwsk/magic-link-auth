package email

// NOTE:
// https://github.com/tangingw/go_smtp/blob/master/send_mail.go

import (
  "os"
  // "strings"
  "net/smtp"
)

var (
  SMTP_SERVER = os.Getenv("SMTP_SERVER")
)

type Sender struct {
  User      string
  Password  string
}

func NewSender(Username, Password string) Sender {
  return Sender{
    Username,
    Password,
  }
}

func (sender Sender) SendMail(recipients []string, subject string, body string) error {
  server := SMTP_SERVER + ":587"
  auth := smtp.PlainAuth("", sender.User, sender.Password, SMTP_SERVER)
  from := sender.User

  mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";

  subjectStr := "Subject: " + subject + "\n"

  msg := []byte(subjectStr + mime + body)

  return smtp.SendMail(server, auth, from, recipients, msg)
}
