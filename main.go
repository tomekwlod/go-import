package main

import (
	"net/mail"

	mailer "go-interface/mailer"
)

func main() {
	m := &mailer.Mailer{}

	m.Send("Subject", "body here", mail.Address{Name: "Tomek Wlodarczyk", Address: "twl@phase-ii.com"})
}
