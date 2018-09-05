package remind

import (
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_should_remind_everyday_from_a_given_days_before_birthdate_until_birthday(t *testing.T) {
	birthday := testdata.BirthDate(2016, time.August, 22)

	aug23 := testdata.LocalDate(2018, time.August, 23)
	aug22 := testdata.LocalDate(2018, time.August, 22)
	aug21 := testdata.LocalDate(2018, time.August, 21)
	aug20 := testdata.LocalDate(2018, time.August, 20)
	aug19 := testdata.LocalDate(2018, time.August, 19)

	remindTomorrow := Reminder{CurrentDate: aug23, NbDaysBeforeBDay: 2, EveryDayUntilBDay: true}.remindEveryDay(birthday)
	assert.Equal(t, false, remindTomorrow)
	remindBirthday := Reminder{CurrentDate: aug22, NbDaysBeforeBDay: 2, EveryDayUntilBDay: true}.remindEveryDay(birthday)
	assert.Equal(t, true, remindBirthday)
	remindOneDayBefore := Reminder{CurrentDate: aug21, NbDaysBeforeBDay: 2, EveryDayUntilBDay: true}.remindEveryDay(birthday)
	assert.Equal(t, true, remindOneDayBefore)
	remindTwoDayBefore := Reminder{CurrentDate: aug20, NbDaysBeforeBDay: 2, EveryDayUntilBDay: true}.remindEveryDay(birthday)
	assert.Equal(t, true, remindTwoDayBefore)
	remindThreeDayBefore := Reminder{CurrentDate: aug19, NbDaysBeforeBDay: 2, EveryDayUntilBDay: true}.remindEveryDay(birthday)
	assert.Equal(t, false, remindThreeDayBefore)
}

func Test_should_remind_once_a_given_days_before_until_birthday(t *testing.T) {
	birthday := testdata.BirthDate(2016, time.August, 22)
	aug23 := testdata.LocalDate(2018, time.August, 23)
	aug22 := testdata.LocalDate(2018, time.August, 22)
	aug21 := testdata.LocalDate(2018, time.August, 21)
	aug20 := testdata.LocalDate(2018, time.August, 20)
	aug19 := testdata.LocalDate(2018, time.August, 19)

	remind := Reminder{CurrentDate: aug23, NbDaysBeforeBDay: 2, EveryDayUntilBDay: false}.remindOnce(birthday)
	assert.Equal(t, false, remind)
	remind = Reminder{CurrentDate: aug22, NbDaysBeforeBDay: 2, EveryDayUntilBDay: false}.remindOnce(birthday)
	assert.Equal(t, false, remind)
	remind = Reminder{CurrentDate: aug21, NbDaysBeforeBDay: 2, EveryDayUntilBDay: false}.remindOnce(birthday)
	assert.Equal(t, false, remind)
	remind = Reminder{CurrentDate: aug20, NbDaysBeforeBDay: 2, EveryDayUntilBDay: false}.remindOnce(birthday)
	assert.Equal(t, true, remind)
	remind = Reminder{CurrentDate: aug19, NbDaysBeforeBDay: 2, EveryDayUntilBDay: false}.remindOnce(birthday)
	assert.Equal(t, false, remind)
}

func Test_should_remind_once_when_current_day_is_a_birthday(t *testing.T) {
	birthday := testdata.BirthDate(2016, time.August, 22)
	aug22 := testdata.LocalDate(2018, time.August, 22)

	remind := Reminder{CurrentDate: aug22, NbDaysBeforeBDay: 0, EveryDayUntilBDay: false}.remindOnce(birthday)
	assert.Equal(t, true, remind)
}

func Test_should_remind_once_with_different_timezones(t *testing.T) {
	birthday := testdata.BirthDate(2016, time.August, 22)

	LAloc, _ := time.LoadLocation("America/Los_Angeles")
	SydneyLoc, _ := time.LoadLocation("Australia/Sydney")
	ParisLoc, _ := time.LoadLocation("Europe/Paris")

	aug22LA := time.Date(2018, time.August, 22, 0, 0 ,0 , 0, LAloc)

	remind := Reminder{CurrentDate: aug22LA, NbDaysBeforeBDay: 0, EveryDayUntilBDay: false}.remindOnce(birthday)
	assert.Equal(t, true, remind)

	aug22Sydney := time.Date(2018, time.August, 22, 0, 0 ,0 , 0, SydneyLoc)

	remind = Reminder{CurrentDate: aug22Sydney, NbDaysBeforeBDay: 0, EveryDayUntilBDay: false}.remindOnce(birthday)
	assert.Equal(t, true, remind)

	aug22Paris := time.Date(2018, time.August, 22, 0, 0 ,0 , 0, ParisLoc)

	remind = Reminder{CurrentDate: aug22Paris, NbDaysBeforeBDay: 0, EveryDayUntilBDay: false}.remindOnce(birthday)
	assert.Equal(t, true, remind)
}
