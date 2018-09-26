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
Subject: Anniversaire de John -34 an(s)-

Ce sera l'anniversaire de John le 22/08. Pensez à le lui souhaiter!`

	birthday := time.Date(1980, time.August, 22, 0, 0, 0, 0, time.UTC)

	bytes, _ := toMail(remind.ContactBirthday{"John", birthday, 34}, "fr")
	assert.Equal(t, expectedMail, string(bytes))
}

func Test_should_get_mail_in_french_without_age(t *testing.T) {

	expectedMail :=
`To: Birthday Pals
Subject: Anniversaire de John

Ce sera l'anniversaire de John le 22/08. Pensez à le lui souhaiter!`

	birthday := time.Date(0, time.August, 22, 0, 0, 0, 0, time.UTC)

	bytes, _ := toMail(remind.ContactBirthday{"John", birthday, 34}, "fr")
	assert.Equal(t, expectedMail, string(bytes))
}

func Test_should_get_mail_in_english(t *testing.T) {

	expectedMail :=
`To: Birthday Pals
Subject: John's birthday -34 yo-

The 08/22 will be John's birthday. Do not forget to make your wish!`

	birthday := time.Date(1980, time.August, 22, 0, 0, 0, 0, time.UTC)

	bytes, _ := toMail(remind.ContactBirthday{"John", birthday, 34}, "EN")
	assert.Equal(t, expectedMail, string(bytes))
}

func Test_should_throw_error_when_subject_template_malformed(t *testing.T) {

	tmpl := new(fakeTemplate)
	tmpl.On("subject").Return("{{{{{")

	bytes, err := resolveMail(remind.ContactBirthday{}, tmpl)

	assert.Error(t, err)
	assert.Empty(t, bytes)
}

func Test_should_throw_error_when_body_template_malformed(t *testing.T) {

	tmpl := new(fakeTemplate)
	tmpl.On("subject").Return("subject")
	tmpl.On("body").Return("{{{{{")

	bytes, err := resolveMail(remind.ContactBirthday{}, tmpl)

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
