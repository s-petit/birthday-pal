package contact

import (
	"time"
	"github.com/bearbin/go-age"
)

//Contact represents a Contact eligible for reminder because his bday is near.
type Contact struct {
	Name      string
	BirthDate time.Time
	Age int
}

// ShouldRemindBirthday return true if the birthday occurs nbDaysBefore Now
func (c *Contact) ShouldRemindBirthday(nbDaysBefore int, now time.Time) bool {

	delay := del(now, nbDaysBefore)

	c.Age = c.ageAfterBirthday(nbDaysBefore, now)

	return delay.Day() == c.BirthDate.Day() && delay.Month() == c.BirthDate.Month()
}

func del(now time.Time, nbDaysBefore int) time.Time {
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	delay := midnight.AddDate(0, 0, nbDaysBefore)
	return delay
}

// Age return the Age of the contact at his incoming birthday
func (c *Contact) ageAfterBirthday(nbDaysBefore int, now time.Time) int {
	return age.AgeAt(c.BirthDate, del(now, nbDaysBefore))
}