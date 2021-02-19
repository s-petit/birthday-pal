package cli

import (
	"github.com/jawher/mow.cli"
	"github.com/s-petit/birthday-pal/app"
	"github.com/s-petit/birthday-pal/app/contact/auth"
	"github.com/s-petit/birthday-pal/app/contact/request"
	"github.com/s-petit/birthday-pal/app/email"
	"github.com/s-petit/birthday-pal/app/remind"
	"github.com/s-petit/birthday-pal/system"
	"log"
	"os"
)

//Mowcli calls the mow.cli CLI which is the entry point of birthday-pal
func Mowcli(birthdayPal app.App, system system.System) {
	bpal := cli.App("birthday-pal", "Remind me birthdays pls.")

	bpal.Spec = "[--smtp-host] [--smtp-port] [--smtp-user] [--smtp-pass] " +
		"[--days-before] [--less-than-or-equal] [--lang]"

	var (

		// OPTS

		SMTPHost = bpal.String(cli.StringOpt{
			Name:   "smtp-host",
			Desc:   "SMTP server hostname",
			EnvVar: "BPAL_SMTP_HOST",
		})

		SMTPPort = bpal.Int(cli.IntOpt{
			Name:   "smtp-port",
			Value:  587,
			Desc:   "SMTP server port",
			EnvVar: "BPAL_SMTP_PORT",
		})

		SMTPUsername = bpal.String(cli.StringOpt{
			Name:   "smtp-user",
			Desc:   "SMTP server username",
			EnvVar: "BPAL_SMTP_USERNAME",
		})

		SMTPPassword = bpal.String(cli.StringOpt{
			Name:   "smtp-pass",
			Desc:   "SMTP server password",
			EnvVar: "BPAL_SMTP_PASSWORD",
		})

		remindEveryDay = bpal.Bool(cli.BoolOpt{
			Name:  "less-than-or-equal",
			Value: false,
			Desc:  "Activates 'less than or equal' operator to trigger reminders depending on 'days-before' option ('equal' by default)",
		})

		daysBefore = bpal.Int(cli.IntOpt{
			Name:  "d days-before",
			Value: 0,
			Desc:  "Number of days before birthday you want to be reminded.",
		})

		language = bpal.String(cli.StringOpt{
			Name:   "l lang",
			Desc:   "Email language [EN, FR]",
			Value:  "EN",
			EnvVar: "BPAL_LANG",
		})

		// ARGS

	)

	bpal.Command("carddav", "use carddav to retrieve contacts", func(cmd *cli.Cmd) {

		cmd.Spec = "[--user] [--pass] --url RECIPIENTS..."

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

			reminder := remind.Params{
				Today:     system.Now(),
				InNbDays:  *daysBefore,
				Inclusive: *remindEveryDay,
			}

			err := birthdayPal.Exec(contactsProvider, smtp, reminder, *recipients)
			crashIfError(err)
		}
	})

	bpal.Action = func() {
		bpal.PrintHelp()
	}

	bpal.Run(os.Args)

}

func crashIfError(err error) {
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}
