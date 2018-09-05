package main

import (
	"fmt"
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

func Test_main(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	d := testdata.StartSMTPServer()

	os.Args = []string{"",
		fmt.Sprintf("--carddav-url=%s/contact", srv.URL),
		"--smtp-host=localhost",
		"--smtp-port=2525",
		"--smtp-user=user@test",
		"--smtp-pass=smtp-pass",
		"recipient@test",
	}

	assert.NotPanics(t, main)

	d.Shutdown()
}

func Test_remind_birthdays(t *testing.T) {

	client := new(fakeClient)
	smtp := new(fakeSender)

	vcards := `
BEGIN:VCARD
VERSION:3.0
FN:Alexis Foo
BDAY:19831028
END:VCARD
BEGIN:VCARD
VERSION:3.0
FN:John Bar
BDAY:19860831
END:VCARD
`

	recipients := []string{"spe@mail.com", "wsh@prov.fr"}

	contactToRemind := remind.ContactBirthday{Name: "John Bar", BirthDate: testdata.BirthDate(1986, time.August, 31), Age: 32}
	reminder := remind.Reminder{CurrentDate: testdata.LocalDate(2018, time.August, 30), NbDaysBeforeBDay: 1}

	client.On("Get").Return(vcards, nil)
	smtp.On("Send", contactToRemind, recipients).Times(1)

	remindBirthdays(client, smtp, reminder, recipients)

	client.AssertExpectations(t)
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

type fakeClient struct {
	mock.Mock
}

func (c *fakeClient) Get() (string, error) {
	args := c.Called()
	return args.String(0), args.Error(1)
}

type fakeSender struct {
	mock.Mock
}

func (c *fakeSender) Send(contact remind.ContactBirthday, recipients []string) error {
	c.Called(contact, recipients)
	return nil
}
