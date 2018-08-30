package birthday

import (
	"github.com/s-petit/birthday-pal/contact"
	"time"
)

//ContactsToRemind filters every contacts which the bday occurs in daysBefore days
func ContactsToRemind(contacts []contact.Contact, daysBefore int) []contact.Contact {

	var contactsToRemind []contact.Contact

	for _, c := range contacts {
		if c.ShouldRemindBirthday(time.Now(), daysBefore) {
			contactsToRemind = append(contactsToRemind, c)
		}
	}

	return contactsToRemind
}
