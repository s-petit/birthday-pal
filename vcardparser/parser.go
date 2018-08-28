package vcardparser

import (
	"errors"
	"fmt"
	"github.com/mapaiva/vcard-go"
	"github.com/s-petit/birthday-pal/birthday"
	"io"
	"log"
	"regexp"
	"strings"
	"time"
)

//RemindContact represents a Contact eligible for reminder because his bday is near.
type RemindContact struct {
	Name      string
	BirthDate time.Time
	Age       int
}

//ContactsToRemind returns every contacts which the bday occurs in daysBefore days
func ContactsToRemind(cards []vcard.VCard, daysBefore int) []RemindContact {

	var s []RemindContact

	for _, card := range cards {

		date, _ := ParseVCardBirthDay(card)
		now := time.Now()
		shouldRemind := birthday.ShouldRemind(now, date, daysBefore)

		if shouldRemind {
			age := birthday.Age(now, date)
			s = append(s, RemindContact{Name: card.FormattedName, BirthDate: date, Age: age})
		}

	}

	return s
}

//ParseContacts parses one or many vcards to a vcard.VCard struct.
func ParseContacts(contacts string) []vcard.VCard {

	reader := strings.NewReader(contacts)
	multiReader := io.MultiReader(reader)

	cards, e := vcard.GetVCardsByReader(multiReader)

	if e != nil {
		fmt.Println("An error occurred during VCard parsing. Please check that your URL refers to a CardDav endpoint.")
		log.Fatal("ERROR: ", e)
		//os.Exit(1)
	}

	return cards
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
