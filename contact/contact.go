package contact

import (
	"github.com/bearbin/go-age"
	"time"
)

//Contact represents a Contact eligible for reminder because his bday is near.
type Contact struct {
	Name      string
	BirthDate time.Time
}

// ShouldRemindBirthday return true if the birthday occurs nbDaysBefore now
func (r *Contact) ShouldRemindBirthday(now time.Time, nbDaysBefore int) bool {

	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	delay := midnight.AddDate(0, 0, nbDaysBefore)

	return delay.Day() == r.BirthDate.Day() && delay.Month() == r.BirthDate.Month()
}

// Age return the Age of the contact at his incoming birthday
func (r *Contact) Age(now time.Time) int {
	return age.AgeAt(r.BirthDate, now)
}
