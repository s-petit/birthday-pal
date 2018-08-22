package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/s-petit/birthday-pal/birthday"
	"github.com/s-petit/birthday-pal/carddav"
	"github.com/s-petit/birthday-pal/email"
	"github.com/s-petit/birthday-pal/vcardparser"
	"os"
	"time"
)

func main() {

	app := cli.App("bpal", "Remind me birthdays pls.")

	//app.Spec = "URL... RECIPIENTS PASSWORD"
	app.Spec = "URL USERNAME PASSWORD"

	var (
		//recipients = app.StringsArg("RECIPIENTS", nil, "Reminders email recipients")

		url      = app.StringArg("URL", "", "cardDav URL")
		username = app.StringArg("USERNAME", "", "basic auth username")
		password = app.StringArg("PASSWORD", "", "basic auth password")
	)

	app.Action = func() {
		contacts := carddav.Contacts(*url, *username, *password)
		cards := vcardparser.ParseContacts(contacts)

		cpt := 0
		daysBefore := 1

		for _, card := range cards {
			date, _ := vcardparser.ParseVCardBirthDay(card)
			remind := birthday.Remind(time.Now(), date, daysBefore)

			if remind {
				email.Send(card.FormattedName, date)
				cpt++
			}
			//fmt.Printf("nom %s, anniv %s, formatted %s, remind %s \n", card.FormattedName, card.BirthDay, date, remind)
		}

		fmt.Printf("Rappels envoyés pour les %d anniversaire(s) qui auront lieu dans les %d jours.", cpt, daysBefore)

		//fmt.Println(contacts)
	}

	// Invoke the app passing in os.Args
	app.Run(os.Args)
}
