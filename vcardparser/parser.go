package vcardparser

import (
	"errors"
	"github.com/mapaiva/vcard-go"
	"io"
	"regexp"
	"strings"
	"time"
)

//ParseContacts parses one or many vcards to a vcard.VCard struct.
func ParseContacts(contacts string) ([]vcard.VCard, error) {

	reader := strings.NewReader(contacts)
	multiReader := io.MultiReader(reader)

	return vcard.GetVCardsByReader(multiReader)
}

//ParseVCardBirthDay parse a Vcard BirthDay to a valid time
func ParseVCardBirthDay(vcard vcard.VCard) (time.Time, error) {

	birthdate := vcard.BirthDay

	//YYYY-MM-DD
	vcardBdayAcceptedFormat := regexp.MustCompile("(\\d{4})-(\\d{2})-(\\d{2})")
	//YYYYMMDD
	vcardBdayAcceptedFormat2 := regexp.MustCompile("(\\d{8})")

	if vcardBdayAcceptedFormat.MatchString(birthdate) {
		return time.Parse("2006-01-02 00:00:00 -0000", birthdate+" 00:00:00 -0000")
	}
	if vcardBdayAcceptedFormat2.MatchString(birthdate) {

		runes := []rune(birthdate)
		year := string(runes[0:4])
		month := string(runes[4:6])
		day := string(runes[6:8])

		return time.Parse("2006-01-02 00:00:00 -0000", year+"-"+month+"-"+day+" 00:00:00 -0000")
	}

	return time.Time{}, errors.New("unknown vcard bday format")

}
