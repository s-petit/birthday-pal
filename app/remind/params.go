package remind

import (
	"time"
)

//Params holds remind context and conditions passed to the app as parameters
type Params struct {
	// should be now() in most cases, but this field is useful for testing purposes
	Today    time.Time
	InNbDays int
	// when false, check birthdays which matches exactly Today + InNbDays
	// when true, check birthdays which matches every days between Today and Today + InNbDays
	Inclusive bool
}

//RemindDay returns the day to remind
func (r Params) RemindDay() time.Time {
	return r.todayAtMidnightUTC().AddDate(0, 0, r.InNbDays)
}

func (r Params) todayAtMidnightUTC() time.Time {
	return time.Date(r.Today.Year(), r.Today.Month(), r.Today.Day(), 0, 0, 0, 0, time.UTC)
}
