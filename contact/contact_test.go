package contact

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"github.com/s-petit/birthday-pal/testdata"
)

//TODO testdata.BirthDate = date naissance birthDay = Anniver

func Test_should_remind_when_birthday_is_in_one_day(t *testing.T) {

	birthdate := testdata.BirthDate(2016, time.August, 22)
	now := testdata.LocalDate(2018, time.August, 21)

	contact := Contact{Name: "John", BirthDate: birthdate}

	remind := contact.ShouldRemindBirthday( 1, now)

	assert.Equal(t, true, remind)
}

func Test_should_not_remind_when_birthday_is_in_more_than_one_day(t *testing.T) {

	birthdate := testdata.BirthDate(2016, time.August, 22)
	now := testdata.LocalDate(2018, time.August, 20)

	contact := Contact{Name: "John", BirthDate: birthdate}

	remind := contact.ShouldRemindBirthday(1, now)

	assert.Equal(t, false, remind)
}

func Test_should_not_remind_when_birthday_is_today(t *testing.T) {

	birthdate := testdata.BirthDate(2016, time.August, 22)
	now := testdata.LocalDate(2018, time.August, 22)

	contact := Contact{Name: "John", BirthDate: birthdate}
	remind := contact.ShouldRemindBirthday(1)

	assert.Equal(t, false, remind)
}

func Test_should_not_remind_when_birthday_is_in_the_past(t *testing.T) {

	birthdate := testdata.BirthDate(2016, time.August, 22)
	now := testdata.LocalDate(2018, time.August, 23)

	contact := Contact{Name: "John", BirthDate: birthdate}
	remind := contact.ShouldRemindBirthday(1, now)

	assert.Equal(t, false, remind)
}

// TODO more tests with timezones subtilities
