package main

import (
	"github.com/jawher/mow.cli"
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/email"
	"github.com/s-petit/birthday-pal/remind"
	"github.com/s-petit/birthday-pal/request"
	"log"
	"os"
	"time"
)

func main() {

	app := cli.App("birthday-pal", "Remind me birthdays pls.")

	app.Spec = "[--carddav-url] [--carddav-user] [--carddav-pass] " +
		"[--smtp-host] [--smtp-port] [--smtp-user] [--smtp-pass] " +
		"[--days-before] [--remind-everyday] " +
		"RECIPIENTS..."

	var (

		// OPTS

		cardDavURL = app.String(cli.StringOpt{
			Name:   "carddav-url",
			Desc:   "CardDAV server URL",
			EnvVar: "BPAL_CARDDAV_URL",
		})

		cardDavUsername = app.String(cli.StringOpt{
			Name:   "carddav-user",
			Desc:   "CardDAV server username",
			EnvVar: "BPAL_CARDDAV_USERNAME",
		})

		cardDavPassword = app.String(cli.StringOpt{
			Name:   "carddav-pass",
			Desc:   "CardDAV server password",
			EnvVar: "BPAL_CARDDAV_PASSWORD",
		})

		SMTPHost = app.String(cli.StringOpt{
			Name:   "smtp-host",
			Desc:   "SMTP server hostname",
			EnvVar: "BPAL_SMTP_HOST",
		})

		SMTPPort = app.Int(cli.IntOpt{
			Name:   "smtp-port",
			Value:  587,
			Desc:   "SMTP server port",
			EnvVar: "BPAL_SMTP_PORT",
		})

		SMTPUsername = app.String(cli.StringOpt{
			Name:   "smtp-user",
			Desc:   "SMTP server username",
			EnvVar: "BPAL_SMTP_USERNAME",
		})

		SMTPPassword = app.String(cli.StringOpt{
			Name:   "smtp-pass",
			Desc:   "SMTP server password",
			EnvVar: "BPAL_SMTP_PASSWORD",
		})

		remindEveryDay = app.Bool(cli.BoolOpt{
			Name:  "e remind-everyday",
			Value: false,
			Desc:  "Send a reminder everyday until birthday instead of once.",
		})

		daysBefore = app.Int(cli.IntOpt{
			Name:  "d days-before",
			Value: 0,
			Desc:  "Number of days before birthday you want to be reminded.",
		})

		// ARGS

		recipients = app.Strings(cli.StringsArg{
			Name: "RECIPIENTS",
			Desc: "Reminders email recipients",
		})
	)

	app.Action = func() {

		// a tester
		// cardav avec basic
		// cardav avec oauth (theorique)
		// google avec oauth

		//TODO songer a remettre url dans basic auth
		auth := auth.BasicAuth{
			Username: *cardDavUsername,
			Password: *cardDavPassword,
		}

		//auth.Get(*cardDavURL)

		/*		oauth := http.OAuth2{
					Auth: google.authentication{Scope: people.ContactsReadonlyScope},
				}

				oauth.OauthClient().Get(*cardDavURL)
		*/

		contactsProvider := request.NewContactsProvider(*cardDavURL, auth)

		//provider, _ := http.GoogleContactsProvider{}.Get(ores)

		//fmt.Println(provider)

		smtp := email.SMTPClient{
			Host:     *SMTPHost,
			Port:     *SMTPPort,
			Username: *SMTPUsername,
			Password: *SMTPPassword,
		}

		reminder := remind.Reminder{
			CurrentDate:       time.Now(),
			NbDaysBeforeBDay:  *daysBefore,
			EveryDayUntilBDay: *remindEveryDay,
		}

		remindBirthdays(contactsProvider, smtp, reminder, *recipients)
	}

	app.Run(os.Args)
}

func crashIfError(err error) {
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

func remindBirthdays(contactsProvider request.ContactsProvider, smtp email.Sender, reminder remind.Reminder, recipients []string) {

	/*	if (client.)

		cardDavPayload, err := client.Get()
		crashIfError(err)*/
	contacts, err := contactsProvider.Get()
	crashIfError(err)

	remindContacts := remind.ContactsToRemind(contacts, reminder)

	for _, contact := range remindContacts {
		err := smtp.Send(contact, recipients)
		crashIfError(err)
	}

	log.Printf("--> %d Reminder(s) sent. %d contact(s) will celebrate their birthday(s) in %d day(s), on %s.", len(remindContacts), len(remindContacts), reminder.NbDaysBeforeBDay, reminder.RemindDay().Format("Mon, 02 Jan 2006"))
}
