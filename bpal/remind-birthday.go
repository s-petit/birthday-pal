package bpal

import (
	"github.com/s-petit/birthday-pal/email"
	"github.com/s-petit/birthday-pal/remind"
	"github.com/s-petit/birthday-pal/request"
	"log"
)

//Pal represents the entity which helps fellow humains in the earth
type Pal interface {
	RemindBirthdays(contactsProvider request.ContactsProvider, smtp email.Sender, reminder remind.Reminder, recipients []string) error
}

//BirthdayPal represents the entity which helps people to remind birthdays.
type BirthdayPal struct {
}

//RemindBirthdays fetches contacts, then retrieve birthdays to remind, and finally send reminders to recipients
func (bp BirthdayPal) RemindBirthdays(contactsProvider request.ContactsProvider, smtp email.Sender, reminder remind.Reminder, recipients []string) error {

	contacts, err := contactsProvider.GetContacts()
	if err != nil {
		return err
	}

	remindContacts := remind.ContactsToRemind(contacts, reminder)

	for _, contact := range remindContacts {
		err := smtp.Send(contact, recipients)
		if err != nil {
			return err
		}
	}

	log.Printf("--> %d Reminder(s) sent. %d contact(s) will celebrate their birthday(s) in %d day(s), on %s.", len(remindContacts), len(remindContacts), reminder.NbDaysBeforeBDay, reminder.RemindDay().Format("Mon, 02 Jan 2006"))

	return nil
}
