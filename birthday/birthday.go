package birthday

import (
	"time"
)


// Birthdate represents a birth date, without any hour or timezone
type BirthDate struct {
	Year     int
	Month     time.Month
	Day int
}

// ShouldRemind return true if the birthday occurs nbDaysBefore now
func (b *BirthDate) ShouldRemind(now time.Time, nbDaysBefore int) bool {

	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	delay := midnight.AddDate(0, 0, nbDaysBefore)

	return delay.Day() == b.Day && delay.Month() == b.Month
}

