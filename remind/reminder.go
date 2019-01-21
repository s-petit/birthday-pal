package remind

import (
	"github.com/s-petit/birthday-pal/contact"
)

//Reminder contains business logic between Params and Contacts
type Reminder struct {
	Contacts     []contact.Contact
	RemindParams Params
}

//ContactsToRemind filters every contacts which the bday matches the reminder's conditions.
func (r Reminder) ContactsToRemind() []contact.Contact {

	var contactsToRemind []contact.Contact

	for _, c := range r.Contacts {

		if r.shouldRemindFor(c) {
			contactsToRemind = append(contactsToRemind, c)
		}
	}

	return contactsToRemind
}

//ShouldRemind returns true when the birthdate should be reminded
func (r Reminder) shouldRemindFor(c contact.Contact) bool {
	return r.remindExactDay(c) || r.remindPeriod(c)
}

func (r Reminder) remindExactDay(c contact.Contact) bool {

	remindParams := r.RemindParams
	remindDay := remindParams.RemindDay()
	birthDate := c.BirthDate

	return !remindParams.Inclusive && remindDay.Day() == birthDate.Day() && remindDay.Month() == birthDate.Month()
}

func (r Reminder) remindPeriod(c contact.Contact) bool {

	remindParams := r.RemindParams
	dateAtMidnight := remindParams.dateAtMidnight()
	remindDay := remindParams.RemindDay()
	birthDate := c.BirthDate

	return remindParams.Inclusive && (birthDate.Day() <= remindDay.Day() && birthDate.Month() <= remindDay.Month()) &&
		(birthDate.Day() >= dateAtMidnight.Day() && birthDate.Month() >= dateAtMidnight.Month())
}
