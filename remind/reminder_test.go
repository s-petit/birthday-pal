package remind

import (
	"fmt"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var may30 = testdata.LocalDate(2018, time.May, 30)
var may31 = testdata.LocalDate(2018, time.May, 31)
var jun1 = testdata.LocalDate(2018, time.June, 1)
var dec31 = testdata.LocalDate(2018, time.December, 31)
var dec30 = testdata.LocalDate(2018, time.December, 30)
var jan1 = testdata.LocalDate(2019, time.January, 1)
var aug21 = testdata.LocalDate(2018, time.August, 21)
var aug20 = testdata.LocalDate(2018, time.August, 20)
var aug19 = testdata.LocalDate(2018, time.August, 19)

var john = contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.August, 20)}
var sara = contact.Contact{Name: "Sara", BirthDate: testdata.BirthDate(1994, time.August, 21)}
var rob = contact.Contact{Name: "Rob", BirthDate: testdata.BirthDate(1994, time.May, 31)}
var jane = contact.Contact{Name: "Jane", BirthDate: testdata.BirthDate(1994, time.June, 1)}
var jill = contact.Contact{Name: "Jill", BirthDate: testdata.BirthDate(1994, time.December, 31)}
var nick = contact.Contact{Name: "Nick", BirthDate: testdata.BirthDate(1994, time.January, 1)}

var contacts = []contact.Contact{john, sara}

type dataprovider struct {
	in  Params
	out []contact.Contact
}

var dataProvider = []dataprovider{
	{Params{Today: aug19, InNbDays: 0, Inclusive: false}, []contact.Contact{}},
	{Params{Today: aug19, InNbDays: 1, Inclusive: false}, []contact.Contact{john}},
	{Params{Today: aug19, InNbDays: 2, Inclusive: false}, []contact.Contact{sara}},
	{Params{Today: aug19, InNbDays: 3, Inclusive: false}, []contact.Contact{}},
	{Params{Today: aug19, InNbDays: 0, Inclusive: true}, []contact.Contact{}},
	{Params{Today: aug19, InNbDays: 1, Inclusive: true}, []contact.Contact{john}},
	{Params{Today: aug19, InNbDays: 2, Inclusive: true}, []contact.Contact{john, sara}},
	{Params{Today: aug19, InNbDays: 3, Inclusive: true}, []contact.Contact{john, sara}},
	{Params{Today: aug20, InNbDays: 0, Inclusive: false}, []contact.Contact{john}},
}

var dataProviderOnSeveralMonths = []dataprovider{
	{Params{Today: may30, InNbDays: 0, Inclusive: false}, []contact.Contact{}},
	{Params{Today: may30, InNbDays: 1, Inclusive: false}, []contact.Contact{rob}},
	{Params{Today: may30, InNbDays: 2, Inclusive: false}, []contact.Contact{jane}},
	{Params{Today: may30, InNbDays: 3, Inclusive: false}, []contact.Contact{}},
	{Params{Today: may30, InNbDays: 0, Inclusive: true}, []contact.Contact{}},
	{Params{Today: may30, InNbDays: 1, Inclusive: true}, []contact.Contact{rob}},
	{Params{Today: may30, InNbDays: 2, Inclusive: true}, []contact.Contact{rob, jane}},
	{Params{Today: may30, InNbDays: 3, Inclusive: true}, []contact.Contact{rob, jane}},
	{Params{Today: may31, InNbDays: 0, Inclusive: false}, []contact.Contact{rob}},
}

var dataProviderOnSeveralYears = []dataprovider{
	{Params{Today: dec30, InNbDays: 0, Inclusive: false}, []contact.Contact{}},
	{Params{Today: dec30, InNbDays: 1, Inclusive: false}, []contact.Contact{jill}},
	{Params{Today: dec30, InNbDays: 2, Inclusive: false}, []contact.Contact{nick}},
	{Params{Today: dec30, InNbDays: 3, Inclusive: false}, []contact.Contact{}},
	{Params{Today: dec30, InNbDays: 0, Inclusive: true}, []contact.Contact{}},
	{Params{Today: dec30, InNbDays: 1, Inclusive: true}, []contact.Contact{jill}},
	{Params{Today: dec30, InNbDays: 2, Inclusive: true}, []contact.Contact{jill, nick}},
	{Params{Today: dec30, InNbDays: 3, Inclusive: true}, []contact.Contact{jill, nick}},
	{Params{Today: dec31, InNbDays: 0, Inclusive: false}, []contact.Contact{jill}},
}

