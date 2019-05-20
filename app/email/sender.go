package email

import (
	"github.com/s-petit/birthday-pal/app/email/i18n"
	"log"
	"net/smtp"
	"strconv"
)

// Sender holds methods necessary for sending reminder emails.
type Sender interface {
	Send(emailContacts i18n.Contacts, recipients []string) error
}

// SMTPClient represents a SMTP client with its credentials
type SMTPClient struct {
	Host     string
	Port     int
	Username string
	Password string
	Language string
}

func (ss SMTPClient) hostPort() string {

	if len(ss.Host) < 1 {
		log.Printf("WARNING: SMTP host is empty")
	}

	return ss.Host + ":" + strconv.Itoa(ss.Port)
}

//Send sends an email to recipients about the related contact incoming birthday.
func (ss SMTPClient) Send(emailContacts i18n.Contacts, recipients []string) error {

	if len(emailContacts.Contacts) < 1 {
		return nil
	}

	auth := smtp.PlainAuth("", ss.Username, ss.Password, ss.Host)

	mail, err := i18n.ToMail(emailContacts, ss.Language)
	if err != nil {
		return err
	}

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the subjectBody all in one step.
	err = smtp.SendMail(
		ss.hostPort(),
		auth,
		ss.Username,
		recipients,
		mail,
	)

	return err
}
