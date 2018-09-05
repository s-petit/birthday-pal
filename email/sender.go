package email

import (
	"github.com/s-petit/birthday-pal/remind"
	"net/smtp"
	"strconv"
)

// Sender holds methods necessary for sending reminder emails.
type Sender interface {
	Send(contactToRemind remind.ContactBirthday, recipients []string) error
}

// SMTPClient represents a SMTP client with its credentials
type SMTPClient struct {
	Host     string
	Port     int
	Username string
	Password string
}

func (ss SMTPClient) hostPort() string {
	return ss.Host + ":" + strconv.Itoa(ss.Port)
}

//Send sends an email to recipients about the related contact incoming birthday.
func (ss SMTPClient) Send(contact remind.ContactBirthday, recipients []string) error {

	auth := smtp.PlainAuth("", ss.Username, ss.Password, ss.Host)

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
