package remind

import (
	"github.com/s-petit/birthday-pal/contact"
	"time"
	"errors"
)

//ContactBirthday represents a Contact with birthday information
type ContactBirthday struct {
	Name      string
	BirthDate time.Time
	Age       int
}

//ContactsToRemind filters every contacts which the bday matches the reminder's conditions.
func ContactsToRemind(contacts []contact.Contact, reminder Reminder) []ContactBirthday {

	var contactsToRemind []ContactBirthday

	for _, c := range contacts {

		if reminder.ShouldRemind(c.BirthDate) {
			contactsToRemind = append(contactsToRemind, ContactBirthday{c.Name, c.BirthDate, c.Age(reminder.RemindDay())})
		}
	}

	return contactsToRemind
}


func WeeklyDigestContactsToRemind(contacts []contact.Contact, reminder Reminder) []ContactBirthday {

	if reminder.CurrentDate.Weekday() != time.Monday {
		return []ContactBirthday{}
	}

	weeklyReminder := reminder
	weeklyReminder.NbDaysBeforeBDay = 7
	weeklyReminder.EveryDayUntilBDay = true

	return ContactsToRemind(contacts, weeklyReminder)
}


func MonthlyDigestContactsToRemind(contacts []contact.Contact, reminder Reminder) []ContactBirthday {

	if reminder.CurrentDate.Day() != 1 {
		return []ContactBirthday{}
	}

	reminder.CurrentDate.Month()


	monthlyReminder := reminder
	monthlyReminder.NbDaysBeforeBDay = 7
	monthlyReminder.EveryDayUntilBDay = true

	return ContactsToRemind(contacts, monthlyReminder)
}


func searchLastDayOfMonth(date time.Time) (int, error) {

	// the last day of a month is between 28 and 31
	daysToCheck := []int{27, 28, 29, 30}

	for _, c := range daysToCheck {

		lol := date.AddDate(0, 0, c)
		if lol.Month() != date.Month() {
			return c-1, nil
		}
	}
	return 0, errors.New("waaaaaa")
}
