package contact

import (
	"github.com/s-petit/birthday-pal/remind"
	"time"
)

//Contact represents a Contact.
type Contact struct {
	Name      string
	BirthDate time.Time
}

//ContactsToRemind filters every contacts which the bday matches the reminder's conditions.
func ContactsToRemind(contacts []Contact, reminder remind.Reminder) []Contact {

	var contactsToRemind []Contact

	for _, c := range contacts {

		if reminder.ShouldRemind(c.BirthDate) {
			contactsToRemind = append(contactsToRemind, c)
		}
	}

	return contactsToRemind
}

