package vcard

import (
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//TODO testdata.BirthDate = date naissance birthDay = Anniver

func Test_should_calculate_age(t *testing.T) {

	birthday := testdata.BirthDate(1986, time.August, 22)
	date := testdata.LocalDate(2018, time.August, 23)

	c := Contact{"John", birthday}

	age := c.Age(date)

	assert.Equal(t, 32, age)
}

func Test_should_calculate_age_one_day_before_birthday(t *testing.T) {

	birthday := testdata.BirthDate(1986, time.August, 22)
	date := testdata.LocalDate(2018, time.August, 21)

	c := Contact{"John", birthday}

	age := c.Age(date)

	assert.Equal(t, 31, age)
}

func Test_age_should_be_negative_when_birthday_is_in_the_future(t *testing.T) {

	birthday := testdata.BirthDate(2019, time.August, 22)
	date := testdata.LocalDate(2018, time.August, 23)

	c := Contact{"John", birthday}

	age := c.Age(date)

	assert.Equal(t, -1, age)
}

func Test_age_should_be_0_for_new_born(t *testing.T) {
	birthday := testdata.BirthDate(2017, time.November, 22)
	date := testdata.LocalDate(2018, time.August, 23)

	c := Contact{Name: "John", BirthDate: birthday}

	age := c.Age(date)

	assert.Equal(t, 0, age)
}
