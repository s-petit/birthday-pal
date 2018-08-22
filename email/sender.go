package email

import (
	"time"
	"net/smtp"
	"log"
)

// Send sends an email to remind the birthday of the related contact
func Send(name string, birthday time.Time) {

	sender := "spetit@enjoycode.fr"
	recipients := []string{"hikaru95@gmail.com", "stephane.petit95@gmail.com"}
	host := "smtp.fastmail.com"
	port := "587"

	//TODO a regenerer !!
	password := "awlh45n29jke5vsv"
	auth := smtp.PlainAuth("", sender, password, host)

	// TODO indiquer l'age dans le mail
	// TODO mail quotidien pour l'anniv du jour (egalité parfaite)
	// TODO mail hebdo pour les anniv de la semaine (notion d'inclure les 7 jours).
	// TODO gerer les recipients
	//TODO CI, makefile, test coverage

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		host+":"+port,
		auth,
		sender,
		recipients,
		[]byte(frenchMailBody(name, birthday)),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func frenchMailBody(name string, birthday time.Time) string {
	return "To: recipient@example.net\r\n" +
		"Subject: Anniversaire de " + name + "!\r\n" +
		"\r\n" +
		"Ce sera l'anniversaire de " + name + " le " + formatFrenchDate(birthday) + ". Pensez à le lui souhaiter !\r\n"
}

func formatFrenchDate(birthday time.Time) string {
	const layout = "02/01"
	return birthday.Format(layout)
}

func formatEnglishDate(birthday time.Time) string {
	const layout = "01/02"
	return birthday.Format(layout)
}
