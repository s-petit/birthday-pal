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

	app.Spec = "URL USERNAME PASSWORD DAYSBEFORE REMINDEVERYDAY SMTPHOST SMTPPORT SMTPUSER SMTPPASS RECIPIENTS..."

	var (
		cardDavURL      = app.StringArg("URL", "", "cardDav URL")
		cardDavUsername = app.StringArg("USERNAME", "", "basic auth username")
		cardDavPassword = app.StringArg("PASSWORD", "", "basic auth password")
		daysBefore      = app.IntArg("DAYSBEFORE", 1, "Send Reminder Days Before Birthday")
		remindEveryDay  = app.BoolArg("REMINDEVERYDAY", false, "Send only one reminder n days before bday or one reminder per day until bday ")
		SMTPURL         = app.StringArg("SMTPHOST", "", "SMTP URL")
		SMTPPort        = app.StringArg("SMTPPORT", "", "SMTP URL")
		SMTPUsername    = app.StringArg("SMTPUSER", "", "SMTP username")
		SMTPPassword    = app.StringArg("SMTPPASS", "", "SMTP password")
		recipients      = app.StringsArg("RECIPIENTS", nil, "Reminders email recipients")
	)

	app.Action = func() {
		client := carddav.BasicAuthRequest{URL: *cardDavURL, Username: *cardDavUsername, Password: *cardDavPassword}
		smtp := email.SMTPSender{Host: *SMTPURL, Port: *SMTPPort, Username: *SMTPUsername, Password: *SMTPPassword}
		reminder := birthday.Reminder{CurrentDate: time.Now(), NbDaysBeforeBDay: *daysBefore, EveryDayUntilBDay: *remindEveryDay}

		remindBirthdays(client, smtp, *recipients, reminder)
	}

	app.Run(os.Args)
}

func crashIfError(err error) {
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

func remindBirthdays(client carddav.Request, smtp email.Sender, recipients []string, reminder birthday.Reminder) {
	cardDavPayload, err := client.Get()
	crashIfError(err)

	contacts, err := vcard.ParseContacts(cardDavPayload)
	crashIfError(err)

	remindContacts := birthday.ContactsToRemind(contacts, reminder)

	for _, contact := range remindContacts {
		err := smtp.Send(contact, recipients)
		crashIfError(err)
	}

	fmt.Printf("Rappels envoyés pour les %d anniversaire(s) qui auront lieu dans les %d jours.", len(remindContacts), reminder.NbDaysBeforeBDay)
}
