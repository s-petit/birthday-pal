package vcard

import (
	"github.com/mapaiva/vcard-go"
	"github.com/s-petit/birthday-pal/contact"
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_should_parse_one_contact_from_string_to_vcard(t *testing.T) {

	contact := `
BEGIN:VCARD
VERSION:3.0
FN:Alexis Foo
N:Foo;Alexis;;;
BDAY:19831028
END:VCARD
`
	vcards, err := parseVCard(contact)
	assert.Equal(t, 1, len(vcards))
	assert.Equal(t, "19831028", vcards[0].BirthDay)
	assert.NoError(t, err)
}

func Test_should_parse_two_contacts_from_string_to_vcard(t *testing.T) {

	contact := `
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
`

	vcards, err := parseVCard(contact)
	assert.Equal(t, 2, len(vcards))
	assert.Equal(t, "19831028", vcards[0].BirthDay)
	assert.Equal(t, "19860425", vcards[1].BirthDay)
	assert.NoError(t, err)
}

func Test_should_return_empty_vcard_struct_when_malformed_vcard(t *testing.T) {

	contact := "malformedString"

	vcards, err := parseVCard(contact)
	assert.Equal(t, []vcard.VCard([]vcard.VCard{}), vcards)
	assert.NoError(t, err)
}

func Test_should_parse_one_contact_from_vcard_to_contact(t *testing.T) {

	card := vcard.VCard{}
	card.FormattedName = "John Doe"
	card.BirthDay = "19830409"

	c, err := parseContact(card)
	assert.Equal(t, contact.Contact{"John Doe", testdata.BirthDate(1983, time.April, 9)}, c)
	assert.NoError(t, err)
}

func Test_parsecontanct_should_return_error_when_bdate_malformed(t *testing.T) {

	card := vcard.VCard{}
	card.FormattedName = "John Doe"
	card.BirthDay = "foo"

	c, err := parseContact(card)
	assert.Equal(t, contact.Contact{}, c)
	assert.Error(t, err)
}

func Test_parseDate_YYYYMMDD_should_return_time(t *testing.T) {

	date, err := parseVCardBirthDay(vcard.VCard{BirthDay: "20161225"})

	assert.Equal(t, time.Date(2016, time.December, 25, 0, 0, 0, 0, time.UTC), date)
	assert.NoError(t, err)
}

func Test_parseDate_YYYYMMDD_should_return_error_when_date_is_malformed(t *testing.T) {
	_, err := parseVCardBirthDay(vcard.VCard{BirthDay: "20162025"})
	assert.Error(t, err)
}

func Test_parseDate_YYYYMMDD_should_return_error_when_date_is_malformed2(t *testing.T) {
	_, err := parseVCardBirthDay(vcard.VCard{BirthDay: "9999999"})
	assert.Error(t, err)
}

func Test_parseDate_YYYYMMDD_should_return_error_when_date_is_malformed3(t *testing.T) {
	_, err := parseVCardBirthDay(vcard.VCard{BirthDay: "notADateR"})
	assert.Error(t, err)
}

func Test_parseDate_YYYY_MM_DD_should_return_time(t *testing.T) {

	date, err := parseVCardBirthDay(vcard.VCard{BirthDay: "2016-12-25"})

	assert.Equal(t, time.Date(2016, time.December, 25, 0, 0, 0, 0, time.UTC), date)
	assert.NoError(t, err)
}

func Test_parseDate_YYYY_MM_DD_should_return_error_when_date_is_malformed(t *testing.T) {
	_, err := parseVCardBirthDay(vcard.VCard{BirthDay: "2016-20-25"})
	assert.Error(t, err)
}

func Test_parseDate_should_ignore_empty_bday(t *testing.T) {
	date, err := parseVCardBirthDay(vcard.VCard{BirthDay: ""})
	assert.Equal(t, time.Time{}, date)
	assert.NoError(t, err)
}

func Test_parseDate_should_return_error_when_layout_is_unknown(t *testing.T) {
	_, err := parseVCardBirthDay(vcard.VCard{BirthDay: "25-12-2016"})
	assert.Error(t, err)
}

func Test_should_parse_two_contacts_from_string_to_contact(t *testing.T) {

	vcard := `
BEGIN:VCARD
VERSION:3.0
FN:Alexis Foo
BDAY:19831228
END:VCARD
BEGIN:VCARD
VERSION:3.0
FN:John Bar
BDAY:19861125
END:VCARD
`

	contacts, err := ParseContacts(vcard)
	assert.Equal(t, 2, len(contacts))
	assert.Equal(t, contact.Contact{"Alexis Foo", testdata.BirthDate(1983, time.December, 28)}, contacts[0])
	assert.Equal(t, contact.Contact{"John Bar", testdata.BirthDate(1986, time.November, 25)}, contacts[1])
	assert.NoError(t, err)
}

func Test_ParseContacts_should_return_error_when_vcard_is_malformed(t *testing.T) {

	vcard := "doh"

	contacts, err := ParseContacts(vcard)

	var emptyContacts []contact.Contact
	assert.Equal(t, emptyContacts, contacts)
	assert.NoError(t, err)
}

func Test_ParseContacts_should_return_error_when_vcard_bdate_is_malformed(t *testing.T) {

	vcard := `BEGIN:VCARD
VERSION:3.0
FN:John Bar
BDAY:doh
END:VCARD`

	contacts, err := ParseContacts(vcard)

	var emptyContacts []contact.Contact
	assert.Equal(t, emptyContacts, contacts)
	assert.Error(t, err)
}
