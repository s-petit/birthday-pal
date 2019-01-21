package remind

import (
	"time"
)

//Params holds remind context and conditions passed to the app as parameters
type Params struct {
	// should be now() in most cases, but this field is useful for testing purposes
	CurrentDate time.Time
	InNbDays    int
	// when false, check birthdays which matches exactly CurrentDate + InNbDays
	// when true, check birthdays which matches every days between CurrentDate and CurrentDate + InNbDays
	Inclusive bool
}

//RemindDay returns the day to remind
func (r Params) RemindDay() time.Time {
	return r.dateAtMidnight().AddDate(0, 0, r.InNbDays)
}

func (r Params) dateAtMidnight() time.Time {
	return time.Date(r.CurrentDate.Year(), r.CurrentDate.Month(), r.CurrentDate.Day(), 0, 0, 0, 0, time.Local)
}
