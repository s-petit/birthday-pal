package remind

import (
	"github.com/s-petit/birthday-pal/vcard"
	"time"
)

//ContactBirthday represents a Contact eligible for reminder because his bday is near.
type ContactBirthday struct {
	Name      string
	BirthDate time.Time
	Age       int
}

//ContactsToRemind filters every contacts which the bday occurs in daysBefore days
func ContactsToRemind(contacts []vcard.Contact, reminder Reminder) []ContactBirthday {

	var contactsToRemind []ContactBirthday

	for _, c := range contacts {

		if reminder.remindEveryDay(c.BirthDate) || reminder.remindOnce(c.BirthDate) {
			contactsToRemind = append(contactsToRemind, ContactBirthday{c.Name, c.BirthDate, c.Age(reminder.RemindDay())})
		}
	}

	return contactsToRemind
}
