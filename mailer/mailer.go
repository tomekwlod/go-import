package mailer

import (
	"fmt"
	"net/mail"
)

type Mailer struct{}

func (m *Mailer) Send(subject, body string, to ...mail.Address) error {

	fmt.Printf("\n\nsending an email to `%s` with a subject of: `%s` and a body:\n>>\n%s\n<<\n\n", to[0].Address, subject, body)

	return nil
}

func New() *Mailer {
	return &Mailer{}
}
