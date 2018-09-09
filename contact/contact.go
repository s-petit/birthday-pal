package contact

import (
	"github.com/bearbin/go-age"
	"time"
)

//Contact represents a Vcard Contact.
type Contact struct {
	Name      string
	BirthDate time.Time
}

// Age return the Age of the contact at a given date
func (c Contact) Age(date time.Time) int {
	return age.AgeAt(c.BirthDate, date)
}
