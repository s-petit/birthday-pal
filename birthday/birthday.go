package birthday

import (
	"github.com/bearbin/go-age"
	"time"
)

// Remind return true if the birthday occurs nbDaysBefore now
func Remind(now time.Time, birthDate time.Time, nbDaysBefore int) bool {

	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	delay := midnight.AddDate(0, 0, nbDaysBefore)

	return delay.Day() == birthDate.Day() && delay.Month() == birthDate.Month()
}

// Age return the age of the contact at his incoming birthday
func Age(now time.Time, birthDate time.Time) int {
	return age.AgeAt(birthDate, now)
}
