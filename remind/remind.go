package remind

import (
	"github.com/s-petit/birthday-pal/contact"
	"time"
)

//Reminder holds remind context and conditions
type Reminder struct {
	// should be now() in most cases, but this field is useful for testing purposes
	CurrentDate time.Time
	InNbDays    int
	// when false, check birthdays which matches exactly CurrentDate + InNbDays
	// when true, check birthdays which matches every days between CurrentDate and CurrentDate + InNbDays
	Inclusive bool
}

//RemindDay returns the day to remind
func (r Reminder) RemindDay() time.Time {
	return r.dateAtMidnight().AddDate(0, 0, r.InNbDays)
}


//ContactsToRemind filters every contacts which the bday matches the reminder's conditions.
func (r Reminder) ContactsToRemind(contacts []contact.Contact) []contact.Contact {

	var contactsToRemind []contact.Contact

	for _, c := range contacts {

		if r.shouldRemind(c) {
			contactsToRemind = append(contactsToRemind, c)
		}
	}

	return contactsToRemind
}


//ShouldRemind returns true when the birthdate should be reminded
func (r Reminder) shouldRemind(c contact.Contact) bool {
	return r.remindDay(c) || r.remindPeriod(c)
}


func (r Reminder) remindDay(c contact.Contact) bool {

	remindDay := r.RemindDay()
	birthDate := c.BirthDate

	return !r.Inclusive && remindDay.Day() == birthDate.Day() && remindDay.Month() == birthDate.Month()
}

func (r Reminder) remindPeriod(c contact.Contact) bool {

	dateAtMidnight := r.dateAtMidnight()
	remindDay := r.RemindDay()
	birthDate := c.BirthDate

	return r.Inclusive && (birthDate.Day() <= remindDay.Day() && birthDate.Month() <= remindDay.Month()) &&
		(birthDate.Day() >= dateAtMidnight.Day() && birthDate.Month() >= dateAtMidnight.Month())
}

func (r Reminder) dateAtMidnight() time.Time {
	return time.Date(r.CurrentDate.Year(), r.CurrentDate.Month(), r.CurrentDate.Day(), 0, 0, 0, 0, time.Local)
}
