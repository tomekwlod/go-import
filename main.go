package main

import (
	"net/mail"

	mailer "go-interface/mailer"
)

type mailerInterface interface {
	Send(subject, body string, to ...mail.Address) error
}

func sendEmail(m mailerInterface, to ...mail.Address) {

	m.Send("Subject", "body here", to...)
}

func main() {
	m := &mailer.Mailer{}

	sendEmail(m, mail.Address{Name: "Tomek Wlodarczyk", Address: "twl@phase-ii.com"})
}
