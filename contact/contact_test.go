package contact

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//TODO birthDate = date naissance birthDay = Anniver

func Test_should_calculate_age(t *testing.T) {

	birthday := birthDate(1986, time.August, 22)
	now := localDate(2018, time.August, 23)

	contact := Contact{"John", birthday}

	assert.Equal(t, 32, contact.Age(now))
}

func Test_should_calculate_age_one_day_before_birthday(t *testing.T) {

	birthday := birthDate(1986, time.August, 22)
	now := localDate(2018, time.August, 21)

	contact := Contact{"John", birthday}

	assert.Equal(t, 31, contact.Age(now))
}

func Test_age_should_be_negative_when_birthday_is_in_the_future(t *testing.T) {

	birthday := birthDate(2019, time.August, 22)
	now := localDate(2018, time.August, 23)

	contact := Contact{"John", birthday}

	assert.Equal(t, -1, contact.Age(now))
}

func Test_age_should_be_0_for_new_born(t *testing.T) {
	birthday := birthDate(2017, time.November, 22)
	now := localDate(2018, time.August, 23)

	contact := Contact{"John", birthday}

	assert.Equal(t, 0, contact.Age(now))
}

func Test_should_remind_when_birthday_is_in_one_day(t *testing.T) {

	birthdate := birthDate(2016, time.August, 22)
	now := localDate(2018, time.August, 21)

	contact := Contact{"John", birthdate}

	remind := contact.ShouldRemindBirthday(now, 1)

	assert.Equal(t, true, remind)
}

func Test_should_not_remind_when_birthday_is_in_more_than_one_day(t *testing.T) {

	birthdate := birthDate(2016, time.August, 22)
	now := localDate(2018, time.August, 20)

	contact := Contact{"John", birthdate}

	remind := contact.ShouldRemindBirthday(now, 1)

	assert.Equal(t, false, remind)
}

func Test_should_not_remind_when_birthday_is_today(t *testing.T) {

	birthdate := birthDate(2016, time.August, 22)
	now := localDate(2018, time.August, 22)

	contact := Contact{"John", birthdate}
	remind := contact.ShouldRemindBirthday(now, 1)

	assert.Equal(t, false, remind)
}

func Test_should_not_remind_when_birthday_is_in_the_past(t *testing.T) {

	birthdate := birthDate(2016, time.August, 22)
	now := localDate(2018, time.August, 23)

	contact := Contact{"John", birthdate}
	remind := contact.ShouldRemindBirthday(now, 1)

	assert.Equal(t, false, remind)
}

func birthDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func localDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

// TODO more tests with timezones subtilities
