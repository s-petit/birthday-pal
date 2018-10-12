package remind

import (
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_remind_contacts(t *testing.T) {
	date := testdata.LocalDate(2018, time.November, 21)
	c := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.November, 23)}
	c2 := contact.Contact{Name: "Bill", BirthDate: testdata.BirthDate(1980, time.November, 28)}

	contactsToRemind := ContactsToRemind([]contact.Contact{c, c2}, Reminder{CurrentDate: date, NbDaysBeforeBDay: 2})

	expected := ContactBirthday{Name: "John", BirthDate: testdata.BirthDate(1980, time.November, 23), Age: 38}

	assert.Equal(t, []ContactBirthday{expected}, contactsToRemind)
}

func Test_remind_contacts_of_the_week(t *testing.T) {
	date := testdata.LocalDate(2018, time.October, 15) // monday
	c := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.October, 15)}
	c2 := contact.Contact{Name: "Bill", BirthDate: testdata.BirthDate(1980, time.October, 21)}
	c3 := contact.Contact{Name: "Dana", BirthDate: testdata.BirthDate(1980, time.October, 22)}
	c4 := contact.Contact{Name: "Lana", BirthDate: testdata.BirthDate(1980, time.October, 14)}

	contactsToRemind := WeeklyDigestContactsToRemind(
		[]contact.Contact{c, c2, c3, c4}, Reminder{CurrentDate: date, WeeklyDigest: true},
	)

	expected := []ContactBirthday{
		{Name: "John", BirthDate: testdata.BirthDate(1980, time.October, 15), Age: 38},
		{Name: "Bill", BirthDate: testdata.BirthDate(1980, time.October, 21), Age: 38},
	}

	assert.Equal(t, expected, contactsToRemind)
}

func Test_should_not_remind_contacts_of_the_week_when_not_monday(t *testing.T) {
	date := testdata.LocalDate(2018, time.October, 16) // not monday
	c := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.October, 15)}
	c2 := contact.Contact{Name: "Bill", BirthDate: testdata.BirthDate(1980, time.October, 21)}
	c3 := contact.Contact{Name: "Dana", BirthDate: testdata.BirthDate(1980, time.October, 22)}
	c4 := contact.Contact{Name: "Lana", BirthDate: testdata.BirthDate(1980, time.October, 14)}

	contactsToRemind := WeeklyDigestContactsToRemind(
		[]contact.Contact{c, c2, c3, c4}, Reminder{CurrentDate: date, WeeklyDigest: true},
	)

	expected := []ContactBirthday{}

	assert.Equal(t, expected, contactsToRemind)
}


func Test_remind_contacts_of_the_month(t *testing.T) {
	date := testdata.LocalDate(2018, time.October, 1) // 1st of the month
	c := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.October, 1)}
	c2 := contact.Contact{Name: "Bill", BirthDate: testdata.BirthDate(1980, time.October, 31)}
	c3 := contact.Contact{Name: "Dana", BirthDate: testdata.BirthDate(1980, time.November, 1)}
	c4 := contact.Contact{Name: "Lana", BirthDate: testdata.BirthDate(1980, time.September, 30)}

	contactsToRemind := MonthlyDigestContactsToRemind(
		[]contact.Contact{c, c2, c3, c4}, Reminder{CurrentDate: date, WeeklyDigest: true},
	)

	expected := []ContactBirthday{
		{Name: "John", BirthDate: testdata.BirthDate(1980, time.October, 1), Age: 38},
		{Name: "Bill", BirthDate: testdata.BirthDate(1980, time.October, 31), Age: 38},
	}

	assert.Equal(t, expected, contactsToRemind)
}

func Test_should_not_remind_contacts_of_the_week_when_not_first_day_of_month(t *testing.T) {
	date := testdata.LocalDate(2018, time.October, 2) // not 1st of the month
	c := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.October, 1)}
	c2 := contact.Contact{Name: "Bill", BirthDate: testdata.BirthDate(1980, time.October, 31)}
	c3 := contact.Contact{Name: "Dana", BirthDate: testdata.BirthDate(1980, time.November, 1)}
	c4 := contact.Contact{Name: "Lana", BirthDate: testdata.BirthDate(1980, time.September, 30)}

	contactsToRemind := MonthlyDigestContactsToRemind(
		[]contact.Contact{c, c2, c3, c4}, Reminder{CurrentDate: date, WeeklyDigest: true},
	)

	expected := []ContactBirthday{}

	assert.Equal(t, expected, contactsToRemind)
}

func Test_last_day_of_30_days_month(t *testing.T) {
	date := testdata.LocalDate(2018, time.November, 21)

	lastDayMonth := searchLastDayOfMonth(date)
	assert.Equal(t, 30, lastDayMonth)
}

func Test_last_day_of_31_days_month(t *testing.T) {
	date := testdata.LocalDate(2018, time.January, 21)

	lastDayMonth := searchLastDayOfMonth(date)
	assert.Equal(t, 31, lastDayMonth)
}

func Test_last_day_of_28_days_month(t *testing.T) {
	date := testdata.LocalDate(2018, time.February, 21)

	lastDayMonth := searchLastDayOfMonth(date)
	assert.Equal(t, 28, lastDayMonth)
}

func Test_last_day_of_29_days_month(t *testing.T) {
	date := testdata.LocalDate(2016, time.February, 21)

	lastDayMonth := searchLastDayOfMonth(date)
	assert.Equal(t, 29, lastDayMonth)
}
