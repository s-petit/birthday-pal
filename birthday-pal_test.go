package main

import (
	"github.com/s-petit/birthday-pal/contact"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

//TODO refacto sur le projet entier : privilegier les pointeurs
//TODO ajouter de la validation sur les args, notamment urls et emails
// https://goreportcard.com/report/github.com/vektra/mockery

//TODO faire un test moins fragile, moins dependant de la date courante
func Test_remind_birthdays(t *testing.T) {
	//os.Args = []string{"", "https://mycarddav/com/contacts", "carddav-user", "carddav-pass", "recipient-email"}
	/*	os.Args[1] = "https://mycarddav.com/contacts"
		os.Args[2] = "carddav-user"
		os.Args[3] = "carddav-pass"
		os.Args[4] = "recipient-email"*/

	client := new(FakeClient)
	smtp := new(FakeSender)

	vcards := `
BEGIN:VCARD
VERSION:3.0
FN:Alexis Foo
BDAY:19831028
END:VCARD
BEGIN:VCARD
VERSION:3.0
FN:Florence Bar
BDAY:19860831
END:VCARD
`

	recipients := []string{"spe@mail.com", "wsh@prov.fr"}
	c := contact.Contact{Name: "Florence Bar", BirthDate: birthDate(1986, time.August, 31)}

	client.On("Get").Return(vcards, nil)
	smtp.On("Send", c, recipients).Times(1)

	remindBirthdays(client, smtp, recipients, 1)

	client.AssertExpectations(t)
	smtp.AssertExpectations(t)

}

//TODO SPE dans go a t on un scope test comme maven ?
func birthDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

type FakeClient struct {
	mock.Mock
}

func (c *FakeClient) Get() (string, error) {
	args := c.Called()
	return args.String(0), args.Error(1)
}

type FakeSender struct {
	mock.Mock
}

func (c *FakeSender) Send(contact contact.Contact, recipients []string) error {
	c.Called(contact, recipients)
	return nil
}
