package main

import (
	"fmt"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/remind"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"
)

func handler() http.Handler {

	h := func(w http.ResponseWriter, r *http.Request) {
		vcard := `
BEGIN:VCARD
VERSION:3.0
FN:Alexis Foo
N:Foo;Alexis;;;
BDAY:%s
END:VCARD
`
		bday := vcardBday(time.Now())
		io.WriteString(w, fmt.Sprintf(vcard, bday))
	}

	r := http.NewServeMux()
	r.HandleFunc("/contact", h)
	return r
}

func Test_main_with_carddav(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	d := testdata.StartSMTPServer()
	defer d.Shutdown()

	os.Args = []string{"",
		"--smtp-host=localhost",
		"--smtp-port=2525",
		"--smtp-user=user@test",
		"--smtp-pass=smtp-pass",
		"carddav",
		fmt.Sprintf("%s/contact", srv.URL),
		"recipient@test",
	}

	assert.NotPanics(t, main)
}

func Test_remind_birthdays(t *testing.T) {

	contactProvider := new(fakeContactProvider)
	smtp := new(fakeSender)

	al := contact.Contact{Name: "Al Foo", BirthDate: testdata.BirthDate(1983, time.October, 28)}
	john := contact.Contact{Name: "John Bar", BirthDate: testdata.BirthDate(1986, time.August, 31)}
	con := []contact.Contact{al, john}

	recipients := []string{"spe@mail.com", "wsh@prov.fr"}

	contactToRemind := remind.ContactBirthday{Name: "John Bar", BirthDate: testdata.BirthDate(1986, time.August, 31), Age: 32}
	reminder := remind.Reminder{CurrentDate: testdata.LocalDate(2018, time.August, 30), NbDaysBeforeBDay: 1}

	contactProvider.On("GetContacts").Return(con, nil)
	smtp.On("Send", contactToRemind, recipients).Times(1)

	remindBirthdays(contactProvider, smtp, reminder, recipients)

	contactProvider.AssertExpectations(t)
	smtp.AssertExpectations(t)

}

func vcardBday(date time.Time) string {
	year := strconv.Itoa(date.Year())
	month := strconv.Itoa(int(date.Month()))
	day := strconv.Itoa(date.Day())

	if int(date.Month()) < 10 {
		month = "0" + month
	}
	if date.Day() < 10 {
		day = "0" + day
	}
	bday := year + month + day
	return bday
}

type fakeContactProvider struct {
	mock.Mock
}

func (c *fakeContactProvider) GetContacts() ([]contact.Contact, error) {
	args := c.Called()

	var s []contact.Contact
	var ok bool
	if s, ok = args.Get(0).([]contact.Contact); !ok {
		panic(fmt.Sprintf("assert: arguments: Int(%d) failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	return s, args.Error(1)
}

type fakeSender struct {
	mock.Mock
}

func (c *fakeSender) Send(contact remind.ContactBirthday, recipients []string) error {
	c.Called(contact, recipients)
	return nil
}
