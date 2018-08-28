package main

import (
	"github.com/s-petit/birthday-pal/vcardparser"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

//TODO refacto sur le projet entier : privilegier les pointeurs
//TODO ajouter de la validation sur les args, notamment urls et emails
// https://goreportcard.com/report/github.com/vektra/mockery
//TODO move remindcontact elsewhere

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
BDAY:19860829
END:VCARD
`

	recipients := []string{"spe@mail.com", "wsh@prov.fr"}
	contact := vcardparser.RemindContact{Name: "Florence Bar", BirthDate: time.Date(1986, time.August, 29, 0, 0, 0, 0, time.UTC), Age: 31}

	client.On("Get").Return(vcards)
	smtp.On("Send", contact, recipients).Times(1)

	remindBirthdays(client, smtp, recipients, 1)

	client.AssertExpectations(t)
	smtp.AssertExpectations(t)

}

//TODO implemter un struct simplifie de birthdate avec mois et annee plutot que time.

type FakeClient struct {
	mock.Mock
}

func (c FakeClient) Get() (string, error) {
	args := c.Called()
	return args.String(0), nil
}

type FakeSender struct {
	mock.Mock
}

func (c FakeSender) Send(contact vcardparser.RemindContact, recipients []string) {
	c.Called(contact, recipients)
}
