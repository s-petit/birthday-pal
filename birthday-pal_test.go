package main

import (
	"fmt"
	"github.com/s-petit/birthday-pal/remind"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

//TODO refacto sur le projet entier : privilegier les pointeurs
//TODO ajouter de la validation sur les args, notamment urls et emails

func handler() http.Handler {

	h := func(w http.ResponseWriter, r *http.Request) {
		vcard := `
BEGIN:VCARD
VERSION:3.0
FN:Alexis Foo
N:Foo;Alexis;;;
BDAY:19831028
END:VCARD
`
		io.WriteString(w, vcard)
	}

	r := http.NewServeMux()
	r.HandleFunc("/contact", h)
	return r
}

func Test_main(t *testing.T) {

	srv := httptest.NewServer(handler())
	defer srv.Close()

	d := testdata.StartSMTPServer()

	os.Args = []string{"", fmt.Sprintf("%s/contact", srv.URL), "carddav-user", "carddav-pass", "1", "false", "localhost", "2525", "smtp-user", "smtp-pass", "recipient-email"}

	main()

	// TODO add an assertion
	// TODO why the test does not crash with a wrong smtp creds ?

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

	remindBirthdays(client, smtp, recipients, reminder)

	client.AssertExpectations(t)
	smtp.AssertExpectations(t)

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
