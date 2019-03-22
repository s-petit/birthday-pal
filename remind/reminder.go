package remind

import (
	"github.com/s-petit/birthday-pal/contact"
	"time"
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
	remindDayMinusAge := setYear(remindDay, remindDay.Year()-c.Age(remindDay))

	return !remindParams.Inclusive && remindDayMinusAge.Equal(c.BirthDate)
}

func (r Reminder) remindPeriod(c contact.Contact) bool {

	remindParams := r.RemindParams
	todayAtMidnight := remindParams.todayAtMidnight()
	remindDay := remindParams.RemindDay()

	todayAtMidnightMinusAge := setYear(todayAtMidnight, todayAtMidnight.Year()-c.Age(remindDay))
	remindDayMinusAge := setYear(remindDay, remindDay.Year()-c.Age(remindDay))

	return remindParams.Inclusive && (todayAtMidnightMinusAge.Before(c.BirthDate) || todayAtMidnightMinusAge.Equal(c.BirthDate)) && (remindDayMinusAge.After(c.BirthDate) || remindDayMinusAge.Equal(c.BirthDate))
}

func setYear(date time.Time, year int) time.Time {
	return time.Date(year, date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
}
