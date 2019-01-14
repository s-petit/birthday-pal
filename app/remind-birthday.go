package app

import (
	"github.com/s-petit/birthday-pal/email"
	"github.com/s-petit/birthday-pal/remind"
	"github.com/s-petit/birthday-pal/request"
	"log"
)

//App is an interface which represents an executable application
type App interface {
	Exec(contactsProvider request.ContactsProvider, smtp email.Sender, reminder remind.Reminder, recipients []string) error
}

//BirthdayPal represents this very software, and helps people to remind birthdays by sending reminders.
type BirthdayPal struct {
}

//Exec of BirthdayPal fetches contacts, then retrieve birthdays to remind, and finally send reminders to recipients
func (bp BirthdayPal) Exec(contactsProvider request.ContactsProvider, smtp email.Sender, reminder remind.Reminder, recipients []string) error {

	contacts, err := contactsProvider.GetContacts()
	if err != nil {
		return err
	}

	remindContacts := reminder.ContactsToRemind(contacts)
	contactsEmail := email.Contacts{Contacts: remindContacts, RemindDate: reminder.RemindDay()}

	err = smtp.Send(contactsEmail, recipients)
	if err != nil {
		return err
	}

	log.Printf("--> %d Reminder(s) sent. %d contact(s) will celebrate their birthday(s) in %d day(s), on %s.", len(remindContacts), len(remindContacts), reminder.InNbDays, reminder.RemindDay().Format("Mon, 02 Jan 2006"))

	return nil
}
