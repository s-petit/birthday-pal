package vcardparser

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/mapaiva/vcard-go"
)


func Test_extract_bday_from_one_contact(t *testing.T) {

	contact := `
BEGIN:VCARD
VERSION:3.0
FN:Alexis Foo
N:Foo;Alexis;;;
BDAY:19831028
END:VCARD
`
	vcards := ParseContacts(contact)
	assert.Equal(t, 1, len(vcards))
	assert.Equal(t, "19831028", vcards[0].BirthDay)
}


func Test_extract_bday_from_two_contacts(t *testing.T) {

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

	vcards := ParseContacts(contact)
	assert.Equal(t, 2, len(vcards))
	assert.Equal(t, "19831028", vcards[0].BirthDay)
	assert.Equal(t, "19860425", vcards[1].BirthDay)
}


func Test_parseDate(t *testing.T) {

	date, err := ParseVCardBirthDay(vcard.VCard{BirthDay:"20161225"})

	assert.Equal(t, time.Date(2016, time.December, 25, 0, 0, 0, 0, time.UTC), date)
	assert.NoError(t, err)
}

func Test_parseDate_error(t *testing.T) {
	_, err := ParseVCardBirthDay(vcard.VCard{BirthDay:"20162025"})
	assert.Error(t, err)
}

/*func Test_parseDate_error2(t *testing.T) {
	_, err := ParseVCardBirthDay(vcard.VCard{BirthDay:"2016-20-25 00:00:00 -0000"})
	assert.Error(t, err)
}
*/
