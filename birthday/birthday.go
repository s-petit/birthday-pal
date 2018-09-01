package birthday

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
func ContactsToRemind(contacts []vcard.Contact, daysBefore int, date time.Time) []ContactBirthday {

	var contactsToRemind []ContactBirthday

	for _, c := range contacts {

		midnight := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
		delay := midnight.AddDate(0, 0, daysBefore)

		if delay.Day() == c.BirthDate.Day() && delay.Month() == c.BirthDate.Month() {
			contactsToRemind = append(contactsToRemind, ContactBirthday{c.Name, c.BirthDate, c.Age(delay)})
		}
	}

	return contactsToRemind
}
