package birthday

import (
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/s-petit/birthday-pal/vcard"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//TODO distinguer UT et ITs

func Test_remind_contacts(t *testing.T) {
	date := testdata.LocalDate(2018, time.November, 21)
	c := vcard.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.November, 23)}
	c2 := vcard.Contact{Name: "Bill", BirthDate: testdata.BirthDate(1980, time.November, 28)}

	contactsToRemind := ContactsToRemind([]vcard.Contact{c, c2}, 2, date)

	expected := ContactBirthday{Name: "John", BirthDate: testdata.BirthDate(1980, time.November, 23), Age: 38}

	assert.Equal(t, []ContactBirthday{expected}, contactsToRemind)
}

/*
func Test_should_remind_when_birthday_is_in_one_day(t *testing.T) {

	birthdate := testdata.BirthDate(2016, time.August, 22)
	date := testdata.LocalDate(2018, time.August, 21)

	contact := Contact{Name: "John", BirthDate: birthdate}

	remind := contact.ShouldRemindBirthday( 1, date)

	assert.Equal(t, true, remind)
}

func Test_should_not_remind_when_birthday_is_in_more_than_one_day(t *testing.T) {

	birthdate := testdata.BirthDate(2016, time.August, 22)
	date := testdata.LocalDate(2018, time.August, 20)

	contact := Contact{Name: "John", BirthDate: birthdate}

	remind := contact.ShouldRemindBirthday(1, date)

	assert.Equal(t, false, remind)
}

func Test_should_not_remind_when_birthday_is_today(t *testing.T) {

	birthdate := testdata.BirthDate(2016, time.August, 22)
	date := testdata.LocalDate(2018, time.August, 22)

	contact := Contact{Name: "John", BirthDate: birthdate}
	remind := contact.ShouldRemindBirthday(1)

	assert.Equal(t, false, remind)
}

func Test_should_not_remind_when_birthday_is_in_the_past(t *testing.T) {

	birthdate := testdata.BirthDate(2016, time.August, 22)
	date := testdata.LocalDate(2018, time.August, 23)

	contact := Contact{Name: "John", BirthDate: birthdate}
	remind := contact.ShouldRemindBirthday(1, date)

	assert.Equal(t, false, remind)
}*/

// TODO more tests with timezones subtilities