func Test_ContactsToRemind(t *testing.T) {
	for _, data := range dataProvider {
		performUnitTest(data, t)
	}
}

func Test_ContactsToRemindSeveralMonths(t *testing.T) {
	for _, data := range dataProviderOnSeveralMonths {
		performUnitTest(data, t)
	}
}

func Test_ContactsToRemindSeveralYears(t *testing.T) {
	for _, data := range dataProviderOnSeveralYears {
		performUnitTest(data, t)
	}
}

func Test_should_remind_on_period_including_different_months(t *testing.T) {

	var aug30 = testdata.LocalDate(2018, time.August, 30)

	var john = contact.Contact{Name: "John", BirthDate: testdata.LocalDate(1981, time.August, 31)}
	var sara = contact.Contact{Name: "Sara", BirthDate: testdata.LocalDate(1974, time.September, 1)}

	var contacts = []contact.Contact{john, sara}

	var params = Params{Today: aug30, InNbDays: 3, Inclusive: true}

	reminder := Reminder{contacts, params}
	s := reminder.ContactsToRemind()

	assert.Equal(t, s, contacts)
}

func Test_should_remind_on_period_including_different_years(t *testing.T) {

	var dec30 = testdata.LocalDate(2018, time.December, 30)

	var john = contact.Contact{Name: "John", BirthDate: testdata.LocalDate(1990, time.December, 31)}
	var sara = contact.Contact{Name: "Sara", BirthDate: testdata.LocalDate(1989, time.January, 1)}

	var contacts = []contact.Contact{john, sara}

	var params = Params{Today: dec30, InNbDays: 3, Inclusive: true}

	reminder := Reminder{contacts, params}
	s := reminder.ContactsToRemind()

	assert.Equal(t, s, contacts)
}

func Test_should_remind_once_with_different_timezones(t *testing.T) {

	LAloc, _ := time.LoadLocation("America/Los_Angeles")
	SydneyLoc, _ := time.LoadLocation("Australia/Sydney")
	ParisLoc, _ := time.LoadLocation("Europe/Paris")

	aug20LA := time.Date(2018, time.August, 20, 0, 0, 0, 0, LAloc)

	remind20augLA := Reminder{
		contacts,
		Params{Today: aug20LA, InNbDays: 0, Inclusive: false},
	}.ContactsToRemind()
	assert.Equal(t, []contact.Contact{john}, remind20augLA)

	aug20Sydney := time.Date(2018, time.August, 20, 0, 0, 0, 0, SydneyLoc)

	remind20augSydney := Reminder{
		contacts,
		Params{Today: aug20Sydney, InNbDays: 0, Inclusive: false},
	}.ContactsToRemind()
	assert.Equal(t, []contact.Contact{john}, remind20augSydney)

	aug20Paris := time.Date(2018, time.August, 20, 0, 0, 0, 0, ParisLoc)

	remind20augParis := Reminder{
		contacts,
		Params{Today: aug20Paris, InNbDays: 0, Inclusive: false},
	}.ContactsToRemind()
	assert.Equal(t, []contact.Contact{john}, remind20augParis)
}

func performUnitTest(data dataprovider, t *testing.T) {
	testCase := fmt.Sprintf("[%s, %d, %t]", data.in.Today, data.in.InNbDays, data.in.Inclusive)
	t.Run(testCase, func(t *testing.T) {
		reminder := Reminder{data.out, data.in}
		s := reminder.ContactsToRemind()
		if len(data.out) == 0 {
			assert.Empty(t, s)
		} else {
			assert.Equal(t, s, data.out)
		}
	})
}
