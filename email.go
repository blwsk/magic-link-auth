package main

// NOTE:
// https://github.com/tangingw/go_smtp/blob/master/send_mail.go

import (
  "strings"
  "net/smtp"
)

const (
  SMTP_SERVER = "smtp.gmail.com"
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

func (sender Sender) SendMail(recipients []string, Subject, bodyMessage string) error {
  msg := "From: " + sender.User + "\n" +
    "To: " + strings.Join(recipients, ",") + "\n" +
    "Subject: " + Subject + "\n" + bodyMessage

  return smtp.SendMail(SMTP_SERVER + ":587",
    smtp.PlainAuth("", sender.User, sender.Password, SMTP_SERVER),
    sender.User,
    recipients,
    []byte(msg))
}
