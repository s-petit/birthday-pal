package main

import (
	"testing"
	"os"
	"github.com/s-petit/birthday-pal/email"
	"github.com/stretchr/testify/mock"
	"errors"
)

//TODO refacto sur le projet entier : privilegier les pointeurs
//TODO ajouter de la validation sur les args, notamment urls et emails
// https://goreportcard.com/report/github.com/vektra/mockery

func Test_main(t *testing.T) {
	//TODO how to test with mow.cli ?
	os.Args = []string{"", "https://mycarddav/com/contacts", "carddav-user", "carddav-pass", "recipient-email"}
/*	os.Args[1] = "https://mycarddav.com/contacts"
	os.Args[2] = "carddav-user"
	os.Args[3] = "carddav-pass"
	os.Args[4] = "recipient-email"*/

	client := FakeClient{}
	smtp := FakeSender{}

	client.On("Get").Return("lol")
	smtp.On("Send")

	DoIt(client, smtp, []string{"spe@mail.com", "wsh@prov.fr"})

	}

type FakeClient struct {
	mock.Mock
}

func (c FakeClient) Get() (string, error) {
	args := c.Called()
/*	contact := `
BEGIN:VCARD
VERSION:3.0
FN:Alexis Foo
BDAY:19831028
END:VCARD
BEGIN:VCARD
VERSION:3.0
FN:Florence Bar
BDAY:19860425
END:VCARD
`*/
	return args.String(0), errors.New("lol")
}

type FakeSender struct {
	mock.Mock
}

func (c FakeSender) Send(contact email.Contact, recipients []string)  {
	//assert.Equal(t, []string{"spe@mail.com", "wsh@prov.fr"}, recipients)
}