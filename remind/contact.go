package remind

import (
	"github.com/s-petit/birthday-pal/contact"
	"time"
	"log"
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
	weeklyReminder.NbDaysBeforeBDay = 6
	weeklyReminder.EveryDayUntilBDay = true

	return ContactsToRemind(contacts, weeklyReminder)
}


func MonthlyDigestContactsToRemind(contacts []contact.Contact, reminder Reminder) []ContactBirthday {

	if reminder.CurrentDate.Day() != 1 {
		return []ContactBirthday{}
	}

	lastDayOfMonth := searchLastDayOfMonth(reminder.CurrentDate)

	monthlyReminder := reminder
	monthlyReminder.NbDaysBeforeBDay = lastDayOfMonth - 1
	monthlyReminder.EveryDayUntilBDay = true

	return ContactsToRemind(contacts, monthlyReminder)
}


func searchLastDayOfMonth(date time.Time) int {

	monthFirstDay := time.Date(date.Year(), date.Month(), 1, date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())

	// the last day of a month is between 28 and 31
	possibleLastDays := []int{28, 29, 30, 31}

	for _, dayNumber := range possibleLastDays {

		day := monthFirstDay.AddDate(0, 0, dayNumber)
		if day.Month() != date.Month() {
			return dayNumber
		}
	}
	log.Fatal("Unexpected error. Can not find last day of month for date: " + date.String())
	return -1
}
