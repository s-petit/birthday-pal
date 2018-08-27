package email

import (
	"log"
	"net/smtp"
	"strconv"
	"time"
)

type Sender interface {
	Send(contact Contact, recipients []string)
}

type SmtpSender struct {
	Host string
	Port string
	Username string
	Password string
}

type Contact struct {
	Name string
	BirthDate time.Time
	Age int
}

func (ss SmtpSender) hostPort() string {
	return ss.Host + ":" + ss.Port
}

// Send sends an email to remind the birthday of the related contact
func (ss SmtpSender) Send(contact Contact, recipients []string) {

	auth := smtp.PlainAuth("", ss.Username, ss.Password, ss.Host)

	// TODO mail hebdo pour les anniv de la semaine (notion d'inclure les 7 jours).
	// TODO gerer les recipients

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		ss.hostPort(),
		auth,
		ss.Username,
		recipients,
		[]byte(frenchMailBody(contact)),
	)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO improve this horrible piece of code
func frenchMailBody(contact Contact) string {
	return "To: Birthday Pals \r\n" +
		"Subject: Anniversaire de " + contact.Name + "!\r\n" +
		"\r\n" +
		"Ce sera l'anniversaire de " + contact.Name + " le " + formatFrenchDate(contact.BirthDate) + ". Il aura " + strconv.Itoa(contact.Age) + " ans. Pensez a le lui souhaiter !\r\n"
}

func formatFrenchDate(birthday time.Time) string {
	const layout = "02/01"
	return birthday.Format(layout)
}

func formatEnglishDate(birthday time.Time) string {
	const layout = "01/02"
	return birthday.Format(layout)
}
