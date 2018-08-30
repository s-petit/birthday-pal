package birthday

import (
	"testing"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/testdata"
	"time"
	"github.com/stretchr/testify/assert"
)

//TODO distinguer UT et ITs

func Test_remind_contacts(t *testing.T) {
	now := testdata.LocalDate(2018, time.November, 21)
	c := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.November, 23)}
	c2 := contact.Contact{Name: "Bill", BirthDate: testdata.BirthDate(1980, time.November, 28)}

	contactsToRemind := ContactsToRemind([]contact.Contact{c, c2}, 2, now)

	expected := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.November, 23), Age: 38}

	assert.Equal(t, []contact.Contact{expected}, contactsToRemind)
}

func Test_should_calculate_age(t *testing.T) {

	birthday := testdata.BirthDate(1986, time.August, 22)
	now := testdata.LocalDate(2018, time.August, 23)

	age := ageAfterBirthday(birthday, now)

	assert.Equal(t, 32, age)
}

func Test_should_calculate_age_one_day_before_birthday(t *testing.T) {

	birthday := testdata.BirthDate(1986, time.August, 22)
	now := testdata.LocalDate(2018, time.August, 21)

	age := ageAfterBirthday(birthday, now)

	assert.Equal(t, 31, age)
}

func Test_age_should_be_negative_when_birthday_is_in_the_future(t *testing.T) {

	birthday := testdata.BirthDate(2019, time.August, 22)
	now := testdata.LocalDate(2018, time.August, 23)

	age := ageAfterBirthday(birthday, now)

	assert.Equal(t, -1, age)
}

func Test_age_should_be_0_for_new_born(t *testing.T) {
	birthday := testdata.BirthDate(2017, time.November, 22)
	now := testdata.LocalDate(2018, time.August, 23)

	age := ageAfterBirthday(birthday, now)

	assert.Equal(t, 0, age)
}