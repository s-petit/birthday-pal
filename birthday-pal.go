package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/s-petit/birthday-pal/carddav"
	"github.com/s-petit/birthday-pal/email"
	"github.com/s-petit/birthday-pal/vcardparser"
	"log"
	"os"
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
		client := carddav.ContactClient{URL: *cardDavURL, Username: *cardDavUsername, Password: *cardDavPassword}
		smtp := email.SMTPSender{Host: "smtp.fastmail.com", Port: "587", Username: "spetit@enjoycode.fr", Password: "awlh45n29jke5vsv"}

		remindBirthdays(client, smtp, *recipients, *daysBefore)
	}

	app.Run(os.Args)
}

func remindBirthdays(client carddav.Client, smtp email.Sender, recipients []string, daysBefore int) {
	contacts, err := client.Get()
	if err != nil {
		log.Fatal("ERROR: ", err)
		os.Exit(1)
	}
	cards, err := vcardparser.ParseContacts(contacts)

	if err != nil {
		fmt.Println("An error occurred during VCard parsing. Please check that your URL refers to a CardDav endpoint.")
		log.Fatal("ERROR: ", err)
		os.Exit(1)
	}

	remindContacts := vcardparser.ContactsToRemind(cards, daysBefore)

	for _, contact := range remindContacts {
		smtp.Send(contact, recipients)
	}

	//fmt.Printf("nom %s, anniv %s, formatted %s, shouldRemind %s \n", card.FormattedName, card.BirthDay, date, shouldRemind)
	fmt.Printf("Rappels envoyés pour les %d anniversaire(s) qui auront lieu dans les %d jours.", len(remindContacts), daysBefore)

	//fmt.Println(contacts)
}
