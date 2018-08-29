package contact

import (
	"time"
	"github.com/bearbin/go-age"
	"github.com/s-petit/birthday-pal/birthday"
)

//Contact represents a Contact eligible for reminder because his bday is near.
type Contact struct {
	Name      string
	BirthDate birthday.BirthDate
}

func (r *Contact) LocalBirthDate() time.Time {
	return time.Date(r.BirthDate.Year, r.BirthDate.Month, r.BirthDate.Day, 0, 0, 0, 0, time.Local)
}

// Age return the Age of the contact at his incoming birthday
func (r *Contact) Age(now time.Time) int {
	return age.AgeAt(r.LocalBirthDate(), now)
}

