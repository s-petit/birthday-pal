package birthday

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_should_remind_when_birthday_is_in_one_day(t *testing.T) {

	birthdate := BirthDate{2016, time.August, 22}
	now := time.Date(2018, time.August, 21, 0, 0 ,0, 0, time.Local)

	remind := birthdate.ShouldRemind(now, 1)

	assert.Equal(t, true, remind)
}

func Test_should_not_remind_when_birthday_is_in_more_than_one_day(t *testing.T) {

	birthdate := BirthDate{2016, time.August, 22}
	now := time.Date(2018, time.August, 20, 0, 0 ,0, 0, time.Local)

	remind := birthdate.ShouldRemind(now, 1)

	assert.Equal(t, false, remind)
}

func Test_should_not_remind_when_birthday_is_today(t *testing.T) {

	birthdate := BirthDate{2016, time.August, 22}
	now := time.Date(2018, time.August, 22, 0, 0 ,0, 0, time.Local)

	remind := birthdate.ShouldRemind(now, 1)

	assert.Equal(t, false, remind)
}

func Test_should_not_remind_when_birthday_is_in_the_past(t *testing.T) {

	birthdate := BirthDate{2016, time.August, 22}
	now := time.Date(2018, time.August, 23, 0, 0 ,0, 0, time.Local)

	remind := birthdate.ShouldRemind(now, 1)

	assert.Equal(t, false, remind)
}

