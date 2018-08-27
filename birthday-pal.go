package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/s-petit/birthday-pal/birthday"
	"github.com/s-petit/birthday-pal/carddav"
	"github.com/s-petit/birthday-pal/email"
	"github.com/s-petit/birthday-pal/vcardparser"
	"log"
	"os"
	"time"
)

func main() {

	app := cli.App("bpal", "Remind me birthdays pls.")

	app.Spec = "URL USERNAME PASSWORD RECIPIENTS..."

	var (
		recipients = app.StringsArg("RECIPIENTS", nil, "Reminders email recipients")
		cardDavUrl        = app.StringArg("URL", "", "cardDav URL")
		cardDavUsername   = app.StringArg("USERNAME", "", "basic auth username")
		cardDavPassword   = app.StringArg("PASSWORD", "", "basic auth password")
	)

	app.Action = func() {
		client := carddav.ContactClient{Url: *cardDavUrl, Username: *cardDavUsername, Password: *cardDavPassword}
		smtp := email.SmtpSender{Host: "smtp.fastmail.com", Port: "587", Username: "spetit@enjoycode.fr", Password: "awlh45n29jke5vsv"}

		DoIt(client, smtp, *recipients)
	}

	app.Run(os.Args)
}

func DoIt(client carddav.Client, smtp email.Sender, recipients []string) {
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
	cpt := 0
	//TODO as parameter pls.
	daysBefore := 32
	for _, card := range cards {
		date, _ := vcardparser.ParseVCardBirthDay(card)
		now := time.Now()
		remind := birthday.Remind(now, date, daysBefore)

		if remind {
			age := birthday.Age(now, date)
			contact := email.Contact{Name: card.FormattedName, BirthDate: date, Age: age}
			smtp.Send(contact, recipients)
			cpt++
		}
		//fmt.Printf("nom %s, anniv %s, formatted %s, remind %s \n", card.FormattedName, card.BirthDay, date, remind)
	}
	fmt.Printf("Rappels envoyés pour les %d anniversaire(s) qui auront lieu dans les %d jours.", cpt, daysBefore)

	//fmt.Println(contacts)
}

