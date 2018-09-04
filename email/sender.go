package email

import (
	"github.com/s-petit/birthday-pal/remind"
	"net/smtp"
	"strconv"
)

// Sender represents a SMTP client
type Sender interface {
	Send(contactToRemind remind.ContactBirthday, recipients []string) error
}

// SMTPSender represents a SMTP client
type SMTPSender struct {
	Host     string
	Port     int
	Username string
	Password string
}

func (ss SMTPSender) hostPort() string {
	return ss.Host + ":" + strconv.Itoa(ss.Port)
}

// Send sends an email to remind the birthday of the related contact
func (ss SMTPSender) Send(contact remind.ContactBirthday, recipients []string) error {

	auth := smtp.PlainAuth("", ss.Username, ss.Password, ss.Host)

	// TODO mail hebdo pour les anniv de la semaine (notion d'inclure les 7 jours).

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		ss.hostPort(),
		auth,
		ss.Username,
		recipients,
		[]byte(French(contact)),
	)

	return err
}
