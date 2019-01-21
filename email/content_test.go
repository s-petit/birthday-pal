package email

import (
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/remind"
	"github.com/s-petit/birthday-pal/testdata"
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

	birthday := testdata.BirthDate(1980, time.August, 22)
	currentDate := testdata.LocalDate(2014, time.August, 21)
	remindParams := remind.Params{CurrentDate: currentDate, InNbDays: 1}

	bytes, err := toMail(Contacts{
		Contacts: []contact.Contact{{"John", birthday}}, RemindParams: remindParams,
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

	johnBirthday := testdata.BirthDate(1980, time.August, 22)
	janeBirthday := testdata.BirthDate(1991, time.August, 22)
	jillBirthday := testdata.BirthDate(2012, time.August, 22)
	currentDate := testdata.LocalDate(2014, time.August, 21)
	remindParams := remind.Params{CurrentDate: currentDate, InNbDays: 1}

	bytes, err := toMail(Contacts{
		Contacts: []contact.Contact{
			{"John", johnBirthday},
			{"Jane", janeBirthday},
			{"Jill", jillBirthday},
		}, RemindParams: remindParams,
	}, "fr")
	assert.NoError(t, err)
	assert.Equal(t, expectedMail, string(bytes))
}

func Test_should_get_mail_in_french_for_several_contacts_and_several_days(t *testing.T) {

	expectedMail :=
		`To: Birthday Pals
Subject: Anniversaires du 22/08

Durant les 7 prochains jours, n'oubliez pas de souhaiter l'anniversaire de :

- John (34 ans le 21/08)

- Jane (23 ans le 18/08)

- Jill (2 ans le 15/08)
`

	johnBirthday := testdata.BirthDate(1980, time.August, 21)
	janeBirthday := testdata.BirthDate(1991, time.August, 18)
	jillBirthday := testdata.BirthDate(2012, time.August, 15)

	currentDate := testdata.LocalDate(2014, time.August, 21)
	remindParams := remind.Params{CurrentDate: currentDate, InNbDays: 1, Inclusive: true}

	bytes, err := toMail(Contacts{
		Contacts: []contact.Contact{
			{"John", johnBirthday},
			{"Jane", janeBirthday},
			{"Jill", jillBirthday},
		}, RemindParams: remindParams,
	}, "fr")
	assert.NoError(t, err)
	assert.Equal(t, expectedMail, string(bytes))
}

func Test_should_get_mail_in_english_for_several_contacts_and_several_days(t *testing.T) {

	expectedMail :=
		`To: Birthday Pals
Subject: Your 08/22 birthday reminder

During the next 7 days, don't forget to wish birthdays of :

- John (34 yo the 08/21)

- Jane (23 yo the 08/18)

- Jill (2 yo the 08/15)
`

	johnBirthday := testdata.BirthDate(1980, time.August, 21)
	janeBirthday := testdata.BirthDate(1991, time.August, 18)
	jillBirthday := testdata.BirthDate(2012, time.August, 15)

	currentDate := testdata.LocalDate(2014, time.August, 21)
	remindParams := remind.Params{CurrentDate: currentDate, InNbDays: 1, Inclusive: true}

	bytes, err := toMail(Contacts{
		Contacts: []contact.Contact{
			{"John", johnBirthday},
			{"Jane", janeBirthday},
			{"Jill", jillBirthday},
		}, RemindParams: remindParams,
	}, "en")
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

	invalidBirthday := testdata.BirthDate(0, time.August, 22)
	currentDate := testdata.LocalDate(2014, time.August, 21)
	remindParams := remind.Params{CurrentDate: currentDate, InNbDays: 1}

	bytes, err := toMail(Contacts{
		Contacts: []contact.Contact{{"John", invalidBirthday}}, RemindParams: remindParams,
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
	currentDate := testdata.LocalDate(2014, time.August, 21)
	remindParams := remind.Params{CurrentDate: currentDate, InNbDays: 1}

	bytes, err := toMail(Contacts{
		Contacts: []contact.Contact{{"John", birthday}}, RemindParams: remindParams,
	}, "EN")
	assert.NoError(t, err)
	assert.Equal(t, expectedMail, string(bytes))
}

func Test_should_throw_error_when_subject_template_malformed(t *testing.T) {

	tmpl := new(fakeTemplate)
	tmpl.On("subject").Return("{{{{{")

	bytes, err := resolveMail(Contacts{}, tmpl)

	assert.Error(t, err)
	assert.Empty(t, bytes)
}

func Test_should_throw_error_when_body_template_malformed(t *testing.T) {

	tmpl := new(fakeTemplate)
	tmpl.On("subject").Return("subject")
	tmpl.On("body").Return("{{{{{")

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
