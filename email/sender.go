package email

import (
	"github.com/s-petit/birthday-pal/remind"
	"net/smtp"
	"strconv"
	"time"
)

// Sender holds methods necessary for sending reminder emails.
type Sender interface {
	Send(emailContacts Contacts, recipients []string) error
}

// Contacts holds every contacts related data necessary for the email content.
type Contacts struct {
	Contacts   []remind.ContactBirthday
	RemindDate time.Time
}

// SMTPClient represents a SMTP client with its credentials
type SMTPClient struct {
	Host     string
	Port     int
	Username string
	Password string
	Language string
	EmailType 	 string
}

func (clt SMTPClient) hostPort() string {
	return clt.Host + ":" + strconv.Itoa(clt.Port)
}

//Send sends an email to recipients about the related contact incoming birthday.
func (clt SMTPClient) Send(emailContacts Contacts, recipients []string) error {

	auth := smtp.PlainAuth("", clt.Username, clt.Password, clt.Host)

	mail, err := toMail(emailContacts, clt.Language, clt.EmailType)
	if err != nil {
		return err
	}

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the subjectBody all in one step.
	err = smtp.SendMail(
		clt.hostPort(),
		auth,
		clt.Username,
		recipients,
		mail,
	)

	return err
}

