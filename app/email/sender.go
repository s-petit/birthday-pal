package email

import (
	"github.com/s-petit/birthday-pal/app/contact"
	"github.com/s-petit/birthday-pal/app/remind"
	"log"
	"net/smtp"
	"strconv"
)

// Sender holds methods necessary for sending reminder emails.
type Sender interface {
	Send(emailContacts Contacts, recipients []string) error
}

// Contacts holds every contacts related data necessary for the email content.
type Contacts struct {
	Contacts     []contact.Contact
	RemindParams remind.Params
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

func (ss SMTPClient) auth() smtp.Auth {

	if len(ss.Username) > 0 && len(ss.Password) > 0 {
		return smtp.PlainAuth("", ss.Username, ss.Password, ss.Host)
	}

	return nil
}

//Send sends an email to recipients about the related contact incoming birthday.
func (ss SMTPClient) Send(emailContacts Contacts, recipients []string) error {

	if len(emailContacts.Contacts) < 1 {
		return nil
	}

	mail, err := toMail(emailContacts, ss.Language, ss.Username, recipients)
	if err != nil {
		return err
	}

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the subjectBody all in one step.
	err = smtp.SendMail(
		ss.hostPort(),
		ss.auth(),
		ss.Username,
		recipients,
		mail,
	)

	return err
}
