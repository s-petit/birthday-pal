package cli

import (
	"github.com/jawher/mow.cli"
	"github.com/s-petit/birthday-pal/app"
	"github.com/s-petit/birthday-pal/auth"
	"github.com/s-petit/birthday-pal/email"
	"github.com/s-petit/birthday-pal/remind"
	"github.com/s-petit/birthday-pal/request"
	"github.com/s-petit/birthday-pal/system"
	"google.golang.org/api/people/v1"
	"log"
	"os"
)

//Mowcli calls the mow.cli CLI which is the entry point of birthday-pal
func Mowcli(birthdayPal app.App, system system.System) {
	app := cli.App("birthday-pal", "Remind me birthdays pls.")

	app.Spec = "[--smtp-host] [--smtp-port] [--smtp-user] [--smtp-pass] " +
		"[--days-before] [--remind-everyday] [--lang]"

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

		language = app.String(cli.StringOpt{
			Name:   "l lang",
			Desc:   "Email language [EN, FR]",
			Value:  "EN",
			EnvVar: "BPAL_LANG",
		})

		// ARGS

	)

	app.Command("carddav", "use carddav to retrieve contacts", func(cmd *cli.Cmd) {

		cmd.Spec = "[--user] [--pass] [--url] RECIPIENTS..."

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

			cardDavURL = cmd.String(cli.StringOpt{
				Name:   "url",
				Desc:   "CardDAV server URL",
				EnvVar: "BPAL_CARDDAV_URL",
			})

			//ARGS

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

			contactsProvider := request.CardDavContactsProvider{AuthClient: auth, URL: *cardDavURL}

			smtp := email.SMTPClient{
				Host:     *SMTPHost,
				Port:     *SMTPPort,
				Username: *SMTPUsername,
				Password: *SMTPPassword,
				Language: *language,
			}

			reminder := remind.Reminder{
				CurrentDate:       system.Now(),
				NbDaysBeforeBDay:  *daysBefore,
				EveryDayUntilBDay: *remindEveryDay,
			}

			err := birthdayPal.Exec(contactsProvider, smtp, reminder, *recipients)
			crashIfError(err)
		}
	})

	app.Command("google", "use Google People API to retrieve contacts", func(cmd *cli.Cmd) {

		cmd.Spec = "[--url] SECRET RECIPIENTS"

		var (
			// OPTS

			googleURL = cmd.String(cli.StringOpt{
				Name:   "u url",
				Desc:   "Google API URL",
				Value:  "https://people.googleapis.com/v1/people/me/connections?requestMask.includeField=person.names%2Cperson.birthdays&pageSize=500",
				EnvVar: "BPAL_GOOGLE_API_URL",
			})

			//ARGS

			secret = cmd.String(cli.StringArg{
				Name: "SECRET",
				Desc: "Google OAuth2 client_secret.json",
			})

			recipients = cmd.Strings(cli.StringsArg{
				Name: "RECIPIENTS",
				Desc: "Reminders email recipients",
			})
		)

		cmd.Action = func() {

			auth := auth.OAuth2{
				Scope:      people.ContactsReadonlyScope,
				SecretPath: *secret,
				System:     system,
			}

			contactsProvider := request.GoogleContactsProvider{AuthClient: auth, URL: *googleURL}

			smtp := email.SMTPClient{
				Host:     *SMTPHost,
				Port:     *SMTPPort,
				Username: *SMTPUsername,
				Password: *SMTPPassword,
				Language: *language,
			}

			reminder := remind.Reminder{
				CurrentDate:       system.Now(),
				NbDaysBeforeBDay:  *daysBefore,
				EveryDayUntilBDay: *remindEveryDay,
			}

			err := birthdayPal.Exec(contactsProvider, smtp, reminder, *recipients)
			crashIfError(err)
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

		cmd.Action = func() {

			auth := auth.OAuth2{
				Scope:      people.ContactsReadonlyScope,
				SecretPath: *secret,
				System:     system,
			}

			err := auth.Authenticate()
			crashIfError(err)

			log.Println("Oauth2 authentication successful !")

		}

	})

	app.Action = func() {
		app.PrintHelp()
	}

	app.Run(os.Args)

}

func crashIfError(err error) {
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}
