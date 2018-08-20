package vcard_parser

import (
	"github.com/mapaiva/vcard-go"
	"strings"
	"io"
	"log"
	"time"
)

func ParseContacts(contacts string) []vcard.VCard {

	reader := strings.NewReader(contacts)
	multiReader := io.MultiReader(reader)

	cards, err := vcard.GetVCardsByReader(multiReader)

	if err != nil {
		log.Fatal(err)
	}

	return cards
}

func ParseVCardBirthDay(vcard vcard.VCard) (time.Time, error) {

	birthdate := vcard.BirthDay

	//TODO faire des regex pour matcher les differents formats de BDAY
	if (len(birthdate) != 8) {
		//return time.Time{}, errors.New("lol")
		return time.Parse("2006-01-02 00:00:00 -0000", birthdate + " 00:00:00 -0000")
	} else {
		runes := []rune(birthdate)
		year := string(runes[0:4])
		month := string(runes[4:6])
		day := string(runes[6:8])

		return time.Parse("2006-01-02 00:00:00 -0000", year + "-" + month + "-" +  day + " 00:00:00 -0000")
	}

}
