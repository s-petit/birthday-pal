package remind

import (
	"fmt"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var aug21 = testdata.LocalDate(2018, time.August, 21)
var aug20 = testdata.LocalDate(2018, time.August, 20)
var aug19 = testdata.LocalDate(2018, time.August, 19)

var john = contact.Contact{Name: "John", BirthDate: aug20}
var sara = contact.Contact{Name: "Sara", BirthDate: aug21}

var contacts = []contact.Contact{john, sara}

var dataProvider = []struct {
	in  Params
	out []contact.Contact
}{
	{Params{CurrentDate: aug19, InNbDays: 0, Inclusive: false}, []contact.Contact{}},
	{Params{CurrentDate: aug19, InNbDays: 1, Inclusive: false}, []contact.Contact{john}},
	{Params{CurrentDate: aug19, InNbDays: 2, Inclusive: false}, []contact.Contact{sara}},
	{Params{CurrentDate: aug19, InNbDays: 3, Inclusive: false}, []contact.Contact{}},
	{Params{CurrentDate: aug19, InNbDays: 0, Inclusive: true}, []contact.Contact{}},
	{Params{CurrentDate: aug19, InNbDays: 1, Inclusive: true}, []contact.Contact{john}},
	{Params{CurrentDate: aug19, InNbDays: 2, Inclusive: true}, []contact.Contact{john, sara}},
	{Params{CurrentDate: aug19, InNbDays: 3, Inclusive: true}, []contact.Contact{john, sara}},
	{Params{CurrentDate: aug20, InNbDays: 0, Inclusive: false}, []contact.Contact{john}},
}

func Test_ContactsToRemind(t *testing.T) {
	for _, data := range dataProvider {
		testCase := fmt.Sprintf("[%s, %d, %t]", data.in.CurrentDate, data.in.InNbDays, data.in.Inclusive)
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
}

func Test_should_remind_once_with_different_timezones(t *testing.T) {

	LAloc, _ := time.LoadLocation("America/Los_Angeles")
	SydneyLoc, _ := time.LoadLocation("Australia/Sydney")
	ParisLoc, _ := time.LoadLocation("Europe/Paris")

	aug20LA := time.Date(2018, time.August, 20, 0, 0, 0, 0, LAloc)

	remind20augLA := Reminder{
		contacts,
		Params{CurrentDate: aug20LA, InNbDays: 0, Inclusive: false},
	}.ContactsToRemind()
	assert.Equal(t, []contact.Contact{john}, remind20augLA)

	aug20Sydney := time.Date(2018, time.August, 20, 0, 0, 0, 0, SydneyLoc)

	remind20augSydney := Reminder{
		contacts,
		Params{CurrentDate: aug20Sydney, InNbDays: 0, Inclusive: false},
	}.ContactsToRemind()
	assert.Equal(t, []contact.Contact{john}, remind20augSydney)

	aug20Paris := time.Date(2018, time.August, 20, 0, 0, 0, 0, ParisLoc)

	remind20augParis := Reminder{
		contacts,
		Params{CurrentDate: aug20Paris, InNbDays: 0, Inclusive: false},
	}.ContactsToRemind()
	assert.Equal(t, []contact.Contact{john}, remind20augParis)
}
