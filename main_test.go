package main

import (
	"net/mail"
	"testing"
)

type testMailer struct {
	Body    string
	Subject string
	To      []mail.Address
}

func (m *testMailer) Send(subject, body string, to ...mail.Address) error {
	//fmt.Printf("\n\n[TEST ONLY] sending an email to `%s` with a subject of: `%s` and a body:\n>>\n%s\n<<\n\n", to[0].Address, subject, body)

	m.Body = body
	m.Subject = subject
	m.To = to

	return nil
}

func TestSendEmail(t *testing.T) {
	m := &testMailer{}

	sendEmail(m, mail.Address{Name: "Tomek Wlodarczyk", Address: "twl@phase-ii.com"})

	if m.Body != "body here" {
		t.Errorf("Body expected to be: `%s`, received: %s", "body here", m.Body)
	}
	if m.Subject != "Subject" {
		t.Errorf("Subject expected to be: %s, received: %s", "Subject", m.Subject)
	}
	if m.To[0].Address != "twl@phase-ii.com" {
		t.Errorf("Email expected to be: %s, received: %s", "twl@phase-ii.com", m.To[0].Address)
	}
}
