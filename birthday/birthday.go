package birthday

import (
	"github.com/s-petit/birthday-pal/contact"
	"time"
)

//ContactsToRemind filters every contacts which the bday occurs in daysBefore days
func ContactsToRemind(contacts []contact.Contact, daysBefore int, now time.Time) []contact.Contact {

	var contactsToRemind []contact.Contact

	for _, c := range contacts {
		if c.ShouldRemindBirthday(daysBefore, now) {
			contactsToRemind = append(contactsToRemind, c)
		}
	}

	return contactsToRemind
}


