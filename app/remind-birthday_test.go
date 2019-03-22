package app

import (
	"errors"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/email"
	"github.com/s-petit/birthday-pal/remind"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func Test_remind_birthdays_successful(t *testing.T) {

	contactProvider := new(fakeContactProvider)
	smtp := new(fakeSender)

	al := contact.Contact{Name: "Al Foo", BirthDate: testdata.BirthDate(1983, time.October, 28)}
	john := contact.Contact{Name: "John Bar", BirthDate: testdata.BirthDate(1986, time.August, 31)}
	con := []contact.Contact{al, john}

	recipients := []string{"spe@mail.com", "wsh@prov.fr"}

	remindParams := remind.Params{Today: testdata.LocalDate(2018, time.August, 30), InNbDays: 1}

	contactProvider.On("GetContacts").Return(con, nil)
	emailContacts := email.Contacts{Contacts: []contact.Contact{john}, RemindParams: remindParams}
	smtp.On("Send", emailContacts, recipients).Return(nil)

	err := BirthdayPal{}.Exec(contactProvider, smtp, remindParams, recipients)

	assert.NoError(t, err)
	contactProvider.AssertExpectations(t)
	smtp.AssertExpectations(t)
	smtp.AssertNumberOfCalls(t, "Send", 1)
}

func Test_remind_birthdays_fail_during_contact_retrieving(t *testing.T) {

	contactProvider := new(fakeContactProvider)
	smtp := new(fakeSender)

	contactProvider.On("GetContacts").Return([]contact.Contact{}, errors.New("woops"))

	err := BirthdayPal{}.Exec(contactProvider, smtp, remind.Params{}, []string{})

	assert.Error(t, err)
	contactProvider.AssertExpectations(t)
	smtp.AssertNumberOfCalls(t, "Send", 0)
}

func Test_remind_birthdays_fail_during_mail_sending(t *testing.T) {

	contactProvider := new(fakeContactProvider)
	smtp := new(fakeSender)

	al := contact.Contact{Name: "Al Foo", BirthDate: testdata.BirthDate(1983, time.October, 28)}
	john := contact.Contact{Name: "John Bar", BirthDate: testdata.BirthDate(1986, time.August, 31)}
	con := []contact.Contact{al, john}

	recipients := []string{"spe@mail.com", "wsh@prov.fr"}

	remindParams := remind.Params{Today: testdata.LocalDate(2018, time.August, 30), InNbDays: 1}
	emailContacts := email.Contacts{Contacts: []contact.Contact{john}, RemindParams: remindParams}

	contactProvider.On("GetContacts").Return(con, nil)
	smtp.On("Send", emailContacts, recipients).Return(errors.New("wow"))

	err := BirthdayPal{}.Exec(contactProvider, smtp, remindParams, recipients)

	assert.Error(t, err)
	contactProvider.AssertExpectations(t)
	smtp.AssertExpectations(t)
}

type fakeContactProvider struct {
	mock.Mock
}

func (c *fakeContactProvider) GetContacts() ([]contact.Contact, error) {
	args := c.Called()
	return args.Get(0).([]contact.Contact), args.Error(1)
}

type fakeSender struct {
	mock.Mock
}

func (c *fakeSender) Send(contact email.Contacts, recipients []string) error {
	args := c.Called(contact, recipients)
	return args.Error(0)
}
