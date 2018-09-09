package remind

import (
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/s-petit/birthday-pal/contact/vcard"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_remind_contacts(t *testing.T) {
	date := testdata.LocalDate(2018, time.November, 21)
	c := vcard.Contact{Name: "John", BirthDate: testdata.BirthDate(1980, time.November, 23)}
	c2 := vcard.Contact{Name: "Bill", BirthDate: testdata.BirthDate(1980, time.November, 28)}

	contactsToRemind := ContactsToRemind([]vcard.Contact{c, c2}, Reminder{CurrentDate: date, NbDaysBeforeBDay: 2})

	expected := ContactBirthday{Name: "John", BirthDate: testdata.BirthDate(1980, time.November, 23), Age: 38}

	assert.Equal(t, []ContactBirthday{expected}, contactsToRemind)
}
