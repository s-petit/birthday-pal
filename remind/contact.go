package remind

import (
	"github.com/s-petit/birthday-pal/contact"
)

//TODO SPE remplacer name et birthdate par contact ?
//renommer par contact with age ?
//ContactBirthday represents a Contact with birthday information
type ContactBirthday struct {
	Contact      contact.Contact
	Age       int
}

//ContactsToRemind filters every contacts which the bday matches the reminder's conditions.
func ContactsToRemind(contacts []contact.Contact, reminder Reminder) []ContactBirthday {

	var contactsToRemind []ContactBirthday

	for _, c := range contacts {

		if reminder.ShouldRemind(c.BirthDate) {
			contactsToRemind = append(contactsToRemind, ContactBirthday{c, c.Age(reminder.RemindDay())})
		}
	}

	return contactsToRemind
}
