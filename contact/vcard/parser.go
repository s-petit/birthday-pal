package vcard

import (
	"errors"
	"github.com/mapaiva/vcard-go"
	"io"
	"regexp"
	"strings"
	"time"
	"github.com/s-petit/birthday-pal/contact"
)

//Request holds methods necessary for requesting cardDAV HTTP servers.
type Contact interface {
	Get() ([]contact.Contact, error)
}



//ParseContacts parses a cardDav payload to a Contact struct.
func ParseContacts(cardDavPayload string) ([]contact.Contact, error) {
	vCards, err := parseVCard(cardDavPayload)

	if err != nil {
		return nil, err
	}

	var contacts []contact.Contact
	for _, card := range vCards {
		c, err := parseContact(card)

		if err != nil {
			return nil, err
		}

		contacts = append(contacts, c)
	}
	return contacts, nil
}

//parseContact parses one vcard to a Contact struct.
func parseContact(vcard vcard.VCard) (contact.Contact, error) {

	birthday, err := parseVCardBirthDay(vcard)

	if err != nil {
		return contact.Contact{}, err
	}

	return contact.Contact{Name: vcard.FormattedName, BirthDate: birthday}, nil
}

func parseVCard(contacts string) ([]vcard.VCard, error) {

	reader := strings.NewReader(contacts)
	multiReader := io.MultiReader(reader)

	return vcard.GetVCardsByReader(multiReader)
}

//ParseVCardBirthDay parse a Vcard BirthDay field to a valid golang time
func parseVCardBirthDay(vcard vcard.VCard) (time.Time, error) {

	birthdate := vcard.BirthDay

	if birthdate == "" {
		return time.Time{}, nil
	}

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

		bday, e := time.Parse("2006-01-02 00:00:00 -0000", year+"-"+month+"-"+day+" 00:00:00 -0000")

		return bday, e
	}

	return time.Time{}, errors.New("unknown vcard bday format: " + vcard.BirthDay)

}
