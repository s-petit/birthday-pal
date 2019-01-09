package remind

import (
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


func Test_lol(t *testing.T) {
	aug21 := testdata.LocalDate(2018, time.August, 21)
	aug20 := testdata.LocalDate(2018, time.August, 20)
	aug19 := testdata.LocalDate(2018, time.August, 19)

	johnBday := aug20
	saraBday := aug21

	remindJohn := Reminder{CurrentDate: aug19, NbDaysBeforeBDay: 0, EveryDayUntilBDay: false}.remindOnce(johnBday)
	remindSara := Reminder{CurrentDate: aug19, NbDaysBeforeBDay: 0, EveryDayUntilBDay: false}.remindOnce(saraBday)

	assert.Equal(t, false, remindJohn)
	assert.Equal(t, false, remindSara)

}


//TODO SPE faire en sorte que tous les tests soient sur une date fixe, il faut jouer uniquement sur le -d -e -i

// comme celui ci par exemple
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

func Test_should_remind_all_birthdays_in_the_next_two_days(t *testing.T) {
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

	aug22LA := time.Date(2018, time.August, 22, 0, 0, 0, 0, LAloc)

	remind := Reminder{CurrentDate: aug22LA, NbDaysBeforeBDay: 0, EveryDayUntilBDay: false}.remindOnce(birthday)
	assert.Equal(t, true, remind)

	aug22Sydney := time.Date(2018, time.August, 22, 0, 0, 0, 0, SydneyLoc)

	remind = Reminder{CurrentDate: aug22Sydney, NbDaysBeforeBDay: 0, EveryDayUntilBDay: false}.remindOnce(birthday)
	assert.Equal(t, true, remind)

	aug22Paris := time.Date(2018, time.August, 22, 0, 0, 0, 0, ParisLoc)

	remind = Reminder{CurrentDate: aug22Paris, NbDaysBeforeBDay: 0, EveryDayUntilBDay: false}.remindOnce(birthday)
	assert.Equal(t, true, remind)
}

func Test_should_remind_when_current_day_is_a_birthday(t *testing.T) {
	birthday := testdata.BirthDate(2016, time.August, 22)
	aug22 := testdata.LocalDate(2018, time.August, 22)

	remind := Reminder{CurrentDate: aug22, NbDaysBeforeBDay: 0, EveryDayUntilBDay: false}.ShouldRemind(birthday)
	assert.Equal(t, true, remind)

	remind2 := Reminder{CurrentDate: aug22, NbDaysBeforeBDay: 0, EveryDayUntilBDay: true}.ShouldRemind(birthday)
	assert.Equal(t, true, remind2)
}

/*
func Test_remind_contacts(t *testing.T) {
	date := testdata.LocalDate(2018, time.November, 21)
	c := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.November, 23)}
	c2 := contact.Contact{Name: "Bill", BirthDate: testdata.BirthDate(1980, time.November, 28)}

	contactsToRemind := ContactsToRemind([]contact.Contact{c, c2}, Reminder{CurrentDate: date, NbDaysBeforeBDay: 2})

	expected := ContactBirthday{Contact: c, Age: 38}

	assert.Equal(t, []ContactBirthday{expected}, contactsToRemind)
}

func Test_remind_contacts_birthday_current_day(t *testing.T) {

	aug23 := testdata.LocalDate(2018, time.August, 23)
	aug22 := testdata.LocalDate(2018, time.August, 22)
	aug21 := testdata.LocalDate(2018, time.August, 21)

	john := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.August, 22)}
	sara := contact.Contact{Name: "Sara", BirthDate: testdata.BirthDate(1984, time.August, 23)}

	johnAge := ContactBirthday{Contact: john, Age: 38}
	saraAge := ContactBirthday{Contact: sara, Age: 34}

	bdaysOfTheDay := 0
	contactsToRemind := ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug21, NbDaysBeforeBDay: bdaysOfTheDay, EveryDayUntilBDay: false})
	assert.Empty(t, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug22, NbDaysBeforeBDay: bdaysOfTheDay, EveryDayUntilBDay: false})
	assert.Equal(t, []ContactBirthday{johnAge}, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug23, NbDaysBeforeBDay: bdaysOfTheDay, EveryDayUntilBDay: false})
	assert.Equal(t, []ContactBirthday{saraAge}, contactsToRemind)
}

//TODO SPE renommer tests
func Test_remind_contacts_birthday_current_day_does_not_depend_on_EveryDayUntilBDay(t *testing.T) {

	aug23 := testdata.LocalDate(2018, time.August, 23)
	aug22 := testdata.LocalDate(2018, time.August, 22)
	aug21 := testdata.LocalDate(2018, time.August, 21)

	john := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.August, 22)}
	sara := contact.Contact{Name: "Sara", BirthDate: testdata.BirthDate(1984, time.August, 23)}

	johnAge := ContactBirthday{Contact: john, Age: 38}
	saraAge := ContactBirthday{Contact: sara, Age: 34}

	bdaysOfTheDay := 0
	contactsToRemind := ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug21, NbDaysBeforeBDay: bdaysOfTheDay, EveryDayUntilBDay: true})
	assert.Empty(t, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug22, NbDaysBeforeBDay: bdaysOfTheDay, EveryDayUntilBDay: true})
	assert.Equal(t, []ContactBirthday{johnAge}, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug23, NbDaysBeforeBDay: bdaysOfTheDay, EveryDayUntilBDay: true})
	assert.Equal(t, []ContactBirthday{saraAge}, contactsToRemind)
}

func Test_remind_contacts_birthday_in_one_day(t *testing.T) {

	aug23 := testdata.LocalDate(2018, time.August, 23)
	aug22 := testdata.LocalDate(2018, time.August, 22)
	aug21 := testdata.LocalDate(2018, time.August, 21)

	john := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.August, 22)}
	sara := contact.Contact{Name: "Sara", BirthDate: testdata.BirthDate(1984, time.August, 23)}

	johnAge := ContactBirthday{Contact: john, Age: 38}
	saraAge := ContactBirthday{Contact: sara, Age: 34}

	bdaysTomorrow := 1

	contactsToRemind := ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug21, NbDaysBeforeBDay: bdaysTomorrow, EveryDayUntilBDay: false})
	assert.Equal(t, []ContactBirthday{johnAge}, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug22, NbDaysBeforeBDay: bdaysTomorrow, EveryDayUntilBDay: false})
	assert.Equal(t, []ContactBirthday{saraAge}, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug23, NbDaysBeforeBDay: bdaysTomorrow, EveryDayUntilBDay: false})
	assert.Empty(t, contactsToRemind)
}

func Test_remind_all_contacts_birthday_from_today_to_tomorrow(t *testing.T) {

	aug23 := testdata.LocalDate(2018, time.August, 23)
	aug22 := testdata.LocalDate(2018, time.August, 22)
	aug21 := testdata.LocalDate(2018, time.August, 21)

	john := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.August, 22)}
	sara := contact.Contact{Name: "Sara", BirthDate: testdata.BirthDate(1984, time.August, 23)}

	johnAge := ContactBirthday{Contact: john, Age: 38}
	saraAge := ContactBirthday{Contact: sara, Age: 34}

	bdaysTomorrow := 1

	contactsToRemind := ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug21, NbDaysBeforeBDay: bdaysTomorrow, EveryDayUntilBDay: true})
	assert.Equal(t, []ContactBirthday{johnAge}, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug22, NbDaysBeforeBDay: bdaysTomorrow, EveryDayUntilBDay: true})
	assert.Equal(t, []ContactBirthday{johnAge, saraAge}, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug23, NbDaysBeforeBDay: bdaysTomorrow, EveryDayUntilBDay: true})
	assert.Equal(t, []ContactBirthday{saraAge}, contactsToRemind)
}

func Test_remind_contacts_birthday_in_two_days(t *testing.T) {

	aug23 := testdata.LocalDate(2018, time.August, 23)
	aug22 := testdata.LocalDate(2018, time.August, 22)
	aug21 := testdata.LocalDate(2018, time.August, 21)

	john := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.August, 22)}
	sara := contact.Contact{Name: "Sara", BirthDate: testdata.BirthDate(1984, time.August, 23)}

	saraAge := ContactBirthday{Contact: sara, Age: 34}

	bdaysInTwodays := 2

	contactsToRemind := ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug21, NbDaysBeforeBDay: bdaysInTwodays, EveryDayUntilBDay: false})
	assert.Equal(t, []ContactBirthday{saraAge}, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug22, NbDaysBeforeBDay: bdaysInTwodays, EveryDayUntilBDay: false})
	assert.Empty(t, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug23, NbDaysBeforeBDay: bdaysInTwodays, EveryDayUntilBDay: false})
	assert.Empty(t, contactsToRemind)
}


func Test_remind_all_contacts_birthday_in_the_next_two_days(t *testing.T) {

	aug23 := testdata.LocalDate(2018, time.August, 23)
	aug22 := testdata.LocalDate(2018, time.August, 22)
	aug21 := testdata.LocalDate(2018, time.August, 21)

	john := contact.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.August, 22)}
	sara := contact.Contact{Name: "Sara", BirthDate: testdata.BirthDate(1984, time.August, 23)}

	johnAge := ContactBirthday{Contact: john, Age: 38}
	saraAge := ContactBirthday{Contact: sara, Age: 34}

	bdaysInTwodays := 2

	contactsToRemind := ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug21, NbDaysBeforeBDay: bdaysInTwodays, EveryDayUntilBDay: true})
	assert.Equal(t, []ContactBirthday{johnAge, saraAge}, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug22, NbDaysBeforeBDay: bdaysInTwodays, EveryDayUntilBDay: true})
	assert.Equal(t, []ContactBirthday{johnAge, saraAge}, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug23, NbDaysBeforeBDay: bdaysInTwodays, EveryDayUntilBDay: true})
	assert.Equal(t, []ContactBirthday{saraAge}, contactsToRemind)
}
*/
