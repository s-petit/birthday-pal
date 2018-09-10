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

	expected := ContactBirthday{Name: "John", BirthDate: testdata.BirthDate(1980, time.November, 23), Age: 38}

	assert.Equal(t, []ContactBirthday{expected}, contactsToRemind)
}
