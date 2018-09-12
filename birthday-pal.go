package main

import (
	"github.com/jawher/mow.cli"
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/email"
	"github.com/s-petit/birthday-pal/remind"
	"github.com/s-petit/birthday-pal/request"
	"google.golang.org/api/people/v1"
	"log"
	"os"
	"time"
)

func main() {

	app := cli.App("birthday-pal", "Remind me birthdays pls.")

	app.Spec = "[--smtp-host] [--smtp-port] [--smtp-user] [--smtp-pass] " +
		"[--days-before] [--remind-everyday]"

	var (

		// OPTS

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

	)

	app.Command("carddav", "use carddav to retrieve contacts", func(cmd *cli.Cmd) {

		cmd.Spec = "[--user] [--pass] [URL] RECIPIENTS..."

		var (

			// OPTS

			cardDavUsername = cmd.String(cli.StringOpt{
				Name:   "u user",
				Desc:   "CardDAV server username",
				EnvVar: "BPAL_CARDDAV_USERNAME",
			})

			cardDavPassword = cmd.String(cli.StringOpt{
				Name:   "p pass",
				Desc:   "CardDAV server password",
				EnvVar: "BPAL_CARDDAV_PASSWORD",
			})

			//ARGS

			cardDavURL = cmd.String(cli.StringArg{
				Name:   "URL",
				Desc:   "CardDAV server URL",
				EnvVar: "BPAL_CARDDAV_URL",
			})

			recipients = cmd.Strings(cli.StringsArg{
				Name: "RECIPIENTS",
				Desc: "Reminders email recipients",
			})
		)

		// Run this function when the command is invoked
		cmd.Action = func() {
			auth := auth.BasicAuth{
				Username: *cardDavUsername,
				Password: *cardDavPassword,
			}

			contactsProvider := request.CardDavContactsProvider{Client: auth, URL: *cardDavURL}

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
	})

	app.Command("google", "use Google People API to retrieve contacts", func(cmd *cli.Cmd) {

		cmd.Spec = "SECRET [URL]"

		var (
			// OPTS

			secret = cmd.String(cli.StringArg{
				Name: "SECRET",
				Desc: "Google OAuth2 client_secret.json",
			})

			//ARGS


			//TODO cette url ne devrait pas etre overridable. trop risque. on ne saura probablement pas traiter le resultat.
			googleURL = cmd.String(cli.StringArg{
				Name:   "URL",
				Desc:   "Google API URL",
				Value:  "https://people.googleapis.com/v1/people/me/connections?requestMask.includeField=person.names%2Cperson.birthdays&pageSize=500",
				EnvVar: "BPAL_GOOGLE_API_URL",
			})
			/*
				recipients = cmd.Strings(cli.StringsArg{
					Name: "RECIPIENTS",
					Desc: "Reminders email recipients",
				})*/
		)

		//TODO faire un birthday-pal smtp ? rendre obligatoire le smtp ou alors faire une erreur claire ?

		//TODO exporter le scope et l'url dans une fontion dediee ?
		cmd.Action = func() {

			auth := auth.OAuth2{
				Scope:      people.ContactsReadonlyScope,
				SecretPath: *secret,
			}

			contactsProvider := request.GoogleContactsProvider{Client: auth, URL: *googleURL}

			//TODO SPE: mutualiser smtp/reminder voire recipient
			/*			smtp := email.SMTPClient{
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
			*/
			contactsProvider.GetContacts()

			//remindBirthdays(contactsProvider, smtp, reminder, *recipients)
		}

	})

	app.Command("oauth", "identify birthday-pal to your oauth api provider", func(cmd *cli.Cmd) {

		cmd.Spec = "SECRET"

		var (
			secret = cmd.String(cli.StringArg{
				Name: "SECRET",
				Desc: "Google OAuth2 client_secret.json",
			})
		)

		//TODO voir si on peut auth plusieurs personnes... mais flemme...
		//TODO exporter le scope et l'url dans une fontion dediee ?
		cmd.Action = func() {

			auth := auth.OAuth2{
				Scope:      people.ContactsReadonlyScope,
				SecretPath: *secret,
			}

			err := auth.Authenticate()
			if err != nil {
				log.Fatal(err)
			} else {
				log.Println("Oauth2 authentication successful !")
			}

		}

	})

	app.Action = func() {

		// a tester
		// cardav avec basic
		// cardav avec oauth (theorique)
		// google avec oauth

		//TODO songer a remettre url dans basic auth
		/*		auth := auth.BasicAuth{
				Username: *cardDavUsername,
				Password: *cardDavPassword,
			}*/

		//auth.Get(*cardDavURL)

		/*		oauth := http.OAuth2{
					Auth: google.authentication{Scope: people.ContactsReadonlyScope},
				}

				oauth.oauthClient().Get(*cardDavURL)
		*/

		//contactsProvider := request.CardDavContactsProvider{auth, *cardDavURL}

		//provider, _ := http.GoogleContactsProvider{}.Get(ores)

		//fmt.Println(provider)

		/*		smtp := email.SMTPClient{
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

				remindBirthdays(contactsProvider, smtp, reminder, *recipients)*/
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
	contacts, err := contactsProvider.GetContacts()
	crashIfError(err)

	remindContacts := remind.ContactsToRemind(contacts, reminder)

	for _, contact := range remindContacts {
		err := smtp.Send(contact, recipients)
		crashIfError(err)
	}

	log.Printf("--> %d Reminder(s) sent. %d contact(s) will celebrate their birthday(s) in %d day(s), on %s.", len(remindContacts), len(remindContacts), reminder.NbDaysBeforeBDay, reminder.RemindDay().Format("Mon, 02 Jan 2006"))
}
