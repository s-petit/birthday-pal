package email

import (
	"github.com/s-petit/birthday-pal/remind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func Test_should_get_mail_in_french(t *testing.T) {

	expectedMail :=
		`To: Birthday Pals
Subject: Anniversaires du 22/08

Le 22/08, n'oubliez pas de souhaiter l'anniversaire de :

- John (34 ans)
`

	birthday := time.Date(1980, time.August, 22, 0, 0, 0, 0, time.UTC)

	bytes, err := toMail(Contacts{
		Contacts: []remind.ContactBirthday{{"John", birthday, 34}}, RemindDate: birthday,
	}, "fr")
	assert.NoError(t, err)
	assert.Equal(t, expectedMail, string(bytes))
}

func Test_should_get_mail_in_french_for_several_contacts(t *testing.T) {

	expectedMail :=
		`To: Birthday Pals
Subject: Anniversaires du 22/08

Le 22/08, n'oubliez pas de souhaiter l'anniversaire de :

- John (34 ans)

- Jane (23 ans)

- Jill (2 ans)
`

	birthday := time.Date(1980, time.August, 22, 0, 0, 0, 0, time.UTC)

	bytes, err := toMail(Contacts{
		Contacts: []remind.ContactBirthday{
			{"John", birthday, 34},
			{"Jane", birthday, 23},
			{"Jill", birthday, 2},
		}, RemindDate: birthday,
	}, "fr")
	assert.NoError(t, err)
	assert.Equal(t, expectedMail, string(bytes))
}

func Test_should_get_mail_in_french_without_age(t *testing.T) {

	expectedMail :=
		`To: Birthday Pals
Subject: Anniversaires du 22/08

Le 22/08, n'oubliez pas de souhaiter l'anniversaire de :

- John
`

	birthday := time.Date(0, time.August, 22, 0, 0, 0, 0, time.UTC)

	bytes, err := toMail(Contacts{
		Contacts: []remind.ContactBirthday{{"John", birthday, 34}}, RemindDate: birthday,
	}, "fr")
	assert.NoError(t, err)
	assert.Equal(t, expectedMail, string(bytes))
}

func Test_should_get_mail_in_english(t *testing.T) {

	expectedMail :=
		`To: Birthday Pals
Subject: Your 08/22 birthday reminder

The 08/22, don't forget to wish birthdays of :

- John (34 yo)
`

	birthday := time.Date(1980, time.August, 22, 0, 0, 0, 0, time.UTC)

	bytes, err := toMail(Contacts{
		Contacts: []remind.ContactBirthday{{"John", birthday, 34}}, RemindDate: birthday,
	}, "EN")
	assert.NoError(t, err)
	assert.Equal(t, expectedMail, string(bytes))
}

func Test_should_throw_error_when_subject_template_malformed(t *testing.T) {

	tmpl := new(fakeTemplate)
	tmpl.On("simpleReminderSubject").Return("{{{{{")

	bytes, err := resolveMail(Contacts{}, tmpl)

	assert.Error(t, err)
	assert.Empty(t, bytes)
}

func Test_should_throw_error_when_body_template_malformed(t *testing.T) {

	tmpl := new(fakeTemplate)
	tmpl.On("simpleReminderSubject").Return("simpleReminderSubject")
	tmpl.On("simpleReminderBody").Return("{{{{{")

	bytes, err := resolveMail(Contacts{}, tmpl)

	assert.Error(t, err)
	assert.Empty(t, bytes)
}

type fakeTemplate struct {
	mock.Mock
}

func (ft fakeTemplate) subject() string {
	args := ft.Called()
	return args.String(0)
}

func (ft fakeTemplate) body() string {
	args := ft.Called()
	return args.String(0)
}

func (ft fakeTemplate) dateLayout() string {
	return "02/01"
}

func (ft fakeTemplate) formatDate(date time.Time) string {
	return "dasDate"
}
