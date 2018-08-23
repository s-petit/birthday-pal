package email

import (
	"log"
	"net/smtp"
	"strconv"
	"time"
)

// Send sends an email to remind the birthday of the related contact
func Send(name string, birthday time.Time, age int, recipients []string) {

	sender := "spetit@enjoycode.fr"
	host := "smtp.fastmail.com"
	port := "587"

	//TODO a regenerer !!
	password := "awlh45n29jke5vsv"
	auth := smtp.PlainAuth("", sender, password, host)

	// TODO mail hebdo pour les anniv de la semaine (notion d'inclure les 7 jours).
	// TODO gerer les recipients

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		host+":"+port,
		auth,
		sender,
		recipients,
		[]byte(frenchMailBody(name, birthday, age)),
	)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO improve this horrible piece of code
func frenchMailBody(name string, birthday time.Time, age int) string {
	return "To: Birthday Pals \r\n" +
		"Subject: Anniversaire de " + name + "!\r\n" +
		"\r\n" +
		"Ce sera l'anniversaire de " + name + " le " + formatFrenchDate(birthday) + ". Il aura " + strconv.Itoa(age) + " ans. Pensez a le lui souhaiter !\r\n"
}

func formatFrenchDate(birthday time.Time) string {
	const layout = "02/01"
	return birthday.Format(layout)
}

func formatEnglishDate(birthday time.Time) string {
	const layout = "01/02"
	return birthday.Format(layout)
}
