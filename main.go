package main

import (
	"net/mail"

	mailer "go-interface/mailer"
)

func sendEmail(m *mailer.Mailer, to ...mail.Address) {

	m.Send("Subject", "body here", to...)
}

func main() {
	m := &mailer.Mailer{}

	sendEmail(m, mail.Address{Name: "Tomek Wlodarczyk", Address: "twl@phase-ii.com"})
}
