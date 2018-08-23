package vcardparser

import (
	"github.com/mapaiva/vcard-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_should_parse_one_contact(t *testing.T) {

	contact := `
BEGIN:VCARD
VERSION:3.0
FN:Alexis Foo
N:Foo;Alexis;;;
BDAY:19831028
END:VCARD
`
	vcards, err := ParseContacts(contact)
	assert.Equal(t, 1, len(vcards))
	assert.Equal(t, "19831028", vcards[0].BirthDay)
	assert.NoError(t, err)
}

func Test_should_parse_two_contacts(t *testing.T) {

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

	vcards, err := ParseContacts(contact)
	assert.Equal(t, 2, len(vcards))
	assert.Equal(t, "19831028", vcards[0].BirthDay)
	assert.Equal(t, "19860425", vcards[1].BirthDay)
	assert.NoError(t, err)
}

func Test_should_return_empty_vcard_struct_when_malformed_vcard(t *testing.T) {

	contact := "malformedString"

	vcards, err := ParseContacts(contact)
	assert.Equal(t, []vcard.VCard([]vcard.VCard{}), vcards)
	assert.NoError(t, err)
}

func Test_parseDate_YYYYMMDD_should_return_time(t *testing.T) {

	date, err := ParseVCardBirthDay(vcard.VCard{BirthDay: "20161225"})

	assert.Equal(t, time.Date(2016, time.December, 25, 0, 0, 0, 0, time.UTC), date)
	assert.NoError(t, err)
}

func Test_parseDate_YYYYMMDD_should_return_error_when_date_is_malformed(t *testing.T) {
	_, err := ParseVCardBirthDay(vcard.VCard{BirthDay: "20162025"})
	assert.Error(t, err)
}

func Test_parseDate_YYYYMMDD_should_return_error_when_date_is_malformed2(t *testing.T) {
	_, err := ParseVCardBirthDay(vcard.VCard{BirthDay: "9999999"})
	assert.Error(t, err)
}

func Test_parseDate_YYYYMMDD_should_return_error_when_date_is_malformed3(t *testing.T) {
	_, err := ParseVCardBirthDay(vcard.VCard{BirthDay: "notADateR"})
	assert.Error(t, err)
}

func Test_parseDate_YYYY_MM_DD_should_return_time(t *testing.T) {

	date, err := ParseVCardBirthDay(vcard.VCard{BirthDay: "2016-12-25"})

	assert.Equal(t, time.Date(2016, time.December, 25, 0, 0, 0, 0, time.UTC), date)
	assert.NoError(t, err)
}

func Test_parseDate_YYYY_MM_DD_should_return_error_when_date_is_malformed(t *testing.T) {
	_, err := ParseVCardBirthDay(vcard.VCard{BirthDay: "2016-20-25"})
	assert.Error(t, err)
}

func Test_parseDate_should_return_error_when_layout_is_unknown(t *testing.T) {
	_, err := ParseVCardBirthDay(vcard.VCard{BirthDay: "25-12-2016"})
	assert.Error(t, err)
}
