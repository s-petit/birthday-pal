package email

import (
	"github.com/s-petit/birthday-pal/app/contact"
	"github.com/s-petit/birthday-pal/app/remind"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"net/smtp"
	"testing"
	"time"
)

func Test_hostPort(t *testing.T) {
	sender := SMTPClient{Host: "localhost", Port: 2525}
	hostPort := sender.hostPort()
	assert.Equal(t, "localhost:2525", hostPort)
}

func Test_auth_when_username_and_password_not_empty(t *testing.T) {
	sender := SMTPClient{Host: "localhost", Port: 2525, Username: "user", Password: "pwd"}
	auth := sender.auth()
	assert.Equal(t, smtp.PlainAuth("", "user", "pwd", "localhost"), auth)
}

func Test_auth_should_be_nil_when_username_empty(t *testing.T) {
	sender := SMTPClient{Host: "localhost", Port: 2525, Password: "pwd"}
	auth := sender.auth()
	assert.Equal(t, nil, auth)
}

func Test_auth_should_be_nil_when_password_empty(t *testing.T) {
	sender := SMTPClient{Host: "localhost", Port: 2525, Username: "user"}
	auth := sender.auth()
	assert.Equal(t, nil, auth)
}

func Test_auth_should_be_nil_when_username_and_password_empty(t *testing.T) {
	sender := SMTPClient{Host: "localhost", Port: 2525}
	auth := sender.auth()
	assert.Equal(t, nil, auth)
}

func Test_send(t *testing.T) {
	d := testdata.StartSMTPServer()

	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	sender := SMTPClient{Host: "localhost", Port: 2525}
	c := Contacts{
		Contacts: []contact.Contact{{Name: "ttf", BirthDate: birthday}}, RemindParams: remind.Params{},
	}
	e := sender.Send(c, []string{"recipient@test", "recipient2@test"})
	assert.NoError(t, e)
	d.Shutdown()
}

func Test_no_send_when_contact_list_empty(t *testing.T) {
	d := testdata.StartSMTPServer()

	sender := SMTPClient{Host: "localhost", Port: 2525}
	c := Contacts{
		Contacts: []contact.Contact{}, RemindParams: remind.Params{},
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
		Contacts: []contact.Contact{{Name: "ttf", BirthDate: birthday}}, RemindParams: remind.Params{},
	}
	e := sender.Send(c, []string{"recipient@test2"})
	assert.EqualError(t, e, "454 4.1.1 Error: Relay access denied: test2")
	d.Shutdown()
}
