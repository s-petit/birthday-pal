package app

import (
	"github.com/s-petit/birthday-pal/app/contact/request"
	"github.com/s-petit/birthday-pal/app/email"
	"github.com/s-petit/birthday-pal/app/remind"
	"log"
)

//App is an interface which represents an executable application
type App interface {
	Exec(contactsProvider request.ContactsProvider, smtp email.Sender, reminder remind.Params, recipients []string) error
}

//BirthdayPal represents this very software, and helps people to remind birthdays by sending reminders.
type BirthdayPal struct {
}

//Exec of BirthdayPal fetches contacts, then retrieve birthdays to remind, and finally send reminders to recipients
func (bp BirthdayPal) Exec(contactsProvider request.ContactsProvider, smtp email.Sender, remindParams remind.Params, recipients []string) error {

	contacts, err := contactsProvider.GetContacts()
	if err != nil {
		return err
	}

	remindContacts := remind.Reminder{Contacts: contacts, RemindParams: remindParams}.ContactsToRemind()
	contactsEmail := email.Contacts{Contacts: remindContacts, RemindParams: remindParams}

	err = smtp.Send(contactsEmail, recipients)
	if err != nil {
		return err
	}

	log.Printf("--> %d Reminder(s) sent. %d contact(s) will celebrate their birthday(s) in %d day(s), on %s.", len(remindContacts), len(remindContacts), remindParams.InNbDays, remindParams.RemindDay().Format("Mon, 02 Jan 2006"))

	return nil
}
