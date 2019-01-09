package email

import (
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_hostPort(t *testing.T) {
	sender := SMTPClient{Host: "localhost", Port: 2525}
	hostPort := sender.hostPort()
	assert.Equal(t, "localhost:2525", hostPort)
}

func Test_send(t *testing.T) {
	d := testdata.StartSMTPServer()

	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	sender := SMTPClient{Host: "localhost", Port: 2525}
	c := Contacts{
		Contacts: []contact.Contact{{Name: "ttf", BirthDate: birthday}}, RemindDate: birthday,
	}
	e := sender.Send(c, []string{"recipient@test", "recipient2@test"})
	assert.NoError(t, e)
	d.Shutdown()
}

func Test_no_send_when_contact_list_empty(t *testing.T) {
	d := testdata.StartSMTPServer()

	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	sender := SMTPClient{Host: "localhost", Port: 2525}
	c := Contacts{
		Contacts: []contact.Contact{}, RemindDate: birthday,
	}
	e := sender.Send(c, []string{"recipient@test", "recipient2@test"})
	assert.NoError(t, e)
	d.Shutdown()
}

func Test_send_error(t *testing.T) {
	d := testdata.StartSMTPServer()

	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	sender := SMTPClient{Host: "localhost", Port: 2525}
	c := Contacts{
		Contacts: []contact.Contact{{Name: "ttf", BirthDate: birthday}}, RemindDate: birthday,
	}
	e := sender.Send(c, []string{"recipient@test2"})
	assert.EqualError(t, e, "454 4.1.1 Error: Relay access denied: test2")
	d.Shutdown()
}
