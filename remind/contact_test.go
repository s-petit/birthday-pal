package remind

import (
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

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

	contactsToRemind := ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug21, NbDaysBeforeBDay: bdaysInTwodays, EveryDayUntilBDay: false})
	assert.Equal(t, []ContactBirthday{johnAge, saraAge}, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug22, NbDaysBeforeBDay: bdaysInTwodays, EveryDayUntilBDay: false})
	assert.Equal(t, []ContactBirthday{johnAge, saraAge}, contactsToRemind)
	contactsToRemind = ContactsToRemind([]contact.Contact{john, sara}, Reminder{CurrentDate: aug23, NbDaysBeforeBDay: bdaysInTwodays, EveryDayUntilBDay: false})
	assert.Equal(t, []ContactBirthday{saraAge}, contactsToRemind)
}
