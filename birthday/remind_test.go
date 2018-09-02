package birthday

import (
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//TODO split into several tests ?
func Test_should_remind_everyday_from_nb_days_before_until_birthday(t *testing.T) {
	birthday := testdata.BirthDate(2016, time.August, 22)
	aug23 := testdata.LocalDate(2018, time.August, 23)
	aug22 := testdata.LocalDate(2018, time.August, 22)
	aug21 := testdata.LocalDate(2018, time.August, 21)
	aug20 := testdata.LocalDate(2018, time.August, 20)
	aug19 := testdata.LocalDate(2018, time.August, 19)

	remind := Reminder{CurrentDate: aug23, NbDaysBeforeBDay: 2, EveryDayUntilBDay: true}.remindEveryDay(birthday)
	assert.Equal(t, false, remind)
	remind = Reminder{CurrentDate: aug22, NbDaysBeforeBDay: 2, EveryDayUntilBDay: true}.remindEveryDay(birthday)
	assert.Equal(t, true, remind)
	remind = Reminder{CurrentDate: aug21, NbDaysBeforeBDay: 2, EveryDayUntilBDay: true}.remindEveryDay(birthday)
	assert.Equal(t, true, remind)
	remind = Reminder{CurrentDate: aug20, NbDaysBeforeBDay: 2, EveryDayUntilBDay: true}.remindEveryDay(birthday)
	assert.Equal(t, true, remind)
	remind = Reminder{CurrentDate: aug19, NbDaysBeforeBDay: 2, EveryDayUntilBDay: true}.remindEveryDay(birthday)
	assert.Equal(t, false, remind)
}

func Test_should_remind_once_nb_days_before_until_birthday(t *testing.T) {
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
