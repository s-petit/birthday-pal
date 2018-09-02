package remind

import "time"

//Reminder holds remind context and conditions
type Reminder struct {
	CurrentDate       time.Time
	NbDaysBeforeBDay  int
	EveryDayUntilBDay bool
}

func (r Reminder) remindOnce(birthDate time.Time) bool {

	remindDay := r.remindDay()

	return !r.EveryDayUntilBDay && remindDay.Day() == birthDate.Day() && remindDay.Month() == birthDate.Month()
}

func (r Reminder) remindEveryDay(birthDate time.Time) bool {

	dateAtMidnight := r.dateAtMidnight()
	remindDay := r.remindDay()

	return r.EveryDayUntilBDay && (birthDate.Day() <= remindDay.Day() && birthDate.Month() <= remindDay.Month()) &&
		(birthDate.Day() >= dateAtMidnight.Day() && birthDate.Month() >= dateAtMidnight.Month())
}

func (r *Reminder) dateAtMidnight() time.Time {
	return time.Date(r.CurrentDate.Year(), r.CurrentDate.Month(), r.CurrentDate.Day(), 0, 0, 0, 0, time.Local)
}

func (r *Reminder) remindDay() time.Time {
	return r.dateAtMidnight().AddDate(0, 0, r.NbDaysBeforeBDay)
}
