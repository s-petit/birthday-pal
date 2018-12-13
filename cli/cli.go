package cli

import (
	"fmt"
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
	bpal := cli.App("birthday-pal", "Remind me birthdays pls.")

	bpal.Spec = "[--smtp-host] [--smtp-port] [--smtp-user] [--smtp-pass] " +
		"[--days-before] [--remind-everyday] [--lang]"

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
			Name:  "e remind-everyday",
			Value: false,
			Desc:  "Send a reminder everyday until birthday instead of once.",
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

			reminder := remind.Reminder{
				CurrentDate:       system.Now(),
				NbDaysBeforeBDay:  *daysBefore,
				EveryDayUntilBDay: *remindEveryDay,
			}

			err := birthdayPal.Exec(contactsProvider, smtp, reminder, *recipients)
			crashIfError(err)
		}
	})

	bpal.Command("google", "use Google People API to retrieve contacts", func(cmd *cli.Cmd) {

		cmd.Spec = "[--url] PROFILE RECIPIENTS"

		var (
			// OPTS

			googleURL = cmd.String(cli.StringOpt{
				Name:   "u url",
				Desc:   "Google API URL",
				Value:  "https://people.googleapis.com/v1/people/me/connections?requestMask.includeField=person.names%2Cperson.birthdays&pageSize=500",
				EnvVar: "BPAL_GOOGLE_API_URL",
			})

			//ARGS

			profile = cmd.String(cli.StringArg{
				Name: "PROFILE",
				Desc: "birthday-pal oauth saved profile",
			})

			recipients = cmd.Strings(cli.StringsArg{
				Name: "RECIPIENTS",
				Desc: "Reminders email recipients",
			})
		)

		cmd.Action = func() {

			auth := auth.OAuth2Authenticator{
				Scope:   people.ContactsReadonlyScope,
				Profile: auth.OAuthProfile{system, *profile},
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

	bpal.Command("oauth", "identify birthday-pal to your oauth api provider", func(oauth *cli.Cmd) {

		oauth.Command("list", "list authenticated profiles", func(list *cli.Cmd) {

			list.Action = func() {
				profiles, err := auth.OAuthProfile{System: system}.ListProfiles()
				crashIfError(err)
				fmt.Println(profiles)
			}
		})

		oauth.Command("perform", "perform and save authentication for a profile", func(perform *cli.Cmd) {

			perform.Spec = "PROFILE SECRET"

			var (
				profile = perform.String(cli.StringArg{
					Name: "PROFILE",
					Desc: "Define a Oauth Authentication profile name",
				})

				secret = perform.String(cli.StringArg{
					Name: "SECRET",
					Desc: "Local Path to your OAuth2Authenticator client secret json file (for Google, download it on https://console.developers.google.com)",
				})
			)

			perform.Action = func() {

				auth := auth.OAuth2Authenticator{
					Scope:   people.ContactsReadonlyScope,
					Profile: auth.OAuthProfile{system, *profile},
				}

				err := auth.Authenticate(*secret)
				crashIfError(err)

				log.Println("Oauth2 authentication successful !")

			}
		})

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
