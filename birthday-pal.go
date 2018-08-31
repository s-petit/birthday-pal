package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/s-petit/birthday-pal/birthday"
	"github.com/s-petit/birthday-pal/carddav"
	"github.com/s-petit/birthday-pal/email"
	"github.com/s-petit/birthday-pal/vcard"
	"log"
	"os"
	"time"
)

func main() {

	app := cli.App("bpal", "Remind me birthdays pls.")

	app.Spec = "URL USERNAME PASSWORD DAYSBEFORE RECIPIENTS..."

	var (
		recipients      = app.StringsArg("RECIPIENTS", nil, "Reminders email recipients")
		cardDavURL      = app.StringArg("URL", "", "cardDav URL")
		cardDavUsername = app.StringArg("USERNAME", "", "basic auth username")
		cardDavPassword = app.StringArg("PASSWORD", "", "basic auth password")
		daysBefore      = app.IntArg("DAYSBEFORE", 1, "Send Reminder Days Before Birthday")
	)

	app.Action = func() {
		client := carddav.BasicAuthRequest{URL: *cardDavURL, Username: *cardDavUsername, Password: *cardDavPassword}
		smtp := email.SMTPSender{Host: "smtp.fastmail.com", Port: "587", Username: "spetit@enjoycode.fr", Password: "awlh45n29jke5vsv"}

		remindBirthdays(client, smtp, *recipients, *daysBefore, time.Now())
	}

	app.Run(os.Args)
}

func crashIfError(err error) {
	if err != nil {
		log.Fatal("ERROR: ", err)
		os.Exit(1)
	}
}

func remindBirthdays(client carddav.Request, smtp email.Sender, recipients []string, daysBefore int, date time.Time) {
	cardDavPayload, err := client.Get()
	crashIfError(err)

	contacts, err := vcard.ParseContacts(cardDavPayload)
	crashIfError(err)

	remindContacts := birthday.ContactsToRemind(contacts, daysBefore, date)

	for _, contact := range remindContacts {
		err := smtp.Send(contact, recipients)
		crashIfError(err)
	}

	//fmt.Printf("nom %s, anniv %s, formatted %s, shouldRemind %s \n", card.FormattedName, card.BirthDay, date, shouldRemind)
	fmt.Printf("Rappels envoyés pour les %d anniversaire(s) qui auront lieu dans les %d jours.", len(remindContacts), daysBefore)

	//fmt.Println(contacts)
}
