package contact

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func Test_should_calculate_age(t *testing.T) {

	birthday := simpleTime(1986, time.August, 22)
	now := simpleTime(2018, time.August, 23)

	assert.Equal(t, 32, Age(now, birthday))
}

func Test_should_calculate_age_one_day_before_birthday(t *testing.T) {

	birthday := simpleTime(1986, time.August, 22)
	now := simpleTime(2018, time.August, 21)

	assert.Equal(t, 31, Age(now, birthday))
}

func Test_age_should_be_negative_when_birthday_is_in_the_future(t *testing.T) {

	birthday := simpleTime(2019, time.August, 22)
	now := simpleTime(2018, time.August, 23)

	assert.Equal(t, -1, Age(now, birthday))
}

func Test_age_should_be_0_for_new_born(t *testing.T) {
	birthday := simpleTime(2017, time.November, 22)
	now := simpleTime(2018, time.August, 23)

	assert.Equal(t, 0, Age(now, birthday))
}

func simpleTime(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}
