package remind

import (
	"time"
)

//Criteria holds remind criteria passed to the app as parameters
type Criteria struct {
	// should be now() in most cases, but this field is useful for testing purposes
	Today    time.Time
	InNbDays int
	// when false, check birthdays which matches exactly Today + InNbDays
	// when true, check birthdays which matches every days between Today and Today + InNbDays
	Inclusive bool
}

//RemindDay returns the day to remind
func (r Criteria) RemindDay() time.Time {
	return r.todayAtMidnightUTC().AddDate(0, 0, r.InNbDays)
}

func (r Criteria) todayAtMidnightUTC() time.Time {
	return time.Date(r.Today.Year(), r.Today.Month(), r.Today.Day(), 0, 0, 0, 0, time.UTC)
}
